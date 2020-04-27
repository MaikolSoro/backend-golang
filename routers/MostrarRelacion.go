package routers

import (
	"encoding/json"
	"net/http"

	"github.com/maikol/backend-golang/bd"
	"github.com/maikol/backend-golang/models"
)

/*ObtenerRelacion chequea si hay relacion entre 2 usuarios */
func ObtenerRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.RelacionID = ID

	var resp models.MostrarRelacion

	status, err := bd.MostrarRelacion(t)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
