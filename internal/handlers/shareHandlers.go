package handlers

import (
	"net/http"
)


func ShareFileHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "fileId") {
		http.Error(w, "FileId parameter requirerd!", http.StatusBadRequest)
		return
	}
	ok, msg := ArePostParamsPresent(r, "password")
	if !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	// проверка допуска, создание токена
	SuccessResponse(w)
}

func ShareDirHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "directoryId") {
		http.Error(w, "DirectotyId parameter required", http.StatusBadRequest)
		return
	}
	ok, msg := ArePostParamsPresent(r, "password")
	if !ok {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	// проверка допуска, создание токена
	SuccessResponse(w)
}

func ShareCheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "token") {
		http.Error(w, "Token parameter required!", http.StatusBadRequest)
		return
	}
	// проверка токена
	SuccessResponse(w)
}

func ShareMetadataTokenHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "token") {
		http.Error(w, "Token parameter required!", http.StatusBadRequest)
		return
	}
	// проверка токена
	// получение данных
	SuccessResponse(w)
}

func ShareDownloadTokenHandler(w http.ResponseWriter, r *http.Request) {
	if !AreParamsPresent(r, "token") {
		http.Error(w, "Token parameter required!", http.StatusBadRequest)
		return
	}
	// проверка токена
	// получение данных
	SuccessResponse(w)
}
