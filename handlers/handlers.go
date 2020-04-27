package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/maikol/backend-golang/middlew"
	"github.com/maikol/backend-golang/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	router := mux.NewRouter()

	// USUARIO
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	// AVATAR
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	// Banner

	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	// Tweet

	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/obtenerTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	// RELACION
	router.HandleFunc("/insertoRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.InsertoRelacion))).Methods("POST")
	router.HandleFunc("/deleteRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarRelacion))).Methods("DELETE")
	router.HandleFunc("/obtenerRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerRelacion))).Methods("GET")

	// LISTAR USUARIOS
	router.HandleFunc("/listarUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	//	MOSTRAR TWEETS SEGUIDORES
	router.HandleFunc("/obtenerTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.MostrarTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
