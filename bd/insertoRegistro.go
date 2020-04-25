package bd

import (
	"context"
	"time"

	"github.com/maikol/backend-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarRegistro es la parada final con la BD, para insertar los datos del usuario */
func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != null {
		return "", false, err
	}

	ObjID, _ := result.InsertID.(primitive.ObjID)

	return ObjID.String(), true, nil

}
