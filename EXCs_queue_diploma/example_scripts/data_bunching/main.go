package main

import (
	"context"
	"database/sql"
	"excs_queue/internal/config"
	"excs_queue/internal/libdb"
	"excs_queue/internal/storage"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os/signal"
	"syscall"
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

	psqlInfoMain := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBname)

	psqlInfoFirstBunch := fmt.Sprintf("host=%s port=5432 user=scraper password=%s dbname=second_bunch sslmode=disable", cfg.Postgres.Host, cfg.Postgres.Password)

	dbMain, err := sql.Open("postgres", psqlInfoMain)
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	} else {
		log.Println("connected to dbMain")
	}

	dbFirstBunch, err := sql.Open("postgres", psqlInfoFirstBunch)
	if err != nil {
		log.Fatal("failed to connect to database fb:", err.Error())
	} else {
		log.Println("connected to dbFirstBunch")
	}

	dbxMain := sqlx.NewDb(dbMain, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbxMain)

	dbxFB := sqlx.NewDb(dbFirstBunch, "pgx")
	libDBWrapperFB := libdb.NewSQLXDB(dbxFB)

	storageRegistryMain, err := storage.NewStorageRegistry(ctx, libDBWrapper, *cfg)
	if err != nil {
		log.Fatal("failed to init storageRegistryMain:", err.Error())
	}

	storageRegistryFB, err := storage.NewStorageRegistry(ctx, libDBWrapperFB, *cfg)
	if err != nil {
		log.Fatal("failed to init storageRegistryFB:", err.Error())
	}

	err = startJob(ctx, storageRegistryMain, storageRegistryFB)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func startJob(ctx context.Context, storageRegistryMain, storageRegistryFB *storage.Storages) error {
	profiles, err := storageRegistryMain.ProfilesStorage.GetBunchProfiles(ctx)
	if err != nil {
		log.Println("err while get profiles:", err.Error())
		return err
	}

	fmt.Println("got profiles:", len(profiles))

	count := 0

	for _, i := range profiles {
		count += 1
		fmt.Println("profile number:", count)
		log.Println("profileID", i.GetProfileID().String())

		experience, err := storageRegistryMain.ExperienceStorage.GetExperienceByProfileID(ctx, i.GetProfileID())
		if err != nil {
			log.Printf("while get experience for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
		}

		education, err := storageRegistryMain.EducationStorage.GetEducationByProfileID(ctx, i.GetProfileID())
		if err != nil {
			log.Printf("while get education for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
		}

		err = storageRegistryFB.ProfilesStorage.CreateProfile(ctx, i)
		if err != nil {
			log.Printf("while create profile for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
			return err
		}

		for _, exp := range experience {
			err = storageRegistryFB.ExperienceStorage.CreateExperience(ctx, exp)
			if err != nil {
				log.Printf("while create exper for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
				return err
			}

			err = storageRegistryMain.ExperienceStorage.DeleteExperienceByID(ctx, exp.GetExperienceID())
			if err != nil {
				log.Printf("while delete exper for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
				return err
			}
		}

		for _, edu := range education {
			err = storageRegistryFB.EducationStorage.CreateEducation(ctx, edu)
			if err != nil {
				log.Printf("while create education for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
				return err
			}

			err = storageRegistryMain.EducationStorage.DeleteEducationByID(ctx, edu.GetEducationID())
			if err != nil {
				log.Printf("while delete education for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
				return err
			}
		}

		err = storageRegistryMain.ProfilesStorage.DeleteProfileByID(ctx, i.GetProfileID())
		if err != nil {
			log.Printf("while delete profile for profileId: %s err: %s", i.GetProfileID().String(), err.Error())
			return err
		}
	}

	return nil
}
