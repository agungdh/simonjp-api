package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"

	"github.com/agungdh/simonjp-api/internal/models"
)

type ItemHandler struct {
	DB *bun.DB
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	ctx := r.Context()

	if err := h.DB.NewSelect().Model(&items).Order("id DESC").Scan(ctx); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if items == nil {
		items = []models.Item{}
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var item models.Item
	ctx := r.Context()
	if err := h.DB.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, item)
}

type itemInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in itemInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if in.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}

	now := time.Now()
	item := models.Item{
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	ctx := r.Context()
	if _, err := h.DB.NewInsert().Model(&item).Exec(ctx); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var in itemInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if in.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}

	ctx := r.Context()
	var item models.Item
	if err := h.DB.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "item not found")
			return
		}
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	item.Name = in.Name
	item.Description = in.Description
	item.UpdatedAt = time.Now()

	if _, err := h.DB.NewUpdate().Model(&item).Where("id = ?", id).Exec(ctx); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	res, err := h.DB.NewDelete().Model((*models.Item)(nil)).Where("id = ?", id).Exec(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	_ = res
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// Ensure context import is used (for future timeouts).
var _ = context.Background
