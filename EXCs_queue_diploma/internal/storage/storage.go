package storage

import (
	"context"
	"excs_queue/internal/config"
	"excs_queue/internal/libdb"
	"excs_queue/internal/storage/captcha"
	"excs_queue/internal/storage/certification"
	"excs_queue/internal/storage/education"
	"excs_queue/internal/storage/experience"
	"excs_queue/internal/storage/profiles"
	"excs_queue/internal/storage/redis/queue_converter"
)

type Storages struct {
	EducationStorage     education.EducationStorage
	ExperienceStorage    experience.ExperienceStorage
	ProfilesStorage      profiles.ProfilesStorage
	CaptchaStorage       captcha.CaptchaStorage
	CertificationStorage certification.CertificationStorage
	QueueConverter       queue_converter.QueueConverter
}

func NewStorageRegistry(_ context.Context, db libdb.DB, cfg config.Config) (*Storages, error) {
	redisQueue, err := queue_converter.NewQueueConverter(cfg.QueueRedis)
	if err != nil {
		return nil, err
	}

	return &Storages{
		EducationStorage:     education.NewEducationStorage(db),
		ExperienceStorage:    experience.NewExperience(db),
		ProfilesStorage:      profiles.NewProfileStorage(db),
		CaptchaStorage:       captcha.NewCaptchaStorage(db),
		CertificationStorage: certification.NewCertificationStorage(db),
		QueueConverter:       redisQueue,
	}, nil
}
