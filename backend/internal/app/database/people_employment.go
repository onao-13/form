package database

import (
	"backend/internal/app/payload"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PeopleEmployment struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewPeopleEmployment(ctx context.Context, pool *pgxpool.Pool) PeopleEmployment {
	return PeopleEmployment{
		ctx:  ctx,
		pool: pool,
	}
}

func (p PeopleEmployment) Save(people payload.PeopleEmployment) error {
	sql := `
	INSERT INTO
		people_employment(
		  	organization_name,
		  	short_organization_name,
  			type_participant,
		  	position_leader_organization,
		  	fio_leader_organization,
		  	position_responsible_person,
		  	fio_responsible_person,
		  	telephone_number,
		  	email
	  	)
	VALUES(
	       	@organization_name,
			@short_organization_name,
  			@type_participant,
		  	@position_leader_organization,
		  	@fio_leader_organization,
	       	@position_responsible_person,
		  	@fio_responsible_person,
		  	@telephone_number,
		  	@email
	)`

	args := pgx.NamedArgs{
		"organization_name":            people.OrganizationName,
		"short_organization_name":      people.ShortOrganizationName,
		"type_participant":             people.TypeParticipant,
		"position_leader_organization": people.PositionLeaderOrganization,
		"fio_leader_organization":      people.FIOLeaderOrganization,
		"position_responsible_person":  people.PositionResponsiblePerson,
		"fio_responsible_person":       people.FIOResponsiblePerson,
		"telephone_number":             people.TelephoneNumber,
		"email":                        people.Email,
	}

	if _, err := p.pool.Exec(p.ctx, sql, args); err != nil {
		return err
	}

	return nil
}
