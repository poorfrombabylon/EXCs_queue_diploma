package profiles

import (
	profileDomain "excs_queue/internal/domain/profile"
	"github.com/google/uuid"
	"time"
)

type ProfileDTO struct {
	Id            uuid.UUID `db:"id"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	Country       *string   `db:"country"`
	City          *string   `db:"city"`
	State         *string   `db:"state"`
	Gender        *string   `db:"gender"`
	Occupation    *string   `db:"occupation"`
	Summary       *string   `db:"summary"`
	LinkedinID    string    `db:"linkedin_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	LastCheckedAt time.Time `db:"last_checked_at"`
}

func NewProfileFromDTO(dto ProfileDTO) profileDomain.Profile {
	return profileDomain.NewProfileWithID(
		profileDomain.ProfileID(dto.Id),
		dto.FirstName,
		dto.LastName,
		dto.Country,
		dto.City,
		dto.State,
		dto.Gender,
		dto.Occupation,
		dto.Summary,
		dto.LinkedinID,
		dto.CreatedAt,
		dto.UpdatedAt,
		dto.LastCheckedAt,
	)
}

func NewProfilesFromDTOList(dto []ProfileDTO) []profileDomain.Profile {
	result := make([]profileDomain.Profile, 0, len(dto))

	for _, i := range dto {
		result = append(result, NewProfileFromDTO(i))
	}

	return result
}
