package services

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"time"
	"path/filepath"
)

type FileInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SizeBytes  int64  `json:"size_bytes"`
	UploadedAt string `json:"uploaded_at"` // Можно использовать time.Time и настроить его форматирование
	MimeType   string
	Data       io.ReadSeeker
}


func (f *FileInfo) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}


func FileDownload(filename string, w http.ResponseWriter) (string, int) {
	file, err := os.Open(filename)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return "Файл не найден", http.StatusNotFound
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Type", "text/plain")

	_, err = io.Copy(w, file)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusExpectationFailed)
		return "Error load data", http.StatusExpectationFailed
	}
	
	w.WriteHeader(http.StatusOK)
	return "load data", http.StatusOK
}


func (f *FileInfo) FileGetInfo(filename string, w http.ResponseWriter) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return false
	}
	ext := filepath.Ext(info.Name())

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // fallback, если расширение неизвестно
	}

	fileInfo := &FileInfo{
		ID:         "",                               
		Name:       info.Name(),   
		SizeBytes:  info.Size(), 
		UploadedAt: info.ModTime().Format(time.RFC3339), 
		MimeType:   mimeType, 
		Data:       nil,// Поток данных тут не нужен, если отдаём только метаданные
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileInfo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	return true
}