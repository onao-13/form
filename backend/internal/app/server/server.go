package server

import (
	"backend/internal/app/api"
	"backend/internal/app/controller"
	"backend/internal/app/database"
	"backend/internal/app/service"
	"backend/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
)

type Server struct {
	ctx context.Context
	srv *http.Server
	cfg config.Config
	log *log.Logger
}

func NewServer(ctx context.Context, cfg config.Config, log *log.Logger) Server {
	return Server{
		ctx: ctx,
		cfg: cfg,
		log: log,
	}
}

func (s Server) Start() {
	pool, err := pgxpool.New(s.ctx, s.cfg.DbUrl())
	if err != nil {
		panic(fmt.Sprintf("Ошибка подлючения к базе: %s", err))
	}

	s.log.Println("Сервер запускается")

	peopleDb := database.NewPeopleEmployment(s.ctx, pool)

	peopleService := service.NewPeopleEmployment(peopleDb, s.log)

	peopleController := controller.NewPeopleEmployment(peopleService)
	webController := controller.NewWeb()

	handler := api.NewHandler(peopleController, webController)

	s.srv = &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: handler,
	}

	s.log.Println("Сервер запущен на порту " + s.cfg.Port)

	if err := s.srv.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("Ошибка запуска сервера: %s", err))
	}
}
