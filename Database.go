package main

import (
	"gopkg.in/mgo.v2"
)

// DataStore : A wrapper around the mgo instance
type DataStore struct {
	db *mgo.Database
}

// TrackCollection : An mgo Collection
type TrackCollection struct {
	collection *mgo.Collection
}

// AlbumCollection : An mgo Collection
type AlbumCollection struct {
	collection *mgo.Collection
}
