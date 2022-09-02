package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = connectionDB()

var clientOption = options.Client().ApplyURI("mongodb+srv://brayandev:brayandev123@daily.0gtjrhw.mongodb.net/test")

//realizamos la connexion a base de datos
func connectionDB() *mongo.Client{
	client, err:=mongo.Connect(context.TODO(), clientOption)
	if err !=nil{
		log.Fatal(err.Error())
		return client
	}
	
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client	
	}

	log.Print("conexion exitosa a la db")
	return client
}


//revisa que la connection sea correcta
func CheckConnection() int{

	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil{
		return 0
	}
	return 1
}