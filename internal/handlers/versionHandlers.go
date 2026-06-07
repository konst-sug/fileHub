package handlers

import(
	"net/http"
)


func VersionsFileHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "fileId") {
		http.Error(w, "File Not Found", http.StatusBadRequest)
		return
	}
	//проверка наличия файла
	//операция с файлом
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func LoadFileVersionHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "fileId", "VersionId") {
		http.Error(w, "Missing required parameters: fileId and/or versionId", http.StatusBadRequest)
		return
	}
	//проверка наличия файла
	//операция с файлом
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func RestoreFileVersionHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "fileId", "versionId") {
		http.Error(w, "Missing required parameters: fileId and/or versionId", http.StatusBadRequest)
		return
	}
	//проверка наличия файла
	//операция с файлом
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
