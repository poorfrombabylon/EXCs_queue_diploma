package profile

import (
	"github.com/google/uuid"
	"time"
)

type ProfileID uuid.UUID

func (a ProfileID) String() string {
	return uuid.UUID(a).String()
}

type Profile struct {
	id            ProfileID
	firstName     string
	lastName      string
	country       *string
	city          *string
	state         *string
	gender        *string
	occupation    *string
	summary       *string
	linkedInID    string
	createdAt     time.Time
	updatedAt     time.Time
	lastCheckedAt time.Time
}

func NewProfile(
	firstName string,
	lastName string,
	country *string,
	city *string,
	state *string,
	gender *string,
	occupation *string,
	summary *string,
	linkedInID string,
) Profile {
	return Profile{
		id:            ProfileID(uuid.New()),
		firstName:     firstName,
		lastName:      lastName,
		country:       country,
		city:          city,
		state:         state,
		gender:        gender,
		occupation:    occupation,
		summary:       summary,
		linkedInID:    linkedInID,
		createdAt:     time.Now().In(time.UTC),
		updatedAt:     time.Now().In(time.UTC),
		lastCheckedAt: time.Now().In(time.UTC),
	}
}

func NewProfileWithID(
	id ProfileID,
	firstName string,
	lastName string,
	country *string,
	city *string,
	state *string,
	gender *string,
	occupation *string,
	summary *string,
	linkedInID string,
	createdAt time.Time,
	updatedAt time.Time,
	lastCheckedAt time.Time,
) Profile {
	return Profile{
		id:            id,
		firstName:     firstName,
		lastName:      lastName,
		country:       country,
		city:          city,
		state:         state,
		gender:        gender,
		occupation:    occupation,
		summary:       summary,
		linkedInID:    linkedInID,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		lastCheckedAt: lastCheckedAt,
	}
}

func (p Profile) GetProfileID() ProfileID {
	return p.id
}

func (p Profile) GetFirstName() string {
	return p.firstName
}

func (p Profile) GetLastName() string {
	return p.lastName
}

func (p Profile) GetCountry() *string {
	return p.country
}

func (p Profile) GetCity() *string {
	return p.city
}

func (p Profile) GetState() *string {
	return p.state
}

func (p Profile) GetGender() *string {
	return p.gender
}

func (p Profile) GetOccupation() *string {
	return p.occupation
}

func (p Profile) GetSummary() *string {
	return p.summary
}

func (p Profile) GetLinkedInID() string {
	return p.linkedInID
}

func (p Profile) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p Profile) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p Profile) GetLastCheckedAt() time.Time {
	return p.lastCheckedAt
}
