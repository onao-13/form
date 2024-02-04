package payload

import (
	"fmt"
	"net/mail"
)

// PeopleEmployment структура опроса
type PeopleEmployment struct {
	OrganizationName           string `json:"organization_name"`
	ShortOrganizationName      string `json:"short_organization_name"`
	TypeParticipant            string `json:"type_participant"`
	PositionLeaderOrganization string `json:"position_leader_organization"`
	FIOLeaderOrganization      string `json:"fio_leader_organization"`
	PositionResponsiblePerson  string `json:"position_responsible_person"`
	FIOResponsiblePerson       string `json:"fio_responsible_person"`
	TelephoneNumber            string `json:"telephone_number"`
	Email                      string `json:"email"`
}

// Validate валидация структуры
func (p PeopleEmployment) Validate() error {
	switch {
	case len(p.OrganizationName) == 0:
		return fmt.Errorf("название организации не указано")
	case len(p.ShortOrganizationName) == 0:
		return fmt.Errorf("короткое название организации не указано")
	case len(p.TypeParticipant) == 0:
		return fmt.Errorf("тип отвечающего не указан")
	case len(p.PositionLeaderOrganization) == 0:
		return fmt.Errorf("должность управляющего не указана")
	case len(p.FIOLeaderOrganization) == 0:
		return fmt.Errorf("ФИО управляющего не указано")
	case len(p.TelephoneNumber) == 0:
		return fmt.Errorf("неверный формат номера телефона")
	case len(p.Email) == 0:
		return fmt.Errorf("почта не указана")
	case !p.isMailValid():
		return fmt.Errorf("почта неправильного формата")
	}
	return nil
}

// isMailValid валидна ли почта
func (p PeopleEmployment) isMailValid() bool {
	_, err := mail.ParseAddress(p.Email)
	return err == nil
}
