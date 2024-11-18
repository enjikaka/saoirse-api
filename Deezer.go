package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (basicTrack *BasicTrack) getDeezerID() string {
	requestURL := "https://api.deezer.com/search/track?q=" + basicTrack.getQueryString()

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data DeezerSearchTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	var trackID = "N/A"

	if err == nil && len(data.Data) > 0 {
		trackID = strconv.Itoa(data.Data[0].ID)
	}

	return trackID
}

func (basicAlbum *BasicAlbum) getDeezerID() string {
	requestURL := "https://api.deezer.com/search/album?q=" + basicAlbum.getQueryString()

	// log.Println(requestURL)

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data DeezerSearchAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	albumID := "N/A"

	if err == nil && len(data.Data) > 0 {
		albumID = strconv.Itoa(data.Data[0].ID)
	}

	return albumID
}

func getDeezerIDFromISRC(ISRC string) string {
	requestURL := "https://api.deezer.com/2.0/track/isrc:" + ISRC

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data DeezerTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	trackID := "N/A"

	if err == nil && data.ID != 0 {
		trackID = strconv.Itoa(data.ID)
	}

	return trackID
}

func getBasicTrackAndDeezerIDFromISRC(ISRC string) (BasicTrack, string) {
	query := "track/isrc:" + ISRC
	requestURL := "https://api.deezer.com/2.0/" + query

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data DeezerTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	trackID := "N/A"
	var basicTrack BasicTrack

	if err == nil && data.ID != 0 {
		trackID = strconv.Itoa(data.ID)
		basicTrack = BasicTrack{
			TrackName:  data.Title,
			ArtistName: data.Artist.Name,
			AlbumName:  data.Album.Title,
		}
	}

	return basicTrack, trackID
}

func getBasicTrackFromDeezerID(trackID string) BasicTrack {
	requestURL := "https://api.deezer.com/2.0/track/" + trackID

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var track DeezerTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&track)

	var basicTrack BasicTrack

	if err == nil {
		basicTrack = BasicTrack{
			TrackName:  track.Title,
			ArtistName: track.Artist.Name,
			AlbumName:  track.Album.Title,
		}
	}

	return basicTrack
}

func getBasicAlbumFromDeezerAlbumID(albumID string) (BasicAlbum, string) {
	requestURL := "https://api.deezer.com/2.0/album/" + albumID

	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var album DeezerAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&album)

	var UPC string
	var basicAlbum BasicAlbum

	if err == nil {
		UPC = album.Upc
		basicAlbum = BasicAlbum{
			AlbumName:  album.Title,
			ArtistName: album.Artist.Name,
		}
	}

	return basicAlbum, UPC
}

func getDeezerAlbumIDFromUPC(UPC string) (string, error) {
	requestURL := "http://api.deezer.com/album/upc:" + UPC

	// log.Println(requestURL)

	var albumID string

	res, err := http.Get(requestURL)
	if err != nil || res.StatusCode != 200 {
		return albumID, errors.New("Deezer API did not response with 200")
	}
	defer res.Body.Close()

	var data DeezerUPCAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	var potentialError error

	if err == nil && data.ID != 0 {
		albumID = strconv.Itoa(data.ID)
	} else {
		potentialError = errors.New("Deezer could not find album by UPC code")
	}

	return albumID, potentialError
}
