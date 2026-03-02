package handlers

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/kopeykinroman/gopro/internal/storage"
)

type HTTPHandlers struct {
	store *storage.Storage
}

func NewHTTPHandlers(store *storage.Storage) *HTTPHandlers {
	return &HTTPHandlers{
		store: store,
	}
}

// HandleRouteMethod в зависимости от типа метода вызывает подходящий обработчик
func (h *HTTPHandlers) HandleRouteMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/" {
			h.handleAddShortener(w, r)
			return
		}

	case http.MethodGet:
		path := strings.Trim(r.URL.Path, "/")
		if path != "" && !strings.Contains(path, "/") {
			h.handleGetUrl(w, r)
			return
		}
	}

	msg := "Bad Request. Error in Method type or Path. Received by Method [" + r.Method + "] and Path [" + r.URL.Path + "]"
	http.Error(w, msg, http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(shortLink))
}

// HandleGetUrl по короткой ссылке возвращает исходный url
//
// GET /{shortlink}
//
// Success: response := Оригинальный URL как Location и code 307
// Error: response := code 400
func (h *HTTPHandlers) handleGetUrl(w http.ResponseWriter, r *http.Request) {

	// Извлекает оригинальный url
	short := strings.TrimPrefix(r.URL.Path, "/")
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
