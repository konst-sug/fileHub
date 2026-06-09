package handlers
import (
	"io"
	"net/http"
	"os"
	"path/filepath"


	"discword.ru/fCloud/internal/services"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

// var filename = "/home/konst/_work/go-proj-fileCloud/test.txt"

var uploadDir = "/home/konst/java"

type FileInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SizeBytes  int64  `json:"size_bytes"`
	UploadedAt string `json:"uploaded_at"` // Можно использовать time.Time и настроить его форматирование
	MimeType   string
	Data       io.ReadSeeker
}


func FileHandler(w http.ResponseWriter, r *http.Request) {
	// Скачать последнюю версию файла по его ID.
	if !AreParamsPresent(r, "fileId") {
		http.Error(w, "File ID is missing", http.StatusBadRequest)
		return
	}
	// определить путь к файлу по id?
	filePath := r.URL.Query().Get("fileId")
	_, _ = services.FileDownload(filePath, w)
    
}

func FileInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Получить метаданные файла без скачивания самого содержимого.
	if !AreParamsPresent(r, "fileId") {
		http.Error(w, "File ID is missing", http.StatusBadRequest)
		return
	}
	// определить путь к файлу по id?
	fileId := chi.URLParam(r, "fileId")
	var f services.FileInfo
	ok := f.FileGetInfo(fileId, w)
	if !ok {
		return 
	}

}

func LoadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Загрузка нового файла.
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error to red file %v", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dstPath := filepath.Join(uploadDir, handler.Filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		log.Fatalf("Error %v", err)
		http.Error(w, "Error to open file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error to copy file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DelFileHandler(w http.ResponseWriter, r *http.Request) {
	// Удалить файл и версии
	if !AreParamsPresent(r, "fileId") {
		http.Error(w, "File ID is missing", http.StatusBadRequest)
		return
	}
	//проверка наличия файла
	//операция с файлом
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
