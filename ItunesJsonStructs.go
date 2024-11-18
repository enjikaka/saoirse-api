package main

// ItunesAlbum : JSON
type ItunesAlbum struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		ArtistName     string `json:"artistName"`
		CollectionName string `json:"collectionName"`
	} `json:"results"`
}

// ItunesTrack : The structure of the JSON in the response from iTunes API when requesting a track
type ItunesTrack struct {
	Resultcount int `json:"resultCount"`
	Results     []struct {
		Artistname             string `json:"artistName"`
		Collectionname         string `json:"collectionName"`
		Trackname              string `json:"trackName"`
		Collectioncensoredname string `json:"collectionCensoredName"`
	} `json:"results"`
}

// ItunesSearchTracks : The structure of the JSON in the response from iTunes API when searching
type ItunesSearchTracks struct {
	Resultcount int `json:"resultCount"`
	Results     []struct {
		TrackID int `json:"trackId"`
	} `json:"results"`
}

// ItunesSearchAlbums : JSON
type ItunesSearchAlbums struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		CollectionID int `json:"collectionId"`
	} `json:"results"`
}
