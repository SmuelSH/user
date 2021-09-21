package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"userm/userinsert"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/denisenkom/go-mssqldb"
)

var port = ":3000"

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))

	})

	r.Post("/getData", getData())

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println("Serving on " + localAddr.IP.String() + port)

	http.ListenAndServe(port, r)
}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)
		var user userinsert.UserData

		//fmt.Println(request)
		user.Nombres = request["Nombres"]
		user.Apellidos = request["Apellidos"]
		user.DocIden = request["DocIden"]
		user.Correo = request["Correo"]
		user.PrefijoNumCel = request["PrefijoNumCel"]
		user.NumCelular = request["NumCelular"]
		user.TipoLogin = request["TipoLogin"]
		user.Password = request["Password"]
		user.PreguntaSecreta = request["PreguntaSecreta"]
		user.RespuestaSecreta = request["RespuestaSecreta"]

		userinsert.UserInsert(user)

		w.Write([]byte("Registro Correcto....!"))
	}
}

//go mod init user
//go mod tidy
