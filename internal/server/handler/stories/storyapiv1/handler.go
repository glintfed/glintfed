package storyapiv1

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Carousel(w http.ResponseWriter, r *http.Request)
	SelfCarousel(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Publish(w http.ResponseWriter, r *http.Request)
	CarouselNext(w http.ResponseWriter, r *http.Request)
	PublishNext(w http.ResponseWriter, r *http.Request)
	MentionAutocomplete(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Viewed(w http.ResponseWriter, r *http.Request)
	Comment(w http.ResponseWriter, r *http.Request)
	Viewers(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Carousel(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Carousel")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SelfCarousel(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.SelfCarousel")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Add")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Publish(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Publish")
	defer span.End()
	// TODO: Implement
}

func (h *handler) CarouselNext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.CarouselNext")
	defer span.End()
	// TODO: Implement
}

func (h *handler) PublishNext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.PublishNext")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MentionAutocomplete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.MentionAutocomplete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Delete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Viewed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Viewed")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Comment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Comment")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Viewers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Stories.ApiV1.Viewers")
	defer span.End()
	// TODO: Implement
}
