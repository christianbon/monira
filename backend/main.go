package main

import (
	database "backend/database"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "docker123#"
	DB_NAME     = "postgres-db"
)

func main() {
	database.InitDB()

	// Init the mux router
	// router := mux.NewRouter()
	// setupDB()

	// router.HandleFunc("/movies/", GetMovies).Methods("GET")
	// router.HandleFunc("/movies/", CreateMovie).Methods("POST")
	// router.HandleFunc("/movies/{movieid}", DeleteMovie).Methods("DELETE")
	// router.HandleFunc("/movies/", DeleteMovies).Methods("DELETE")

	// // serve the app
	// fmt.Println("Server at 8080")
	// log.Fatal(http.ListenAndServe(":8000", router))
}
