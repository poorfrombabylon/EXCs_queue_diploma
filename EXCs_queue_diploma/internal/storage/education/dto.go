package education

import (
	educationDomain "excs_queue/internal/domain/education"
	profileDomain "excs_queue/internal/domain/profile"
	"github.com/google/uuid"
	"time"
)

type EducationDTO struct {
	Id                       uuid.UUID  `db:"id"`
	ProfileID                uuid.UUID  `db:"profile_id"`
	Education                []byte     `db:"education"`
	FieldOfStudy             *string    `db:"field_of_study"`
	DegreeName               *string    `db:"degree_name"`
	School                   string     `db:"school"`
	SchoolLinkedinProfileUrl *string    `db:"school_linkedin_profile_url"`
	Description              *string    `db:"description"`
	LogoUrl                  *string    `db:"logo_url"`
	Grade                    *string    `db:"grade"`
	ActivitiesAndSocieties   *string    `db:"activities_and_societies"`
	StartDate                *time.Time `db:"start_date"`
	EndDate                  *time.Time `db:"end_date"`
	CreatedAt                time.Time  `db:"created_at"`
	UpdatedAt                time.Time  `db:"updated_at"`
}

func NewEducationFromDTO(dto EducationDTO) educationDomain.Education {
	return educationDomain.NewEducationWithID(
		educationDomain.EducationID(dto.Id),
		profileDomain.ProfileID(dto.ProfileID),
		dto.FieldOfStudy,
		dto.DegreeName,
		dto.School,
		dto.SchoolLinkedinProfileUrl,
		dto.Description,
		dto.LogoUrl,
		dto.Grade,
		dto.ActivitiesAndSocieties,
		dto.StartDate,
		dto.EndDate,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}

func NewEducationListFromDTO(dto []EducationDTO) []educationDomain.Education {
	res := make([]educationDomain.Education, 0, len(dto))

	for _, i := range dto {
		res = append(res, NewEducationFromDTO(i))
	}

	return res
}
