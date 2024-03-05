package handler

import (
	"net/http"
	"awsutils/awsconnect"
	"strings"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extraire le nom du fichier de l'URL
	filename := strings.TrimPrefix(r.URL.Path, "/api/image/")
	if filename == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	file, err := awsconnect.Get(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png") // Ajustez selon le type d'image
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
