package routers

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/rodrigopereza/twitter-go/bd"
	"github.com/rodrigopereza/twitter-go/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	erro := json.NewDecoder(r.Body).Decode(&mensaje)
	if erro != nil {
		http.Error(w, "Error en los datos recibidos: "+erro.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
