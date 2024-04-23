package post

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheus-hrm/trampos/types"
	"github.com/matheus-hrm/trampos/utils"
)

type Handler struct {
	store types.PostsStore
}

func NewHandler(store types.PostsStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", h.handleGetPosts).Methods(http.MethodGet)
	router.HandleFunc("/post/{id}", h.handleGetPostByID).Methods(http.MethodGet)
	router.HandleFunc("/post", h.handleCreatePost).Methods(http.MethodPost)
	router.HandleFunc("/post", h.handleUpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/post/{id}", h.handleDeletePost).Methods(http.MethodDelete)
}

func (h *Handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	utils.WriteJSON(w, http.StatusOK, ps)
}
func (h *Handler) handleGetPostByID(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	
}
func (h *Handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleDeletePost(w http.ResponseWriter, r *http.Request) {

}