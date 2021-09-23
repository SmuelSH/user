package userinsert

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type UserData struct {
	Nombres          string
	Apellidos        string
	DocIden          string
	Correo           string
	PrefijoNumCel    string
	NumCelular       string
	TipoLogin        string
	Password         string
	PreguntaSecreta  string
	RespuestaSecreta string
}

var server = "(local)"

//var port = 1433
var user = "admin"
var password = "Temporal1"

var drivermssql = "mssql"
var database = "PMMApps_NB"

func dbcnn() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)
	db, errdb := sql.Open(drivermssql, connString)

	defer handleOutOfBounds()

	if errdb != nil {
		log.Fatal("Error open db:" + errdb.Error())
	}

	rst, errdb := db.Query("select 1")
	if errdb != nil {

		//log.Fatal("error abrir rst:" + errdb.Error())
		panic("Out of bound access for slice")
	}

	defer rst.Close()

	log.Println("BD conected")

	return db, errdb
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func UserInsert(usrd UserData) {
	var sntran string
	var sDescr string

	cnn, err := dbcnn()
	defer cnn.Close()
	if err == nil {
		rows, err := cnn.Query("pWsClientes ?,?,?,?,?,?,?,?,?,?", usrd.Nombres, usrd.Apellidos, usrd.DocIden, usrd.Correo, usrd.PrefijoNumCel, usrd.NumCelular, usrd.TipoLogin, usrd.Password, usrd.PreguntaSecreta, usrd.RespuestaSecreta)
		defer rows.Close()
		if err != nil {

			log.Fatal("error abrir:" + err.Error())

		}

		for rows.Next() {
			err := rows.Scan(&sntran, &sDescr)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(sntran, sDescr)
		}

	} else {

		fmt.Println("Problemas de Conexion:" + err.Error())

	}

}
