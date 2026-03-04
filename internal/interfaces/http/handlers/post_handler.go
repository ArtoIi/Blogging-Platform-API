package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArtoIi/Blogging-Platform-API/internal/application"
	"github.com/ArtoIi/Blogging-Platform-API/internal/domain"
)

type PostHandler struct {
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post domain.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search")

	posts, err := h.service.GetAllPosts(searchTerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	post, err := h.service.GetPostByID(id)
	if err != nil {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)

	if err := h.service.DeletePost(id); err != nil {
		http.Error(w, "ID não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)
	var post domain.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	post.ID = id

	if err := h.service.UpdatePost(&post); err != nil {
		http.Error(w, "ID não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}
