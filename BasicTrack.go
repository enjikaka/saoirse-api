package main

import (
	"errors"
	"net/url"
)

// BasicTrack : Minimum viable options needed to make a track string search
type BasicTrack struct {
	TrackName  string
	ArtistName string
	AlbumName  string
}

// TODO make one with album too as fallback
func (basicTrack *BasicTrack) getQueryString() string {
	trackName := cleanStringForSearch(basicTrack.TrackName)
	artistName := cleanStringForSearch(basicTrack.ArtistName)
	// albumName := cleanStringForSearch(basicTrack.AlbumName)

	// Prevent too much redundant information (Helps with "Single"-albums for example)
	/*
		if strings.Compare(albumName, trackName) > -1 {
			albumName = ""
		}
	*/
	finalString := url.QueryEscape(trackName + " " + artistName)

	// log.Println(finalString)

	return finalString
}

func (basicTrack *BasicTrack) validate() error {
	if len(basicTrack.TrackName) == 0 && len(basicTrack.ArtistName) == 0 {
		return errors.New("404")
	}

	return nil
}
