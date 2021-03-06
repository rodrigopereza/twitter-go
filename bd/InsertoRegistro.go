package db

import (
	"context"
	"time"

	"github.com/rodrigopereza/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil

}
