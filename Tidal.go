package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO: Refactor to TIDAL new Open API! https://developer.tidal.com/

// TidalToken : Token for TIDAL API calls.
var TidalToken = os.Getenv("TIDAL_TOKEN")

func (basicTrack *BasicTrack) getTidalID() string {
	requestURL := "http://api.tidal.com/v1/search?query=" + basicTrack.getQueryString() + "&limit=1&offset=0&types=TRACKS&countryCode=NO"

	log.Println(requestURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data TidalTrackSearch
	err = decoder.Decode(&data)

	var ID = "N/A"

	if len(data.Tracks.Items) > 0 {
		track := data.Tracks.Items[0]
		ID = strconv.Itoa(track.ID)
	}

	return ID
}

func (basicAlbum *BasicAlbum) getTidalID() string {
	requestURL := "http://api.tidal.com/v1/search?query=" + basicAlbum.getQueryString() + "&limit=1&offset=0&types=ALBUMS&countryCode=NO"

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data TidalAlbumSearch
	err = decoder.Decode(&data)

	var albumID = "N/A"

	if len(data.Albums.Items) > 0 {
		albumID = strconv.Itoa(data.Albums.Items[0].ID)
	}

	return albumID
}

/*
  Get data from TIDAL API
  Fetches TIDAL ID
*/
func getTidalIDFromISRC(ISRC string) (string, error) {
	requestURL := "http://api.tidal.com/v1/tracks/?isrc=" + ISRC + "&countryCode=NO"

	log.Println(requestURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data TidalIsrcTracks
	err = decoder.Decode(&data)

	var tidalID string
	var potentialError error

	if data.TotalNumberOfItems > 0 {
		tidalID = strconv.Itoa(data.Items[0].ID)
		potentialError = nil
	} else {
		potentialError = errors.New("Could not find track in TIDAL by ISRC")
	}

	return tidalID, potentialError
}

/*
  Get data from TIDAL API
  Fetches TIDAL Album ID
*/
func getTidalAlbumIDFromUPC(UPC string) (string, error) {
	requestURL := "http://api.tidal.com/v1/albums/?upc=" + UPC + "&countryCode=NO"

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data TidalUpcAlbum
	err = decoder.Decode(&data)

	var albumID = "N/A"
	var potentialError error

	if data.TotalNumberOfItems > 0 {
		albumID = strconv.Itoa(data.Items[0].ID)
	} else {
		potentialError = errors.New("Could not find album id from TIDAL via UPC")
	}

	return albumID, potentialError
}

func getArtistNames(artists TidalArtists) string {
	var names = []string{}

	for _, artist := range artists {
		names = append(names, artist.Name)
	}

	return strings.Join(names, ",")
}

func getBasicTrackFromTidalID(ID string) BasicTrack {
	requestURL := "http://api.tidal.com/v1/tracks/" + ID + "?countryCode=NO"

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var track TidalTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&track)

	basicTrack := BasicTrack{
		TrackName:  track.Title,
		ArtistName: getArtistNames(track.Artists),
		AlbumName:  track.Album.Title,
	}

	return basicTrack
}

func getBasicAlbumFromTidalAlbumID(TidalAlbumID string) (BasicAlbum, string) {
	requestURL := "http://api.tidal.com/v1/albums/" + TidalAlbumID + "?countryCode=NO"

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Add("X-Tidal-Token", TidalToken)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var album TidalAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&album)

	UpcID := album.Upc
	basicAlbum := BasicAlbum{
		ArtistName: getArtistNames(album.Artists),
		AlbumName:  album.Title,
	}

	return basicAlbum, UpcID
}
