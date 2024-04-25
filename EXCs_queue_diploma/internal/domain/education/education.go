package education

import (
	profileDomain "excs_queue/internal/domain/profile"
	"github.com/google/uuid"
	"time"
)

type EducationID uuid.UUID

func (e EducationID) String() string {
	return uuid.UUID(e).String()
}

type Education struct {
	id                       EducationID
	profileID                profileDomain.ProfileID
	fieldOfStudy             *string
	degreeName               *string
	school                   string
	schoolLinkedinProfileUrl *string
	description              *string
	logoUrl                  *string
	grade                    *string
	activitiesAndSocieties   *string
	startDate                *time.Time
	endDate                  *time.Time
	createdAt                time.Time
	updatedAt                time.Time
}

func NewEducation(
	profileID profileDomain.ProfileID,
	fieldOfStudy *string,
	degreeName *string,
	school string,
	schoolLinkedinProfileUrl *string,
	description *string,
	logoUrl *string,
	grade *string,
	activitiesAndSocieties *string,
	startDate *time.Time,
	endDate *time.Time,
) Education {
	return Education{
		id:                       EducationID(uuid.New()),
		profileID:                profileID,
		fieldOfStudy:             fieldOfStudy,
		degreeName:               degreeName,
		school:                   school,
		schoolLinkedinProfileUrl: schoolLinkedinProfileUrl,
		description:              description,
		logoUrl:                  logoUrl,
		grade:                    grade,
		activitiesAndSocieties:   activitiesAndSocieties,
		startDate:                startDate,
		endDate:                  endDate,
		createdAt:                time.Now().In(time.UTC),
		updatedAt:                time.Now().In(time.UTC),
	}
}

func NewEducationWithID(
	id EducationID,
	profileID profileDomain.ProfileID,
	fieldOfStudy *string,
	degreeName *string,
	school string,
	schoolLinkedinProfileUrl *string,
	description *string,
	logoUrl *string,
	grade *string,
	activitiesAndSocieties *string,
	startDate *time.Time,
	endDate *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) Education {
	return Education{
		id:                       id,
		profileID:                profileID,
		fieldOfStudy:             fieldOfStudy,
		degreeName:               degreeName,
		school:                   school,
		schoolLinkedinProfileUrl: schoolLinkedinProfileUrl,
		description:              description,
		logoUrl:                  logoUrl,
		grade:                    grade,
		activitiesAndSocieties:   activitiesAndSocieties,
		startDate:                startDate,
		endDate:                  endDate,
		createdAt:                createdAt,
		updatedAt:                updatedAt,
	}
}

func (e Education) GetEducationID() EducationID {
	return e.id
}

func (e Education) GetProfileID() profileDomain.ProfileID {
	return e.profileID
}

func (e Education) GetFieldOfStudy() *string {
	return e.fieldOfStudy
}

func (e Education) GetDegreeName() *string {
	return e.degreeName
}

func (e Education) GetSchool() string {
	return e.school
}

func (e Education) GetSchoolLinkedinProfileUrl() *string {
	return e.schoolLinkedinProfileUrl
}

func (e Education) GetDescription() *string {
	return e.description
}

func (e Education) GetLogoUrl() *string {
	return e.logoUrl
}

func (e Education) GetGrade() *string {
	return e.grade
}

func (e Education) GetActivitiesAndSocieties() *string {
	return e.activitiesAndSocieties
}

func (e Education) GetStartDate() *time.Time {
	return e.startDate
}

func (e Education) GetEndDate() *time.Time {
	return e.endDate
}

func (e Education) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e Education) GetUpdatedAt() time.Time {
	return e.updatedAt
}
