package main

import (
	
	"net/http"
	"os"

	"discword.ru/fCloud/internal/handlers"
	"discword.ru/fCloud/internal/middleware"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)
func init() {
	// Настраиваем форматтер для вывода в формате JSON (удобно для парсинга системами сбора логов)
	log.SetFormatter(&log.JSONFormatter{})
	// Выводим логи в стандартный вывод
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}


func main() {
    r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.ValidatePathMiddleware)

	// public
	r.Get("/login", handlers.LoginHandler)
	r.Post("/register", handlers.RegisterHandler)

	// protected 
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
			r.Route("/files", func(r chi.Router) {
			r.Post("/", handlers.LoadFileHandler)

			r.Get("/", handlers.FileHandler)
			r.Get("/info", handlers.FileInfoHandler)
			r.Delete("/", handlers.DelFileHandler)

			r.Get("/versions", handlers.LoadFileVersionHandler)
			r.Get("/version", handlers.LoadFileVersionHandler)
			r.Post("/restore", handlers.RestoreFileVersionHandler)
		})

			r.Route("/directories", func(r chi.Router) {
				r.Post("/", handlers.CreateDirHandler)
				r.Put("/", handlers.RenameDirHandler)
				r.Delete("/", handlers.DelDirHandler)
				r.Get("/contents", handlers.ContentDirHandler)
			})

			r.Route("/shares", func(r chi.Router) {
				r.Post("/files", handlers.ShareFileHandler)
				r.Post("/directories", handlers.ShareDirHandler)

				r.Get("/", handlers.ShareCheckTokenHandler)
				r.Get("/metadata", handlers.ShareMetadataTokenHandler)
				r.Get("/download", handlers.ShareDownloadTokenHandler)
			})

			r.Route("/user", func(r chi.Router) {
				r.Post("/", handlers.LoadFileHandler)

				r.Get("/", handlers.FileHandler)
				r.Get("/info", handlers.FileInfoHandler)
				r.Delete("/", handlers.DelFileHandler)
			})
	})

r.Post("/config/upload-dir", handlers.SetUploadDirHandler)
   
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}