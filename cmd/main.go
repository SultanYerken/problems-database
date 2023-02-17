package main

import (
	"fmt"
	"log"

	"git.01.alem.school/Sultanye/problems-database/internal/controller"
	"git.01.alem.school/Sultanye/problems-database/internal/controller/httpserver"
	"git.01.alem.school/Sultanye/problems-database/internal/repository"
	"git.01.alem.school/Sultanye/problems-database/internal/usecase"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var PORT = "8080"

func main() {
	db, err := repository.NewPostrgresDB(repository.Config{
		Username: "postgres",
		Password: "qwerty",
		Host:     "db",
		Port:     "5432",
		DBName:   "postgres",
	})
	if err != nil {
		log.Println(err)
		return
	}

	repos := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repos)
	handler := controller.NewHandler(usecase)
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Home).Methods("GET")
	router.HandleFunc("/problem", handler.GetAllProblems).Methods("GET")
	router.HandleFunc("/problem/{id:[0-9]+}", handler.GetProblemById).Methods("GET")
	router.HandleFunc("/problem", handler.CreateProblem).Methods("POST")
	router.HandleFunc("/problem{id:[0-9]+}", handler.EditProblem).Methods("PUT")
	router.HandleFunc("/problem{id:[0-9]+}", handler.DeleteProblem).Methods("DELETE")

	srv := new(httpserver.Server)
	log.Println("Sever go http://localhost:8080")
	if err := srv.Run(PORT, router); err != nil {
		fmt.Println("main srv Run error")
		log.Println(err)
		return
	}
}

//  docker exec -it problems-database_db_1 createdb --username=postgres --owner=postgres testdb // создание БД в образе докера
// migrate -path db/migrations -database 'postgres://postgres:qwerty@localhost:5432/testdb?sslmode=disable' up // миграции в эту БД
// docker exec -it my_db_1 psql -U postgres -c "SELECT * FROM problem" // посмотреть в таблицу

// TODO first run write topics
