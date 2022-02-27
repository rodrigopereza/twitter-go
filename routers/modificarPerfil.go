package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/rodrigopereza/twitter-go/bd"
	"github.com/rodrigopereza/twitter-go/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
