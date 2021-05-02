package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	s.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Can't connect to db now")
		log.Fatal("The error is: ", err)
	} else {
		fmt.Printf("We are connected to the %s database \n", Dbdriver)
	}

}

func (s *Server) Start(addr string, r *mux.Router) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, r))
}
