package main

//import (
//	"context"
//	"database/sql"
//	"encoding/json"
//	"excs_queue/internal/config"
//	educationDomain "excs_queue/internal/domain/education"
//	experienceDomain "excs_queue/internal/domain/experience"
//	profileDomain "excs_queue/internal/domain/profile"
//	queueDomain "excs_queue/internal/domain/queue"
//	"excs_queue/internal/libdb"
//	"excs_queue/internal/storage"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/jmoiron/sqlx"
//	_ "github.com/lib/pq"
//	"log"
//	"os/signal"
//	"strconv"
//	"sync"
//	"syscall"
//	"time"
//)
//
//func main() {
//	var wg sync.WaitGroup
//
//	ctx, cancel := signal.NotifyContext(
//		context.Background(),
//		syscall.SIGHUP,
//		syscall.SIGINT,
//		syscall.SIGTERM,
//		syscall.SIGQUIT,
//	)
//	defer cancel()
//
//	cfg, err := config.InitConfig()
//	if err != nil {
//		log.Fatal("error while init config")
//	}
//
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBname)
//
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		log.Fatal("failed to connect to database:", err.Error())
//	} else {
//		log.Println("connected to db")
//	}
//
//	dbx := sqlx.NewDb(db, "pgx")
//	libDBWrapper := libdb.NewSQLXDB(dbx)
//
//	storageRegistry, err := storage.NewStorageRegistry(ctx, libDBWrapper, *cfg)
//	if err != nil {
//		log.Fatal("failed to init storageRegistry:", err.Error())
//	}
//
//	wg.Add(2)
//
//	go startExperienceParsing(ctx, storageRegistry, &wg)
//	go startEducationParsing(ctx, storageRegistry, &wg)
//
//	wg.Wait()
//}
//
//func startExperienceParsing(ctx context.Context, storageRegistry *storage.Storages, wg *sync.WaitGroup) {
//	bunch := 1
//
//	for {
//		fmt.Println("experience bunch number:", bunch)
//
//		exp, err := storageRegistry.ExperienceStorage.GetOldBunchData(ctx)
//		if err != nil {
//			log.Println("err while GetOldBunchData experience", err.Error())
//		}
//
//		if exp == nil || len(exp) == 0 {
//			break
//		}
//
//		fmt.Println("experience bunch len:", len(exp))
//
//		for _, i := range exp {
//			var res []queueDomain.Experience
//
//			err = json.Unmarshal(i.Experience, &res)
//			if err != nil {
//				log.Println("err while unmarshal experience for profileID:", i.ProfileID.String(), err.Error())
//			}
//
//			for _, j := range res {
//				fmt.Println("experience profile_id", i.ProfileID.String())
//
//				readyExperience := processExperience(j, profileDomain.ProfileID(i.ProfileID), i.CreatedAt)
//
//				err = storageRegistry.ExperienceStorage.CreateExperience(ctx, readyExperience)
//				if err != nil {
//					log.Printf("|prfileId: %s\n|err while create ready exp: %s", i.ProfileID.String(), err.Error())
//				}
//			}
//
//			err = storageRegistry.ExperienceStorage.DeleteExperienceByID(ctx, experienceDomain.ExperienceID(i.ID))
//			if err != nil {
//				log.Printf("|prfileId: %s\n|err while delete exp: %s", i.ProfileID.String(), err.Error())
//			}
//		}
//
//		bunch += 1
//	}
//
//	wg.Done()
//}
//
//func startEducationParsing(ctx context.Context, storageRegistry *storage.Storages, wg *sync.WaitGroup) {
//	bunch := 1
//
//	for {
//		fmt.Println("education bunch number:", bunch)
//
//		edu, err := storageRegistry.EducationStorage.GetOldBunchData(ctx)
//		if err != nil {
//			log.Println("err while GetOldBunchData education", err.Error())
//		}
//
//		if edu == nil || len(edu) == 0 {
//			break
//		}
//
//		fmt.Println("education bunch len:", len(edu))
//
//		for _, i := range edu {
//			var res []queueDomain.Education
//
//			err = json.Unmarshal(i.Education, &res)
//			if err != nil {
//				log.Println("err while unmarshal education for profileID:", i.ProfileID.String(), err.Error())
//			}
//
//			for _, j := range res {
//				fmt.Println("education profile_id", i.ProfileID.String())
//
//				readyEducation := processEducation(j, profileDomain.ProfileID(i.ProfileID), i.CreatedAt)
//
//				err = storageRegistry.EducationStorage.CreateEducation(ctx, readyEducation)
//				if err != nil {
//					log.Printf("|prfileId: %s\n|err while create ready edu: %s", i.ProfileID.String(), err.Error())
//				}
//			}
//
//			err = storageRegistry.EducationStorage.DeleteEducationByID(ctx, educationDomain.EducationID(i.Id))
//			if err != nil {
//				log.Printf("|prfileId: %s\n|err while delete edu: %s", i.ProfileID.String(), err.Error())
//			}
//		}
//
//		bunch += 1
//	}
//
//	wg.Done()
//}
//
//func processEducation(
//	e queueDomain.Education,
//	profileID profileDomain.ProfileID,
//	createdAt time.Time,
//) educationDomain.Education {
//	startDate := getDate(e.StartDate)
//	endDate := getDate(e.EndDate)
//
//	res := educationDomain.NewEducationWithID(
//		educationDomain.EducationID(uuid.New()),
//		profileID,
//		e.FieldOfStudy,
//		e.DegreeName,
//		e.School,
//		e.SchoolLinkedinProfileUrl,
//		e.Description,
//		e.LogoUrl,
//		e.Grade,
//		e.ActivitiesAndSocieties,
//		startDate,
//		endDate,
//		createdAt,
//	)
//
//	return res
//}
//
//func processExperience(
//	e queueDomain.Experience,
//	profileID profileDomain.ProfileID,
//	createdAt time.Time,
//) experienceDomain.Experience {
//	startDate := getDate(e.StartDate)
//	endDate := getDate(e.EndDate)
//
//	res := experienceDomain.NewExperienceWithID(
//		experienceDomain.ExperienceID(uuid.New()),
//		profileID,
//		e.Position,
//		e.CompanyName,
//		e.Location,
//		e.Description,
//		startDate,
//		endDate,
//		createdAt,
//	)
//
//	return res
//}
//
//func getDate(date *queueDomain.Date) *time.Time {
//	if date != nil && date.Day != nil && date.Month != nil && date.Year != nil {
//		if *date.Day == 0 && *date.Month == 0 && *date.Year == 0 {
//			return nil
//		}
//
//		day := strconv.Itoa(*date.Day)
//		if len(day) == 1 {
//			day = "0" + day
//		}
//
//		month := strconv.Itoa(*date.Month)
//		if len(month) == 1 {
//			month = "0" + month
//		}
//
//		res, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%v-%v", *date.Year, month, day))
//		if err != nil {
//			log.Println("err while parsing date:", err.Error())
//		}
//
//		return &res
//	}
//
//	return nil
//}
