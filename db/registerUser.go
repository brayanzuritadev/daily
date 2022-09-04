package db

import (
	"context"
	"github.com/brayanzuritadev/daily/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func RegisterUser(u models.Usuario) (bool, error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("Daily")
	col := db.Collection("user")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return false, err, "No se logro registrar el usuario"
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return true, nil, ObjID.String()
}
