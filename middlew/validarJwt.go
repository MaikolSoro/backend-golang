package middlew

import (
	"net/http"

	"github.com/maikol/backend-golang/routers"
)

/*ValidarJWT permite validar el JWT que nos viene en la petición */
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token !"+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
