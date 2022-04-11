package routers

import (
	"net/http"

	db "github.com/rodrigopereza/twitter-go/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Clone().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
