package handler

import (
	"encoding/json"
	// "github.com/gorilla/mux"
	"learn-golang/pkg/data"
	// "learn-golang/pkg/dto"
	"net/http"
	// "strconv"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	responseWithJson(writer, http.StatusOK, data.Todos)
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
