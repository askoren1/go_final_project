package handler

import "github.com/askoren1/go_final_project/internal/repository"

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler { //функция-конструктор для типа Handler
	return &Handler{
		repo: repo,
	}
}
