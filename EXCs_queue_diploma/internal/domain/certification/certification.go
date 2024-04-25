package certification

import (
	profileDomain "excs_queue/internal/domain/profile"
	"github.com/google/uuid"
	"time"
)

type CertificationID uuid.UUID

func (c CertificationID) String() string {
	return uuid.UUID(c).String()
}

type Certification struct {
	id                   CertificationID
	profileID            profileDomain.ProfileID
	name                 string
	authority            string
	licenseNumber        *string
	displaySource        *string
	url                  *string
	authorityLinkedinURL *string
	startDate            *time.Time
	endDate              *time.Time
	createdAt            time.Time
	updatedAt            time.Time
}

func NewCertification(
	profileID profileDomain.ProfileID,
	name string,
	authority string,
	licenseNumber *string,
	displaySource *string,
	url *string,
	authorityLinkedinURL *string,
	startDate *time.Time,
	endDate *time.Time,
) Certification {
	return Certification{
		id:                   CertificationID(uuid.New()),
		profileID:            profileID,
		name:                 name,
		authority:            authority,
		licenseNumber:        licenseNumber,
		displaySource:        displaySource,
		url:                  url,
		authorityLinkedinURL: authorityLinkedinURL,
		startDate:            startDate,
		endDate:              endDate,
		createdAt:            time.Now().In(time.UTC),
		updatedAt:            time.Now().In(time.UTC),
	}
}

func (c Certification) GetCertificationID() CertificationID {
	return c.id
}

func (c Certification) GetProfileID() profileDomain.ProfileID {
	return c.profileID
}

func (c Certification) GetName() string {
	return c.name
}

func (c Certification) GetAuthority() string {
	return c.authority
}

func (c Certification) GetLicenseNumber() *string {
	return c.licenseNumber
}

func (c Certification) GetDisplaySource() *string {
	return c.displaySource
}

func (c Certification) GetUrl() *string {
	return c.url
}

func (c Certification) GetAuthorityLinkedinURL() *string {
	return c.authorityLinkedinURL
}

func (c Certification) GetStartDate() *time.Time {
	return c.startDate
}

func (c Certification) GetEndDate() *time.Time {
	return c.endDate
}

func (c Certification) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c Certification) GetUpdatedAt() time.Time {
	return c.updatedAt
}
