package controller

import (
	"backend/internal/app/handler"
	"fmt"
	"net/http"
	"os"
)

const (
	FormPage = "frontend/index.html"
)

type Web struct {
}

func NewWeb() Web {
	return Web{}
}

// FormPage загрузка страницы формы опроса
func (web Web) FormPage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile(FormPage)
	if err != nil {
		handler.InternalServerError(w, "Ошибка загрузки страницы")
		return
	}

	fmt.Fprintf(w, string(page))
}
