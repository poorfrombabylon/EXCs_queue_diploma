package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"excs_queue/internal/config"
	certificationDomain "excs_queue/internal/domain/certification"
	educationDomain "excs_queue/internal/domain/education"
	experienceDomain "excs_queue/internal/domain/experience"
	profileDomain "excs_queue/internal/domain/profile"
	queueDomain "excs_queue/internal/domain/queue"
	"excs_queue/internal/libdb"
	"excs_queue/internal/storage"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"log"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("error while init config")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	} else {
		log.Println("connected to db")
	}

	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry, err := storage.NewStorageRegistry(ctx, libDBWrapper, *cfg)
	if err != nil {
		log.Fatal("failed to init storageRegistry:", err.Error())
	}

	err = startJob(ctx, storageRegistry)
}

func startJob(
	ctx context.Context,
	storageRegistry *storage.Storages,
) error {
	subscriber := storageRegistry.QueueConverter.SubscribeToSubscriberChannel(ctx)
	ch := subscriber.Channel()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				errClose := subscriber.Close()
				if errClose != nil {
					return errClose
				}
				fmt.Println("ctx cancel")
			case m := <-ch:
				res := queueDomain.FullProfileInfo{}

				fmt.Println(m.Payload)

				err := json.Unmarshal([]byte(m.Payload), &res)
				if err != nil {
					fmt.Println("err while json unmarshal", err.Error())
				}

				log.Println(res)
				fmt.Println()

				profile := profileDomain.NewProfile(res.FirstName, res.LastName, res.Country, res.City, res.State, res.Gender, res.Occupation, res.Summary, res.PublicIdentifier)

				err = storageRegistry.ProfilesStorage.CreateProfile(ctx, profile)
				if err != nil {
					log.Println("err while create profile", err.Error())
				}

				if len(res.Experiences) != 0 {
					for _, i := range res.Experiences {
						startDate := getDate(i.StartDate)
						endDate := getDate(i.EndDate)

						exp := experienceDomain.NewExperience(
							profile.GetProfileID(),
							i.Position,
							i.CompanyName,
							i.Location,
							i.Description,
							startDate,
							endDate,
						)

						err = storageRegistry.ExperienceStorage.CreateExperience(ctx, exp)
						if err != nil {
							log.Println("err while create experience", err.Error())
						}
					}
				}

				if len(res.Education) != 0 {
					for _, i := range res.Education {
						startDate := getDate(i.StartDate)
						endDate := getDate(i.EndDate)

						edu := educationDomain.NewEducation(
							profile.GetProfileID(),
							i.FieldOfStudy,
							i.DegreeName,
							i.School,
							i.SchoolLinkedinProfileUrl,
							i.Description,
							i.LogoUrl,
							i.Grade,
							i.ActivitiesAndSocieties,
							startDate,
							endDate,
						)

						err = storageRegistry.EducationStorage.CreateEducation(ctx, edu)
						if err != nil {
							log.Println("err while create education", err.Error())
						}
					}
				}

				if len(res.Certifications) != 0 {
					for _, i := range res.Certifications {
						startDate := getDate(i.StartDate)
						endDate := getDate(i.EndDate)

						cert := certificationDomain.NewCertification(
							profile.GetProfileID(),
							i.Name,
							i.Authority,
							i.LicenseNumber,
							i.DisplaySource,
							i.Url,
							i.AuthorityLinkedinURL,
							startDate,
							endDate,
						)

						err = storageRegistry.CertificationStorage.CreateCertification(ctx, cert)
						if err != nil {
							log.Println("err while create cert:", err.Error())
						}
					}
				}

				err = storageRegistry.CaptchaStorage.CreateCaptcha(ctx, res.CaptchaMeet)
				if err != nil {
					log.Println("err while create captcha", err.Error())
				}
			}
		}
	})

	err := group.Wait()
	log.Println("finish")
	if err != nil {
		return err
	}

	return nil
}

func getDate(date *queueDomain.Date) *time.Time {
	if date != nil && date.Day != nil && date.Month != nil && date.Year != nil {
		day := strconv.Itoa(*date.Day)
		if len(day) == 1 {
			day = "0" + day
		}

		month := strconv.Itoa(*date.Month)
		if len(month) == 1 {
			month = "0" + month
		}

		res, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%v-%v", *date.Year, month, day))
		if err != nil {
			log.Println("err while parsing date:", err.Error())
		}

		return &res
	}

	return nil
}
