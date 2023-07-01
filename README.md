# Gorilla/Mux with HTTP

This is a simple app with basic HTTP route handling


```
router.HandleFunc("/movies", getMovies).Methods("GET")
router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
router.HandleFunc("/movies", createMovie).Methods("POST")
router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```