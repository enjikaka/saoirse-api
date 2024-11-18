package main

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Album : The structure for tracks in the Saoirse API
type Album struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Artist    string        `json:"artist" bson:"artist"`
	SpotifyID string        `json:"spotify_id" bson:"spotify_id"`
	TidalID   string        `json:"tidal_id" bson:"tidal_id"`
	DeezerID  string        `json:"deezer_id" bson:"deezer_id"`
	ItunesID  string        `json:"itunes_id" bson:"itunes_id"`
	UpcID     string        `json:"upc_id" bson:"upc_id"`
}

/*
  Inserts a track to the database
*/
func (albums *AlbumCollection) insertAlbum(album Album) {
	foundAlbum, err := albums.fetchAlbum(album.SpotifyID, "spotify")

	if len(foundAlbum.SpotifyID) != len(album.SpotifyID) && foundAlbum.SpotifyID != album.SpotifyID && err != nil {
		albums.collection.Insert(album)
		// log.Println("Added " + album.Name + " by " + album.Artist)
	} else {
		// log.Println("Already have " + foundAlbum.Name + " by " + foundAlbum.Artist)
	}
}

func (albums *AlbumCollection) createNewFromItunesID(ItunesID string) (Album, error) {
	basicAlbum := getBasicAlbumFromItunesAlbumID(ItunesID)

	if len(basicAlbum.AlbumName) == 0 && len(basicAlbum.ArtistName) == 0 {
		return Album{}, errors.New("404")
	}

	foundAlbum, err := albums.fetchAlbum(ItunesID, "itunes")

	if len(foundAlbum.ItunesID) == len(ItunesID) && foundAlbum.ItunesID == ItunesID && err != nil {
		return foundAlbum, nil
	}

	SpotifyID := basicAlbum.getSpotifyID()
	UpcID, upcError := getSpotifyUPCFromID(SpotifyID)

	var TidalID string
	var DeezerID string

	if upcError == nil {
		var tidalError error
		var deezerError error

		TidalID, tidalError = getTidalAlbumIDFromUPC(UpcID)
		DeezerID, deezerError = getDeezerAlbumIDFromUPC(UpcID)

		if tidalError != nil {
			TidalID = basicAlbum.getTidalID()
		}

		if deezerError != nil {
			DeezerID = basicAlbum.getDeezerID()
		}
	}

	album := Album{
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicAlbum.AlbumName,
		Artist:    basicAlbum.ArtistName,
		UpcID:     UpcID,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	albums.insertAlbum(album)

	return album, nil
}

func (albums *AlbumCollection) createNewFromSpotifyID(SpotifyID string) (Album, error) {
	basicAlbum, UpcID, err := getBasicAlbumFromSpotifyAlbumID(SpotifyID)

	if len(basicAlbum.AlbumName) == 0 && len(basicAlbum.ArtistName) == 0 {
		return Album{}, errors.New("404")
	}

	foundAlbum, err := albums.fetchAlbum(SpotifyID, "spotify")

	if len(foundAlbum.SpotifyID) == len(SpotifyID) && foundAlbum.SpotifyID == SpotifyID && err != nil {
		return foundAlbum, nil
	}

	TidalID, tidalError := getTidalAlbumIDFromUPC(UpcID)
	DeezerID, deezerError := getDeezerAlbumIDFromUPC(UpcID)
	ItunesID, itunesError := getItunesAlbumIDFromUPC(UpcID)

	if tidalError != nil {
		TidalID = basicAlbum.getTidalID()
	}

	if deezerError != nil {
		DeezerID = basicAlbum.getDeezerID()
	}

	if itunesError != nil {
		ItunesID = basicAlbum.getItunesID()
	}

	album := Album{
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicAlbum.AlbumName,
		Artist:    basicAlbum.ArtistName,
		UpcID:     UpcID,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	albums.insertAlbum(album)

	return album, nil
}

func (albums *AlbumCollection) createNewFromDeezerID(deezerID string) (Album, error) {
	basicAlbum, UpcID := getBasicAlbumFromDeezerAlbumID(deezerID)

	if len(basicAlbum.AlbumName) == 0 && len(basicAlbum.ArtistName) == 0 {
		return Album{}, errors.New("404")
	}

	foundAlbum, err := albums.fetchAlbum(deezerID, "deezer")

	if len(foundAlbum.DeezerID) == len(deezerID) && foundAlbum.DeezerID == deezerID && err != nil {
		return foundAlbum, nil
	}

	TidalID, tidalError := getTidalAlbumIDFromUPC(UpcID)
	SpotifyID, spotifyError := getSpotifyAlbumIDFromUPC(UpcID)
	ItunesID, itunesError := getItunesAlbumIDFromUPC(UpcID)

	if tidalError != nil {
		TidalID = basicAlbum.getTidalID()
	}

	if spotifyError != nil {
		SpotifyID = basicAlbum.getSpotifyID()
	}

	if itunesError != nil {
		ItunesID = basicAlbum.getItunesID()
	}

	album := Album{
		DeezerID:  deezerID,
		ID:        bson.NewObjectId(),
		Name:      basicAlbum.AlbumName,
		Artist:    basicAlbum.ArtistName,
		UpcID:     UpcID,
		TidalID:   TidalID,
		SpotifyID: SpotifyID,
		ItunesID:  ItunesID,
	}

	albums.insertAlbum(album)

	return album, nil
}

func (albums *AlbumCollection) createNewFromTidalID(TidalID string) (Album, error) {
	basicAlbum, UpcID := getBasicAlbumFromTidalAlbumID(TidalID)

	validationError := basicAlbum.validate()

	if validationError != nil {
		return Album{}, validationError
	}

	// Try to find album in local database
	foundAlbum, err := albums.fetchAlbum(TidalID, "tidal")

	// If id matches with DB, return DB values (else we update the records)
	if len(foundAlbum.TidalID) == len(TidalID) && foundAlbum.TidalID == TidalID && err != nil {
		return foundAlbum, nil
	}

	SpotifyID, spotifyError := getSpotifyAlbumIDFromUPC(UpcID)
	DeezerID, deezerError := getDeezerAlbumIDFromUPC(UpcID)
	ItunesID, itunesError := getItunesAlbumIDFromUPC(UpcID)

	// log.Println("DeezerID")
	// log.Println(DeezerID)
	// log.Println("deezerError")
	// log.Println(deezerError)

	if spotifyError != nil {
		SpotifyID = basicAlbum.getSpotifyID()
	}

	if deezerError != nil {
		DeezerID = basicAlbum.getDeezerID()
	}

	if itunesError != nil {
		ItunesID = basicAlbum.getItunesID()
	}

	album := Album{
		ID:        bson.NewObjectId(),
		Name:      basicAlbum.AlbumName,
		Artist:    basicAlbum.ArtistName,
		UpcID:     UpcID,
		TidalID:   TidalID,
		SpotifyID: SpotifyID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	albums.insertAlbum(album)

	return album, nil
}

func (albums *AlbumCollection) fetchAlbum(ID string, service string) (Album, error) {
	result := Album{}
	service += "_id"

	err := albums.collection.Find(bson.M{service: ID}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
