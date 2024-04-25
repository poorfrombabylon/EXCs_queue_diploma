package profiles

import (
	"context"
	profileDomain "excs_queue/internal/domain/profile"
	"excs_queue/internal/libdb"
	"github.com/Masterminds/squirrel"
)

const profilesTable = "profiles"

type ProfilesStorage interface {
	CreateProfile(ctx context.Context, profile profileDomain.Profile) error
	GetBunchProfiles(ctx context.Context) ([]profileDomain.Profile, error)
	DeleteProfileByID(ctx context.Context, profileID profileDomain.ProfileID) error
}

type profileStorage struct {
	db libdb.DB
}

func NewProfileStorage(db libdb.DB) ProfilesStorage {
	return &profileStorage{
		db: db,
	}
}

func (p *profileStorage) CreateProfile(ctx context.Context, profile profileDomain.Profile) error {
	query := squirrel.Insert(profilesTable).
		Columns(
			"id",
			"first_name",
			"last_name",
			"country",
			"city",
			"state",
			"gender",
			"occupation",
			"summary",
			"linkedin_id",
			"created_at",
			"updated_at",
			"last_checked_at",
		).
		Values(
			profile.GetProfileID().String(),
			profile.GetFirstName(),
			profile.GetLastName(),
			profile.GetCountry(),
			profile.GetCity(),
			profile.GetState(),
			profile.GetGender(),
			profile.GetOccupation(),
			profile.GetSummary(),
			profile.GetLinkedInID(),
			profile.GetCreatedAt(),
			profile.GetUpdatedAt(),
			profile.GetLastCheckedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := p.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *profileStorage) GetBunchProfiles(ctx context.Context) ([]profileDomain.Profile, error) {
	query := squirrel.Select(
		"id",
		"first_name",
		"last_name",
		"country",
		"city",
		"state",
		"gender",
		"occupation",
		"summary",
		"linkedin_id",
		"created_at",
		"updated_at",
		"last_checked_at",
	).
		From(profilesTable).
		Limit(12556).
		PlaceholderFormat(squirrel.Dollar)

	var result []ProfileDTO

	err := p.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewProfilesFromDTOList(result), nil
}

func (p *profileStorage) DeleteProfileByID(ctx context.Context, profileID profileDomain.ProfileID) error {
	query := squirrel.Delete(profilesTable).
		Where(squirrel.Eq{"id": profileID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	return p.db.Delete(ctx, query)
}
