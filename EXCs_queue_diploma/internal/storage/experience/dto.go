package experience

import (
	experienceDomain "excs_queue/internal/domain/experience"
	profileDomain "excs_queue/internal/domain/profile"
	"github.com/google/uuid"
	"time"
)

type ExperienceDTO struct {
	ID          uuid.UUID  `db:"id",`
	ProfileID   uuid.UUID  `db:"profile_id"`
	Experience  []byte     `db:"experience"`
	Position    *string    `db:"position"`
	CompanyName string     `db:"company_name"`
	Location    *string    `db:"location"`
	Description *string    `db:"description"`
	StartDate   *time.Time `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

func NewExperienceFromDTO(experience ExperienceDTO) experienceDomain.Experience {
	return experienceDomain.NewExperienceWithID(
		experienceDomain.ExperienceID(experience.ID),
		profileDomain.ProfileID(experience.ProfileID),
		experience.Position,
		experience.CompanyName,
		experience.Location,
		experience.Description,
		experience.StartDate,
		experience.EndDate,
		experience.CreatedAt,
		experience.UpdatedAt,
	)
}

func NewExperienceListFromDTO(dto []ExperienceDTO) []experienceDomain.Experience {
	result := make([]experienceDomain.Experience, 0, len(dto))

	for _, i := range dto {
		result = append(result, NewExperienceFromDTO(i))
	}

	return result
}
