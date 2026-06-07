package handlers


import(
	"net/http"
	"encoding/json"
	"strings"
)

// areParamsPresent проверяет наличие указанных параметров в URL.
// Она принимает запрос и список имен параметров, которые нужно проверить.
// Возвращает true, если ВСЕ указанные параметры присутствуют и не пусты.
func AreParamsPresent(r *http.Request, params ...string) bool {
	for _, param := range params {
		if r.URL.Query().Get(param) == "" {
			return false
		}
	}
	return true
}

func ArePostParamsPresent(r *http.Request, params ...string) (bool, string) {

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		msg := "Invalid content type, expected application/json"
		return false, msg
	}

	var data map[string]any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		msg := "Bad request:" + err.Error()
		return false, msg
	}

	for _, param := range params {
		if _, exists := data[param]; !exists {
			return false, "Missing required parameter:" + param
		}
	}
	return true, "OK"
}

func SuccessResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}