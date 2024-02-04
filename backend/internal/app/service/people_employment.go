package service

import (
	"backend/internal/app/database"
	"backend/internal/app/payload"
	"log"
)

type PeopleEmployment struct {
	db  database.PeopleEmployment
	log *log.Logger
}

func NewPeopleEmployment(db database.PeopleEmployment, log *log.Logger) PeopleEmployment {
	return PeopleEmployment{db: db, log: log}
}

func (p PeopleEmployment) Send(people payload.PeopleEmployment) error {
	if err := p.db.Save(people); err != nil {
		p.log.Println("Ошибка сохранения данных: ", err)
		return err
	}
	return nil
}
