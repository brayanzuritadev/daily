package db

import (
	"context"
	"github.com/brayanzuritadev/daily/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ReviewUser(usuario models.Usuario) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("Daily")
	col := db.Collection("user")

	condition := bson.M{"email": usuario.Email}

	err := col.FindOne(ctx, condition).Decode(&usuario)

	ID := usuario.ID.Hex()

	if err != nil {
		return usuario, false, "error al buscar al usuario"
	}

	return usuario, false, ID
}
