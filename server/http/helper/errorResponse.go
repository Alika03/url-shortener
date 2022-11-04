package helper

import (
	"log"
	"net/http"
)

func ErrorResponse(err error, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		log.Println(err)
		return
	}
	return
}
