package main

import (
	"errors"
	"net/url"
)

// BasicAlbum : Minimum viable options needed to make a track string search
type BasicAlbum struct {
	ArtistName string
	AlbumName  string
}

func (basicAlbum *BasicAlbum) getQueryString() string {
	artistName := cleanStringForSearch(basicAlbum.ArtistName)
	albumName := cleanStringForSearch(basicAlbum.AlbumName)

	finalString := url.QueryEscape(artistName + " " + albumName)

	return finalString
}

func (basicAlbum *BasicAlbum) validate() error {
	if len(basicAlbum.AlbumName) == 0 && len(basicAlbum.ArtistName) == 0 {
		return errors.New("404")
	}

	return nil
}
