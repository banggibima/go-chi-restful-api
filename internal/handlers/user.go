package handlers

import (
	"net/http"

	"github.com/banggibima/go-chi-restful-api/internal/entities"
	"github.com/banggibima/go-chi-restful-api/internal/usecases"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	render.JSON(w, r, users)
}

func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.UserUseCase.GetUserByID(id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	render.JSON(w, r, user)
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input entities.User

	if err := render.DecodeJSON(r.Body, &input); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	user, err := h.UserUseCase.CreateUser(input)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var input entities.User

	if err := render.DecodeJSON(r.Body, &input); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	user, err := h.UserUseCase.UpdateUser(id, input)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	render.JSON(w, r, user)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.UserUseCase.DeleteUser(id); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]interface{}{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusNoContent)
}
