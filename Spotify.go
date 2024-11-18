package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

// GetTokenFormat is local
type GetTokenFormat struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// https://developer.spotify.com/documentation/web-api/tutorials/client-credentials-flow
func getSpotifyToken() string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString("grant_type=client_credentials"))

	if err != nil {
		return "Error :("
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
  clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Origin", "")
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return ""
	}

	var data GetTokenFormat
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)

	return data.AccessToken
}

func spotifyGet(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+getSpotifyToken())

	return client.Do(req)
}

func getSpotifyUPCFromID(spotifyAlbumID string) (string, error) {
	requestURL := "https://api.spotify.com/v1/albums/" + spotifyAlbumID

	res, err := spotifyGet(requestURL)
	if err != nil || res.StatusCode != 200 {
		return "", errors.New("Could not find ISRC")
	}
	defer res.Body.Close()

	var data SpotifyAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	UPC := data.ExternalIds.Upc

	return UPC, nil
}

func getSpotifyISRCFromID(ID string) (string, error) {
	requestURL := "https://api.spotify.com/v1/tracks/" + ID

	res, err := spotifyGet(requestURL)
	if err != nil || res.StatusCode != 200 {
		return "", errors.New("Could not find ISRC")
	}
	defer res.Body.Close()

	var data SpotifyTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	ISRC := data.ExternalIds.Isrc

	return ISRC, nil
}

/*
	Searches for Spotify Album ID by album name and artist name in the Spotify API
*/
func (basicAlbum *BasicAlbum) getSpotifyID() string {
	requestURL := "https://api.spotify.com/v1/search/?type=album&q=" + basicAlbum.getQueryString()

	res, err := spotifyGet(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data SpotifySearchAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	ID := "N/A"

	if data.Albums.Total > 0 {
		ID = data.Albums.Items[0].ID
	}

	return ID
}

/*
  Get data from Spotify Web API
  Fetches ISRC and Spotify ID
*/
func (basicTrack *BasicTrack) getSpotifyID() string {
	requestURL := "https://api.spotify.com/v1/search/?type=track&q=" + basicTrack.getQueryString()

	res, err := spotifyGet(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data SpotifySearchTracks
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	ID := "N/A"

	if len(data.Tracks.Items) > 0 {
		track := data.Tracks.Items[0]
		ID = track.ID
	}

	return ID
}

func getBasicAlbumFromSpotifyAlbumID(ID string) (BasicAlbum, string, error) {
	requestURL := "https://api.spotify.com/v1/albums/" + ID

	res, err := spotifyGet(requestURL)
	if err != nil || res.StatusCode != 200 {
		return BasicAlbum{}, "N/A", errors.New("404")
	}
	defer res.Body.Close()

	var data SpotifyAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	basicAlbum := BasicAlbum{
		ArtistName: data.Artists[0].Name,
		AlbumName:  data.Name,
	}

	UpcID := data.ExternalIds.Upc

	return basicAlbum, UpcID, nil
}

func getBasicTrackFromSpotifyID(ID string) (BasicTrack, string, error) {
	requestURL := "https://api.spotify.com/v1/tracks/" + ID

	var ISRC string
	var basicTrack BasicTrack
	var potentialError error

	ISRC = "N/A"

	res, err := spotifyGet(requestURL)
	if err != nil || res.StatusCode != 200 {
		return basicTrack, ISRC, errors.New("404")
	}
	defer res.Body.Close()

	var data SpotifyTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err == nil {
		basicTrack = BasicTrack{
			TrackName:  data.Name,
			ArtistName: data.Artists[0].Name,
			AlbumName:  data.Album.Name,
		}

		ISRC = data.ExternalIds.Isrc
	} else {
		potentialError = errors.New("Could not get basic track from Spotify ID")
	}

	return basicTrack, ISRC, potentialError
}

func getSpotifyAlbumIDFromUPC(UPC string) (string, error) {
	requestURL := "https://api.spotify.com/v1/search?q=upc:" + UPC + "&type=album"

	var albumID string

	res, err := spotifyGet(requestURL)
	if err != nil || res.StatusCode != 200 {
		return albumID, errors.New("Spotify API did not response with 200")
	}
	defer res.Body.Close()

	var data SpotifyUpcAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	var potentialError error

	if err == nil && data.Albums.Total > 0 {
		albumID = data.Albums.Items[0].ID
	} else {
		potentialError = errors.New("Spotify could not find album by UPC code")
	}

	return albumID, potentialError
}
