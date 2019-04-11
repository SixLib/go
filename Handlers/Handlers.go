package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/sixlib/go/Models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `welcome!`)
}
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	todos := models.Todos{
		models.Todo{ID: 1, NAME: "TOM"},
		models.Todo{ID: 2, NAME: "JERRY"},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
func TodoOfKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "todo key：%v", vars["key"])
}
