package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"gopkg.in/mgo.v2"
)

func main() {
	databaseURL := os.Getenv("MONGO_URL")
	session, err := mgo.Dial(databaseURL)

	log.Println("Spotify token is: " + getSpotifyToken())

	log.Println("Connecting to database on: " + databaseURL)

	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	databaseName := os.Getenv("DB_NAME")
	appDatabase := DataStore{session.DB(databaseName)}

	log.Println("Using database: " + databaseName)

	commonHandlers := alice.New(context.ClearHandler, LoggingHandler, RecoverHandler)

	router := NewRouter()
	router.Get("/track/:service/:ID", commonHandlers.ThenFunc(appDatabase.ServiceTrackHandler))
	router.Get("/album/:service/:ID", commonHandlers.ThenFunc(appDatabase.ServiceAlbumHandler))

	// port := os.Getenv("PORT")
	port := "5000"
	err = http.ListenAndServe(":"+port, router)

	log.Println("API is up and running!")

	if err != nil {
		log.Fatal(err)
	}
}
