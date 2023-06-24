package handlers

import (
	"API/model"
	"API/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var response Response

func TodoRouter(router chi.Router) {
	router.Get("/{id}", Get)
	router.Get("/", GetAll)
	router.Post("/", Create)
	router.Put("/{id}", Update)
	router.Delete("/{id}", Delete)
}

func Create(w http.ResponseWriter, request *http.Request) {
	var todo model.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error parsing %v", request.Body)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := services.Insert(todo)
	if err != nil {
		response = NewResponse(http.StatusInternalServerError, err.Error())
	} else {
		response = NewResponse(http.StatusCreated, fmt.Sprintf("Todo adicionado com sucesso com id: %v", id))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Update(w http.ResponseWriter, request *http.Request) {
	var todo model.Todo
	id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
	if err != nil {
		log.Fatal("Id Not Found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error parsing %v", request.Body)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err = services.Update(id, todo)
	if err != nil {
		response = NewResponse(http.StatusInternalServerError, err.Error())
	} else {
		response = NewResponse(http.StatusCreated, fmt.Sprintf("Todo atualizado com sucesso com id: %v", id))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Delete(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
	if err != nil {
		log.Fatal("Id Not Found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	id, err = services.Delete(id)
	if err != nil {
		response = NewResponse(http.StatusInternalServerError, err.Error())
	} else {
		response = NewResponse(http.StatusCreated, fmt.Sprintf("Todo deletado com sucesso com id: %v", id))
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func Get(w http.ResponseWriter, request *http.Request) {
	var todo model.Todo

	id, err := strconv.ParseInt(chi.URLParam(request, "id"), 10, 64)
	if err != nil {
		log.Fatal("Id Not Found")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	todo, err = services.Get(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

func GetAll(w http.ResponseWriter, request *http.Request) {
	todos, err := services.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
