package routers

import (
	"net/http"

	"github.com/maikol/backend-golang/bd"
	"github.com/maikol/backend-golang/models"
)

/*EliminarRelacion realiza el borrado de la relacion entre usuarios */
func EliminarRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.RelacionID = ID

	status, err := bd.DeleteRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
