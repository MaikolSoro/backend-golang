package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/maikol/backend-golang/bd"
)

/*GetTweet leo o muestro los tweets */
func GetTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // Atoi significa la conversion string a un entero
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagina)
	respuesta, correcto := bd.ObtenerTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
