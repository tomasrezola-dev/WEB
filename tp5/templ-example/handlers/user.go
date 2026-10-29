package handlers

import (
	"net/http"
	"templ-example/models"
	"templ-example/views"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if err := views.UserPage(models.Users).Render(r.Context(), w); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
