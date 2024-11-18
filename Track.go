package main

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Track : The structure for tracks in the Saoirse API
type Track struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Artist    string        `json:"artist" bson:"artist"`
	SpotifyID string        `json:"spotify_id" bson:"spotify_id"`
	TidalID   string        `json:"tidal_id" bson:"tidal_id"`
	DeezerID  string        `json:"deezer_id" bson:"deezer_id"`
	ItunesID  string        `json:"itunes_id" bson:"itunes_id"`
	YouTubeID string        `json:"youtube_id,omitempty" bson:"youtube_id"`
	IsrcID    string        `json:"isrc_id" bson:"isrc_id"`
}

/*
  Inserts a track to the database
*/
func (tracks *TrackCollection) insertTrack(track Track) {
	findTrack, err := tracks.fetchTrack(track.SpotifyID, "spotify")

	if len(findTrack.SpotifyID) != len(track.SpotifyID) && findTrack.SpotifyID != track.SpotifyID && err != nil {
		tracks.collection.Insert(track)
		log.Println("Added " + track.Name + " by " + track.Artist)
	} else {
		log.Println("Already have " + findTrack.Name + " by " + findTrack.Artist)
	}
}

func (tracks *TrackCollection) createNewTrackFromTidalID(tidalID string) (Track, error) {
	basicTrack := getBasicTrackFromTidalID(tidalID)

	validateError := basicTrack.validate()

	if validateError != nil {
		return Track{}, validateError
	}

	findTrack, err := tracks.fetchTrack(tidalID, "tidal")

	if len(findTrack.TidalID) == len(tidalID) && findTrack.TidalID == tidalID && err != nil {
		return findTrack, nil
	}

	var DeezerID string

	SpotifyID := basicTrack.getSpotifyID()
	ISRC, err := getSpotifyISRCFromID(SpotifyID)
	ItunesID := basicTrack.getItunesID()

	if err != nil {
		DeezerID = getDeezerIDFromISRC(ISRC)
	} else {
		DeezerID = basicTrack.getDeezerID()
	}

	track := Track{
		ID:        bson.NewObjectId(),
		Name:      basicTrack.TrackName,
		Artist:    basicTrack.ArtistName,
		SpotifyID: SpotifyID,
		TidalID:   tidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
		IsrcID:    ISRC,
	}

	tracks.insertTrack(track)

	return track, nil
}

func (tracks *TrackCollection) createNewTrackFromDeezerWithISRC(ISRC string) (Track, error) {
	basicTrack, DeezerID := getBasicTrackAndDeezerIDFromISRC(ISRC)

	validateError := basicTrack.validate()

	if validateError != nil {
		return Track{}, validateError
	}

	findTrack, err := tracks.fetchTrack(DeezerID, "deezer")

	if len(findTrack.TidalID) == len(DeezerID) && findTrack.DeezerID == DeezerID && err != nil {
		return findTrack, nil
	}

	// fmt.Println(trackName)
	// fmt.Println(trackArtist)

	SpotifyID := basicTrack.getSpotifyID()
	ItunesID := basicTrack.getItunesID()

	TidalID, tidalError := getTidalIDFromISRC(ISRC)

	if tidalError != nil {
		TidalID = basicTrack.getTidalID()
	}

	// fmt.Println(SpotifyID)

	track := Track{
		IsrcID:    ISRC,
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicTrack.TrackName,
		Artist:    basicTrack.ArtistName,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	tracks.insertTrack(track)

	return track, nil
}

func (tracks *TrackCollection) createNewTrackFromDeezerID(DeezerID string) (Track, error) {
	basicTrack := getBasicTrackFromDeezerID(DeezerID)

	validateError := basicTrack.validate()

	if validateError != nil {
		return Track{}, validateError
	}

	findTrack, err := tracks.fetchTrack(DeezerID, "deezer")

	if len(findTrack.TidalID) == len(DeezerID) && findTrack.DeezerID == DeezerID && err != nil {
		return findTrack, nil
	}

	var TidalID string

	SpotifyID := basicTrack.getSpotifyID()
	ISRC, err := getSpotifyISRCFromID(SpotifyID)
	ItunesID := basicTrack.getSpotifyID()

	if err == nil {
		TidalID = basicTrack.getTidalID()
	} else {
		var tidalError error
		TidalID, tidalError = getTidalIDFromISRC(ISRC)

		if tidalError != nil {
			TidalID = basicTrack.getTidalID()
		}
	}

	track := Track{
		IsrcID:    ISRC,
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicTrack.TrackName,
		Artist:    basicTrack.ArtistName,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	tracks.insertTrack(track)

	return track, nil
}

func (tracks *TrackCollection) createNewTrackFromSpotifyID(SpotifyID string) (Track, error) {
	basicTrack, ISRC, err := getBasicTrackFromSpotifyID(SpotifyID)

	if len(basicTrack.TrackName) == 0 && len(basicTrack.ArtistName) == 0 {
		return Track{}, errors.New("404")
	}

	findTrack, err := tracks.fetchTrack(SpotifyID, "spotify")

	if len(findTrack.SpotifyID) == len(SpotifyID) && findTrack.SpotifyID == SpotifyID && err != nil {
		return findTrack, nil
	}

	TidalID, tidalError := getTidalIDFromISRC(ISRC)

	if tidalError != nil {
		TidalID = basicTrack.getTidalID()
	}

	DeezerID := getDeezerIDFromISRC(ISRC)
	ItunesID := basicTrack.getItunesID()

	track := Track{
		IsrcID:    ISRC,
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicTrack.TrackName,
		Artist:    basicTrack.ArtistName,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	tracks.insertTrack(track)

	return track, nil
}

func (tracks *TrackCollection) createNewTrackFromItunesID(ItunesID string) (Track, error) {
	basicTrack, err := getBasicTrackFromItunesID(ItunesID)

	validateError := basicTrack.validate()

	if validateError != nil {
		return Track{}, validateError
	}

	findTrack, err := tracks.fetchTrack(ItunesID, "itunes")

	if len(findTrack.ItunesID) == len(ItunesID) && findTrack.ItunesID == ItunesID && err != nil {
		return findTrack, nil
	}

	var TidalID string
	var DeezerID string

	SpotifyID := basicTrack.getSpotifyID()
	ISRC, err := getSpotifyISRCFromID(SpotifyID)

	if err == nil {
		TidalID = basicTrack.getTidalID()
		DeezerID = basicTrack.getDeezerID()
	} else {
		var tidalError error
		TidalID, tidalError = getTidalIDFromISRC(ISRC)
		DeezerID = getDeezerIDFromISRC(ISRC)

		if tidalError != nil {
			TidalID = basicTrack.getTidalID()
		}
	}

	if err != nil {
		return Track{}, errors.New("404")
	}

	track := Track{
		IsrcID:    ISRC,
		SpotifyID: SpotifyID,
		ID:        bson.NewObjectId(),
		Name:      basicTrack.TrackName,
		Artist:    basicTrack.ArtistName,
		TidalID:   TidalID,
		DeezerID:  DeezerID,
		ItunesID:  ItunesID,
	}

	tracks.insertTrack(track)

	return track, nil
}

func (tracks *TrackCollection) fetchTrack(ID string, service string) (Track, error) {
	result := Track{}
	service += "_id"

	err := tracks.collection.Find(bson.M{service: ID}).One(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
