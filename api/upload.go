// Fichier: /api/upload.go

package handler

import (
	"encoding/json"
	"net/http"
	"awsutils/awsconnect"
)

// Response structure pour encapsuler la réponse de l'upload
type UploadResponse struct {
	Message  string `json:"message"`
	TrueName string `json:"trueName"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	trueName, err := awsconnect.Put(header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Créez une instance de UploadResponse avec le message et le trueName
	response := UploadResponse{
		Message:  "Image uploaded and path saved successfully",
		TrueName: trueName,
	}

	// Définissez le Content-Type de la réponse à application/json
	w.Header().Set("Content-Type", "application/json")

	// Encodez la réponse en JSON et écrivez-la dans le ResponseWriter
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
