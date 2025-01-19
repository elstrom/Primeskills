package handlers

import (
	"comments-api/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// CreateComment - API untuk membuat komentar
func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	// Parse JSON body ke struct Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Simpan komentar ke database
	if err := h.DB.Create(&comment).Error; err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	// Kirim response dengan data komentar yang baru
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetComments - API untuk mendapatkan daftar komentar
func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	var comments []models.Comment

	// Ambil semua komentar dari database
	if err := h.DB.Find(&comments).Error; err != nil {
		http.Error(w, "Failed to get comments", http.StatusInternalServerError)
		return
	}

	// Kirim response dengan daftar komentar
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// DeleteComment - API untuk menghapus komentar berdasarkan ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Ambil ID komentar dari URL parameter
	id := mux.Vars(r)["id"]

	// Cari komentar berdasarkan ID
	var comment models.Comment
	if err := h.DB.First(&comment, id).Error; err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	// Hapus komentar dari database
	if err := h.DB.Delete(&comment).Error; err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	// Kirim response success
	w.WriteHeader(http.StatusNoContent)
}
