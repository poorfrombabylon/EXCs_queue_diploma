package certification

import (
	"context"
	"excs_queue/internal/domain/certification"
	"excs_queue/internal/libdb"
	"github.com/Masterminds/squirrel"
)

const certificationTable = "certification"

type CertificationStorage interface {
	CreateCertification(ctx context.Context, certification certification.Certification) error
}

type certificationStorage struct {
	db libdb.DB
}

func NewCertificationStorage(db libdb.DB) CertificationStorage {
	return &certificationStorage{db: db}
}

func (c *certificationStorage) CreateCertification(ctx context.Context, certification certification.Certification) error {
	query := squirrel.Insert(certificationTable).
		Columns(
			"id",
			"profile_id",
			"name",
			"authority",
			"license_number",
			"display_source",
			"url",
			"authority_linkedin_url",
			"start_date",
			"end_date",
			"created_at",
			"updated_at",
		).
		Values(
			certification.GetCertificationID().String(),
			certification.GetProfileID().String(),
			certification.GetName(),
			certification.GetAuthority(),
			certification.GetLicenseNumber(),
			certification.GetDisplaySource(),
			certification.GetUrl(),
			certification.GetAuthorityLinkedinURL(),
			certification.GetStartDate(),
			certification.GetEndDate(),
			certification.GetCreatedAt(),
			certification.GetUpdatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
