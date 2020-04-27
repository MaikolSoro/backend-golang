package routers

import (
	"net/http"

	"github.com/maikol/backend-golang/bd"
)

/*EliminarTweet permite borrar un tweet determinado */
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al  intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
