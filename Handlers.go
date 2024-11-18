package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// IndexHandler : Handles the index route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// LoggingHandler : Handles loggin HTTP information
func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v from %s \n", r.Method, r.URL.String(), t2.Sub(t1), r.RemoteAddr)
	}

	return http.HandlerFunc(fn)
}

// RecoverHandler : Handles potensial errors with 500 response code
func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// ServiceAlbumHandler : Handles the /album route
func (ds *DataStore) ServiceAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params) // Get the variable from request URL
	service := params.ByName("service")
	ID := params.ByName("ID")

	albums := AlbumCollection{ds.db.C("albums")}
	album, err := albums.fetchAlbum(ID, service)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		switch service {
		case "deezer":
			album, err = albums.createNewFromDeezerID(ID)
		// case "isrc":
		//	album, err = albums.createNewAlbumFromDeezerWithISRC(ID)
		case "spotify":
			album, err = albums.createNewFromSpotifyID(ID)
		case "tidal":
			album, err = albums.createNewFromTidalID(ID)
		case "itunes":
			album, err = albums.createNewFromItunesID(ID)
		}

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
	}

	json.NewEncoder(w).Encode(album)
}

// ServiceTrackHandler : Handles the /track route
func (ds *DataStore) ServiceTrackHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params) // Get the variable from request URL
	service := params.ByName("service")
	ID := params.ByName("ID")

	tracks := TrackCollection{ds.db.C("tracks")}
	track, err := tracks.fetchTrack(ID, service)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Origin", "saoirse.audio")
	// w.Header().Add("Access-Control-Allow-Origin", "chrome-extension://mapfnomankohneoccopodeldnialinlp")

	if err != nil {
		switch service {
		case "deezer":
			track, err = tracks.createNewTrackFromDeezerID(ID)
		case "isrc":
			track, err = tracks.createNewTrackFromDeezerWithISRC(ID)
		case "spotify":
			track, err = tracks.createNewTrackFromSpotifyID(ID)
		case "tidal":
			track, err = tracks.createNewTrackFromTidalID(ID)
		case "itunes":
			track, err = tracks.createNewTrackFromItunesID(ID)
		}

		if err != nil {
			panic(err)
			http.Error(w, http.StatusText(404), 404)
			return
		}
	}

	json.NewEncoder(w).Encode(track)
}
