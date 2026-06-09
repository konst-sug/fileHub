package middleware

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var validPaths = []string{
	"/files/restore/",
	"/directories/contents/",
	"/directories/",
	"/shares/metadata",
	"/shares/download",
	"/shares/directories",
}

func ValidatePathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentPath := r.URL.Path

		log.Printf("[DEBUG] Request path: '%s'", currentPath)
		// log.Printf("[DEBUG] Request URL: %s", r.URL.String())
		// log.Printf("[DEBUG] Request RawPath: '%s'", r.URL.RawPath)

		isValid := false
		for _, validPath := range validPaths {
			match := currentPath == validPath
			// log.Printf("[DEBUG] [%d] '%s' == '%s' -> %v", i, currentPath, validPath, match)
			if match {
				isValid = true
				log.Printf("[DEBUG] MATCH FOUND!")
				break
			}
		}

		log.Printf("[DEBUG] isValid: %v", isValid)

		if !isValid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. Получить токен из заголовка Authorization
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
    		return	
		}
		parts := strings.SplitN(authToken, " ", 2)
       
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Bad Authorization header", http.StatusUnauthorized)
    		return
		}
		 // 2. Валидировать его
		token := parts[1]
		authorized := processToken(token)
        // 3. Получить пользователя и его права
        // 4. Проверить, есть ли у пользователя право на этот метод и путь
        //    Например, для POST /files нужно право can_upload_files

        // Если проверка провалилась:
        if !authorized {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Если всё хорошо, передаем управление дальше
        next.ServeHTTP(w, r)
    })
}

func processToken(token string) bool{
	return true
}