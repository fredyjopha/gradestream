package api

import (
	"io"
	"net/http"
	"os"
)

func handleIngest(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("batch")
	if err != nil {
		http.Error(w,"fichier manquant", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dstTmp, err := os.CreateTemp("", "gradestram-batch-*")
	if err != nil {
		http.Error(w, "erreur serveur", http.StatusInternalServerError)
		return
	}
	defer dstTmp.Close()

	if _, err := io.Copy(dstTmp, file); err != nil {
		http.Error(w, "erreur d'écriture", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}