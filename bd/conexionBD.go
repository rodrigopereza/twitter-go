package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@twitterdb.eavfm.mongodb.net/twitter?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar a la DB*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa a la base de datos")
	return client
}

func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}
