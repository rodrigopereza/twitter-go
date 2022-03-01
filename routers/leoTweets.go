package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/rodrigopereza/twitter-go/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro Id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Error en el parametro pagina debe ser un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	fmt.Println("ID = " + ID)

	respuesta, correcto := db.LeoTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
