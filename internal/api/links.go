package api

import (
	"cortico/internal/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

func Short(w http.ResponseWriter, r *http.Request) {
	linkService := r.Context().Value("links").(models.LinksService)

	link := models.Link{}
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		slog.Error("Error decoding the request", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(http.StatusBadRequest, "invalid JSON request data"))
		return
	}

	if err := link.ValidateURL(); err != nil {
		slog.Error("Error validating the URL", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewResponse(http.StatusBadRequest, "invalid URL"))
		return
	}

	if err := link.GenerateShortLink(); err != nil {
		slog.Error("Error generating the short URL", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(http.StatusBadRequest, "error generating the short URL"))
		return
	}

	if err := linkService.Create(&link); err != nil {
		slog.Error("Error storing the link", "err", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.NewAPIError(http.StatusBadRequest, "error storing the link"))
		return
	}

	response := models.NewResponse(http.StatusCreated, link.FullLink())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
