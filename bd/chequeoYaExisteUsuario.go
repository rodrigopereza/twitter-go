package db

import (
	"context"
	"time"

	"github.com/rodrigopereza/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuario")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)
	//ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, err
	}

	return resultado, true, nil

}
