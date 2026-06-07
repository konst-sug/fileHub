package handlers

import(
	"net/http"
	
)


func CreateDirHandler(w http.ResponseWriter, r *http.Request) {
	err, msg := ArePostParamsPresent(r, "name")
	if !err {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	//операция создания директории
	SuccessResponse(w)
}

func RenameDirHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "directoryId") {
		http.Error(w, "Missing directoryId in params", http.StatusBadRequest)
		return
	}
	//операция с директорией
	SuccessResponse(w)
}

func DelDirHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "directoryId") {
		http.Error(w, "Missing directoryId in params", http.StatusBadRequest)
		return
	}
	// операция с директорией
	SuccessResponse(w)
}

func ContentDirHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "directoryId") {
		http.Error(w, "Missing directoryId in params", http.StatusBadRequest)
		return
	}
	// операция с директорией
	SuccessResponse(w)
}

func SetUploadDirHandler(w http.ResponseWriter, r *http.Request) {
	err, msg := ArePostParamsPresent(r, "name")
	if !err {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	//операция создания директории
	SuccessResponse(w)
}