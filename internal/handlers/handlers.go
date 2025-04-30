package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/AeddGynvael3110/sprint6/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	result, err := service.DetectAndConvert(string(data))
	if err != nil {
		http.Error(w, "Error while converting data", http.StatusInternalServerError)
		return
	}

	// Создание директории, если она не существует
	os.MkdirAll("output", os.ModePerm)

	fileName := "output/" + time.Now().UTC().Format("2006-01-02_15-04-05") + ".txt"
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Unable to write to file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
