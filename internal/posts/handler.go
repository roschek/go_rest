package posts

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/roschek/go_rest.git/internal/handlers"
	"github.com/roschek/go_rest.git/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	postsURL = "/api/posts"
	postURL  = "/api/post/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandlerPost(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(postsURL, h.GetList)
	router.GET(postURL, h.GetPosttUUID)
	router.POST(postURL, h.CreatePost)
	router.PUT(postURL, h.UpdatePost)
	router.PATCH(postURL, h.PartialyUpdatePost)
	router.DELETE(postURL, h.DeletePost)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is users list"))

}

func (h *handler) GetPosttUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Some interesting Text"))

}
func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Some interesting Text"))

}
func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Some interesting Text"))

}
func (h *handler) PartialyUpdatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Some interesting Text"))

}
func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Some interesting Text"))

}
