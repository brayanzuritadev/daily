package db

import (
	"context"
	"github.com/brayanzuritadev/daily/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ReviewUser(usuario models.User) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("Daily")
	col := db.Collection("user")

	condition := bson.M{"email": usuario.Email}

	err := col.FindOne(ctx, condition).Decode(&usuario)

	ID := usuario.ID.Hex()

	if err != nil {
		return usuario, false, ID
	}

	return usuario, true, ID
}
