package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/maikol/backend-golang/bd"
)

/*MostrarTweetsSeguidores lee los tweets de todos nuestros seguidores */
func MostrarTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // lo convertimos a entero pagina

	if err != nil {
		http.Error(w, "Debe enviar el parametro página como entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.ObtenerTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
