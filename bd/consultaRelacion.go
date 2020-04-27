package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/maikol/backend-golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*MostrarRelacion consulta la relacion entre 2 usuarios */
func MostrarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":  t.UsuarioID,
		"relacionid": t.RelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
