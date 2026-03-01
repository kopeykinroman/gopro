package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kopeykinroman/gopro/internal/handlers"
	"github.com/kopeykinroman/gopro/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestHandleRouteMethod(t *testing.T) {
	store := storage.NewStorage()
	h := handlers.NewHTTPHandlers(store)

	// Добавил валидную ссылку в хранилище
	originalURL := "https://yandex.ru"
	short := store.Add(originalURL)

	tests := []struct {
		name           string
		method         string
		path           string
		body           string
		contentType    string
		expectedStatus int
		expectedLoc    string
	}{
		//POST + "/" + correct
		{
			name:           "POST. Success",
			method:         http.MethodPost,
			path:           "/",
			body:           "https://habr.ru/",
			contentType:    "text/plain",
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "POST. Re-adding site ",
			method:         http.MethodPost,
			path:           "/",
			body:           "https://habr.ru/",
			contentType:    "text/plain",
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "POST. End Point not correct",
			method:         http.MethodPost,
			path:           "/lalala",
			body:           "https://ya.ru",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST. ContentType not correct",
			method:         http.MethodPost,
			path:           "/",
			body:           "https://ya.ru",
			contentType:    "application/json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST. Body is clear",
			method:         http.MethodPost,
			path:           "/",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},

		{
			name:           "GET. Success",
			method:         http.MethodGet,
			path:           "/" + short,
			contentType:    "text/plain",
			expectedStatus: http.StatusTemporaryRedirect,
		},
		{
			name:           "GET. Not found",
			method:         http.MethodGet,
			path:           "/unknown",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Method DELETE - Unsupported method",
			method:         http.MethodDelete,
			path:           "/",
			body:           "https://vlgu.ru/",
			contentType:    "text/plain",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bodyReader *strings.Reader
			if tt.body != "" {
				bodyReader = strings.NewReader(tt.body)
			} else {
				bodyReader = strings.NewReader("")
			}

			req := httptest.NewRequest(tt.method, tt.path, bodyReader)

			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}

			rec := httptest.NewRecorder()

			h.HandleRouteMethod(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedLoc != "" {
				assert.Equal(t, tt.expectedLoc, rec.Header().Get("Location"))
			}
		})
	}
}
