package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		data = map[string]bool{
			"success": true,
		}
	}

	content, _ := json.Marshal(data)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	_, _ = w.Write(content)
}

func Index(w http.ResponseWriter, r *http.Request) {
	content := map[string]bool{
		"success": true,
	}

	Write(w, http.StatusOK, content)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	content := map[string]interface{}{
		"success": false,
		"message": http.StatusText(http.StatusNotFound),
	}
	Write(w, http.StatusNotFound, content)
}

func BadRequest(w http.ResponseWriter, message string) {
	content := map[string]interface{}{
		"success": false,
		"message": message,
	}
	Write(w, http.StatusBadRequest, content)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	content := map[string]interface{}{
		"success": false,
		"message": http.StatusText(http.StatusMethodNotAllowed),
	}
	Write(w, http.StatusMethodNotAllowed, content)
}
