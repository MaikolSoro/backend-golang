package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/maikol/backend-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ListarUsuarios  lee usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacionan conmigo */
func ListarUsuarios(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetLimit((page - 1) * 20)
	findOptions.SetSkip(20)
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.RelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = MostrarRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.RelacionID == ID {
			incluir = false
		}
		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true

}
