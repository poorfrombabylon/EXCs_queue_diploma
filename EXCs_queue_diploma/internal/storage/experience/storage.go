package experience

import (
	"context"
	experienceDomain "excs_queue/internal/domain/experience"
	profileDomain "excs_queue/internal/domain/profile"
	"excs_queue/internal/libdb"
	"github.com/Masterminds/squirrel"
)

const experienceTable = "experience"

type ExperienceStorage interface {
	CreateExperience(ctx context.Context, experience experienceDomain.Experience) error
	GetExperienceByProfileID(ctx context.Context, profileID profileDomain.ProfileID) ([]experienceDomain.Experience, error)
	DeleteExperienceByID(ctx context.Context, id experienceDomain.ExperienceID) error
	GetOldBunchData(ctx context.Context) ([]ExperienceDTO, error)
}

type experienceStorage struct {
	db libdb.DB
}

func NewExperience(db libdb.DB) ExperienceStorage {
	return &experienceStorage{
		db: db,
	}
}

func (e *experienceStorage) CreateExperience(ctx context.Context, experience experienceDomain.Experience) error {
	query := squirrel.Insert(experienceTable).
		Columns(
			"id",
			"profile_id",
			"experience",
			"position",
			"company_name",
			"location",
			"description",
			"start_date",
			"end_date",
			"created_at",
			"updated_at",
		).
		Values(
			experience.GetExperienceID().String(),
			experience.GetProfileID().String(),
			nil,
			experience.GetPosition(),
			experience.GetCompanyName(),
			experience.GetLocation(),
			experience.GetDescription(),
			experience.GetStartDate(),
			experience.GetEndDate(),
			experience.GetCreatedAt(),
			experience.GetUpdatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := e.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (e *experienceStorage) GetExperienceByProfileID(ctx context.Context, profileID profileDomain.ProfileID) ([]experienceDomain.Experience, error) {
	query := squirrel.Select(
		"id",
		"profile_id",
		"experience",
		"position",
		"company_name",
		"location",
		"description",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From(experienceTable).
		Where(squirrel.Eq{"profile_id": profileID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []ExperienceDTO

	err := e.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewExperienceListFromDTO(result), nil
}

func (e *experienceStorage) GetOldBunchData(ctx context.Context) ([]ExperienceDTO, error) {
	query := squirrel.Select(
		"id",
		"profile_id",
		"experience",
		"position",
		"company_name",
		"location",
		"description",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From(experienceTable).
		Where(squirrel.Eq{"company_name": ""}).
		Limit(1000).
		PlaceholderFormat(squirrel.Dollar)

	var result []ExperienceDTO

	err := e.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (e *experienceStorage) DeleteExperienceByID(ctx context.Context, id experienceDomain.ExperienceID) error {
	query := squirrel.Delete(experienceTable).
		Where(squirrel.Eq{"id": id.String()}).
		PlaceholderFormat(squirrel.Dollar)

	return e.db.Delete(ctx, query)
}
