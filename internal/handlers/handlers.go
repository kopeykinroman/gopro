package handlers

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kopeykinroman/gopro/internal/storage"
)

type HTTPHandlers struct {
	store   *storage.Storage
	router  http.Handler
	baseURL string
}

// Добавил Functional Options Pattern что бы не исправлять код тестов
type Option func(*HTTPHandlers)

func WithBaseURL(baseURL string) Option {
	return func(h *HTTPHandlers) {
		h.baseURL = baseURL
	}
}

func NewHTTPHandlers(store *storage.Storage, opts ...Option) *HTTPHandlers {
	h := &HTTPHandlers{
		store:   store,
		baseURL: "http://localhost:8080", // Значение по умолчанию
	}

	for _, opt := range opts {
		opt(h)
	}

	h.initRouter()
	return h

}

func (h *HTTPHandlers) initRouter() {
	r := chi.NewRouter()

	r.MethodFunc(http.MethodPost, "/", h.handleAddShortener)
	r.MethodFunc(http.MethodGet, "/{short:[^/]+}", h.handleGetUrl)

	r.NotFound(h.badRequest)
	r.MethodNotAllowed(h.badRequest)

	h.router = r
}

// HandleRouteMethod в зависимости от типа метода вызывает подходящий обработчик
func (h *HTTPHandlers) HandleRouteMethod(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// HandleAddShortener добавляет в хранилище новый url
//
// POST /
//
// request: Content-Type = text/plain, body = URL
//
// Success: response := shortlink как text/plan и code 201 Created
// Error: response := code 400
func (h *HTTPHandlers) handleAddShortener(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Проверяем что заголовок = text/plain
	contentType := r.Header.Get("Content-Type")

	if !strings.HasPrefix(contentType, "text/plain") {
		msg := "Bad Request. Error in Header. Expected text/plain. Received by [" + contentType + "]"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Чтение тела запроса и его проверка
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad Request. Failed to read body.", http.StatusBadRequest)
		return
	}

	url := strings.TrimSpace(string(body))
	if url == "" {
		msg := "Bad Request. Error in Body. Expected length > 0. Received by [" + string(body) + "]"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	shortLink := h.store.Add(url)
	fullShortLink := h.baseURL + "/" + shortLink

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fullShortLink))
}

// HandleGetUrl по короткой ссылке возвращает исходный url
//
// GET /{shortlink}
//
// Success: response := Оригинальный URL как Location и code 307
// Error: response := code 400
func (h *HTTPHandlers) handleGetUrl(w http.ResponseWriter, r *http.Request) {
	// Извлекает оригинальный url
	short := chi.URLParam(r, "short")

	if short == "" || strings.Contains(short, "/") {
		h.badRequest(w, r)
		return
	}

	url, err := h.store.Get(short)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "ShortLink not found", http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *HTTPHandlers) badRequest(w http.ResponseWriter, r *http.Request) {
	msg := "Bad Request. Error in Method type or Path. Received by Method [" +
		r.Method + "] and Path [" + r.URL.Path + "]"

	http.Error(w, msg, http.StatusBadRequest)
}
