package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

/*
  Get data from iTunes Search API
  Fetches iTunes ID
*/
func (basicTrack *BasicTrack) getItunesID() string {
	requestURL := "https://itunes.apple.com/search?term=" + basicTrack.getQueryString() + "&media=music&entity=musicTrack"

	// log.Println(requestURL)

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data ItunesSearchTracks
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	var ID = "N/A"

	if len(data.Results) > 0 {
		track := data.Results[0]
		ID = strconv.Itoa(track.TrackID)
	}

	return ID
}

func (basicAlbum *BasicAlbum) getItunesID() string {
	requestURL := "https://itunes.apple.com/search?term=" + basicAlbum.getQueryString() + "&media=music&entity=musicTrack"

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data ItunesSearchAlbums
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	var albumID = "N/A"

	if data.ResultCount > 0 {
		albumID = strconv.Itoa(data.Results[0].CollectionID)
	}

	return albumID
}

func getBasicTrackFromItunesID(ID string) (BasicTrack, error) {
	requestURL := "https://itunes.apple.com/lookup?id=" + ID + "&media=music&entity=musicTrack"

	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// log.Println(requestURL)

	var data ItunesTrack
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	name, artist, album := "", "", ""

	if len(data.Results) > 0 {
		track := data.Results[0]
		name = track.Trackname
		artist = track.Artistname
		album = track.Collectionname
		err = nil
	} else {
		err = errors.New("404")
	}

	// log.Println("Getting data from iTunes")

	basicTrack := BasicTrack{
		TrackName:  name,
		ArtistName: artist,
		AlbumName:  album,
	}

	return basicTrack, err
}

func getBasicAlbumFromItunesAlbumID(albumID string) BasicAlbum {
	requestURL := "https://itunes.apple.com/lookup?id=" + albumID + "&media=music&entity=album"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var basicAlbum BasicAlbum

	var data ItunesAlbum
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)

	if err == nil && data.ResultCount > 0 {
		basicAlbum = BasicAlbum{
			AlbumName:  data.Results[0].CollectionName,
			ArtistName: data.Results[0].ArtistName,
		}
	}

	return basicAlbum
}

func getItunesAlbumIDFromUPC(UPC string) (string, error) {
	requestURL := "https://itunes.apple.com/lookup?upc=" + UPC

	var albumID string

	res, err := http.Get(requestURL)
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
