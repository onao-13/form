package controller

import (
	"backend/internal/app/handler"
	"backend/internal/app/payload"
	"backend/internal/app/service"
	"encoding/json"
	"net/http"
)

type PeopleEmployment struct {
	service service.PeopleEmployment
}

func NewPeopleEmployment(service service.PeopleEmployment) PeopleEmployment {
	return PeopleEmployment{service: service}
}

// Send отправка данных на сервер
func (p PeopleEmployment) Send(w http.ResponseWriter, r *http.Request) {
	var (
		people payload.PeopleEmployment
		err    error
	)

	if err = json.NewDecoder(r.Body).Decode(&people); err != nil {
		handler.BadRequest(w, "Ошибка декодирования JSON")
		return
	}

	if err = people.Validate(); err != nil {
		handler.BadRequest(w, err.Error())
		return
	}

	if err = p.service.Send(people); err != nil {
		handler.InternalServerError(w, "Ошибка сохранения данных")
		return
	}

	handler.Create(w, "")
}
