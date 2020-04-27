package bd

import (
	"context"
	"time"

	"github.com/maikol/backend-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ObtenerTweets lee los tweets de un perfil */
func ObtenerTweets(ID string, pagina int64) ([]*models.ObtenerTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.ObtenerTweets

	condicion := bson.M{
		"userid": ID,
	}
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // set sort para ordenerlos
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}
	for cursor.Next(context.TODO()) {
		var registro models.ObtenerTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
