package education

import (
	"context"
	educationDomain "excs_queue/internal/domain/education"
	profileDomain "excs_queue/internal/domain/profile"
	"excs_queue/internal/libdb"
	"github.com/Masterminds/squirrel"
)

const educationTable = "education"

type EducationStorage interface {
	CreateEducation(ctx context.Context, education educationDomain.Education) error
	GetEducationByProfileID(ctx context.Context, profileID profileDomain.ProfileID) ([]educationDomain.Education, error)
	GetOldBunchData(ctx context.Context) ([]EducationDTO, error)
	DeleteEducationByID(ctx context.Context, id educationDomain.EducationID) error
}

type educationStorage struct {
	db libdb.DB
}

func NewEducationStorage(db libdb.DB) EducationStorage {
	return &educationStorage{
		db: db,
	}
}

func (e *educationStorage) CreateEducation(ctx context.Context, education educationDomain.Education) error {
	query := squirrel.Insert(educationTable).
		Columns(
			"id",
			"profile_id",
			"education",
			"field_of_study",
			"degree_name",
			"school",
			"school_linkedin_profile_url",
			"description",
			"logo_url",
			"grade",
			"activities_and_societies",
			"start_date",
			"end_date",
			"created_at",
			"updated_at",
		).
		Values(
			education.GetEducationID().String(),
			education.GetProfileID().String(),
			nil,
			education.GetFieldOfStudy(),
			education.GetDegreeName(),
			education.GetSchool(),
			education.GetSchoolLinkedinProfileUrl(),
			education.GetDescription(),
			education.GetLogoUrl(),
			education.GetGrade(),
			education.GetActivitiesAndSocieties(),
			education.GetStartDate(),
			education.GetEndDate(),
			education.GetCreatedAt(),
			education.GetUpdatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := e.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (e *educationStorage) GetEducationByProfileID(ctx context.Context, profileID profileDomain.ProfileID) ([]educationDomain.Education, error) {
	query := squirrel.Select(
		"id",
		"profile_id",
		"education",
		"field_of_study",
		"degree_name",
		"school",
		"school_linkedin_profile_url",
		"description",
		"logo_url",
		"grade",
		"activities_and_societies",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From(educationTable).
		Where(squirrel.Eq{"profile_id": profileID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var res []EducationDTO

	err := e.db.Select(ctx, query, &res)
	if err != nil {
		return nil, err
	}

	return NewEducationListFromDTO(res), nil
}

func (e *educationStorage) GetOldBunchData(ctx context.Context) ([]EducationDTO, error) {
	query := squirrel.Select(
		"id",
		"profile_id",
		"education",
		"field_of_study",
		"degree_name",
		"school",
		"school_linkedin_profile_url",
		"description",
		"logo_url",
		"grade",
		"activities_and_societies",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	).
		From(educationTable).
		Where(squirrel.Eq{"school": ""}).
		Limit(1000).
		PlaceholderFormat(squirrel.Dollar)

	var res []EducationDTO

	err := e.db.Select(ctx, query, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (e *educationStorage) DeleteEducationByID(ctx context.Context, id educationDomain.EducationID) error {
	query := squirrel.Delete(educationTable).
		Where(squirrel.Eq{"id": id.String()}).
		PlaceholderFormat(squirrel.Dollar)

	return e.db.Delete(ctx, query)
}
