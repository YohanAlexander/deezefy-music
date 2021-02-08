package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// Recovery adiciona logging dos server panics
func Recovery(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err) // May be log this error? Send to sentry?
			jsonBody, _ := json.Marshal(map[string]string{
				"error": "There was an internal server error",
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonBody)
		}
	}()
	next(w, r)
}
