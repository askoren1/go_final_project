package handler

import (
	"net/http"
)

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	//Провеверка на входящие данные
	// Если ошибка, возвращаем 404 ошибку
	if err := h.repo.GetList(); err != nil {
		// Возвращаем 500 ошибку

	}
    // Возвращаем OK
	w.Write([]byte("OK"))

}


