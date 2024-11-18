package main

// TidalUpcAlbum : JSON strcut when getting album from UPC
type TidalUpcAlbum struct {
	Limit              int `json:"limit"`
	Offset             int `json:"offset"`
	TotalNumberOfItems int `json:"totalNumberOfItems"`
	Items              []struct {
		ID                   int         `json:"id"`
		Title                string      `json:"title"`
		Duration             int         `json:"duration"`
		StreamReady          bool        `json:"streamReady"`
		StreamStartDate      string      `json:"streamStartDate"`
		AllowStreaming       bool        `json:"allowStreaming"`
		PremiumStreamingOnly bool        `json:"premiumStreamingOnly"`
		NumberOfTracks       int         `json:"numberOfTracks"`
		NumberOfVideos       int         `json:"numberOfVideos"`
		NumberOfVolumes      int         `json:"numberOfVolumes"`
		ReleaseDate          string      `json:"releaseDate"`
		Copyright            string      `json:"copyright"`
		Type                 string      `json:"type"`
		Version              interface{} `json:"version"`
		URL                  string      `json:"url"`
		Cover                string      `json:"cover"`
		VideoCover           interface{} `json:"videoCover"`
		Explicit             bool        `json:"explicit"`
		Upc                  string      `json:"upc"`
		Popularity           int         `json:"popularity"`
		AudioQuality         string      `json:"audioQuality"`
		Artists              TidalAlbum  `json:"artists"`
	} `json:"items"`
}

// TidalAlbumSearch : JSON structure from TIDAL API when searching for albums
type TidalAlbumSearch struct {
	Albums struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
	} `json:"albums"`
}

// TidalAlbum : sss
type TidalAlbum struct {
	ID                   int          `json:"id"`
	Title                string       `json:"title"`
	Duration             int          `json:"duration"`
	StreamReady          bool         `json:"streamReady"`
	StreamStartDate      string       `json:"streamStartDate"`
	AllowStreaming       bool         `json:"allowStreaming"`
	PremiumStreamingOnly bool         `json:"premiumStreamingOnly"`
	NumberOfTracks       int          `json:"numberOfTracks"`
	NumberOfVideos       int          `json:"numberOfVideos"`
	NumberOfVolumes      int          `json:"numberOfVolumes"`
	ReleaseDate          string       `json:"releaseDate"`
	Copyright            string       `json:"copyright"`
	Type                 string       `json:"type"`
	Version              interface{}  `json:"version"`
	URL                  string       `json:"url"`
	Cover                string       `json:"cover"`
	VideoCover           interface{}  `json:"videoCover"`
	Explicit             bool         `json:"explicit"`
	Upc                  string       `json:"upc"`
	Popularity           int          `json:"popularity"`
	AudioQuality         string       `json:"audioQuality"`
	Artists              TidalArtists `json:"artists"`
}

// TidalTrackSearch : JSON structure from TIDAL API when searching for tracks
type TidalTrackSearch struct {
	Tracks struct {
		Items []struct {
			ID int `json:"id"`
		} `json:"items"`
	} `json:"tracks"`
}

// TidalArtists : TidalArtists
type TidalArtists []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// TidalTrack : The structure of the JSON in the response from TIDAL API when requesting a track
type TidalTrack struct {
	ID                   int          `json:"id"`
	Title                string       `json:"title"`
	Duration             int          `json:"duration"`
	Replaygain           float64      `json:"replayGain"`
	Peak                 float64      `json:"peak"`
	Allowstreaming       bool         `json:"allowStreaming"`
	Streamready          bool         `json:"streamReady"`
	Streamstartdate      string       `json:"streamStartDate"`
	Premiumstreamingonly bool         `json:"premiumStreamingOnly"`
	Tracknumber          int          `json:"trackNumber"`
	Volumenumber         int          `json:"volumeNumber"`
	Version              interface{}  `json:"version"`
	Popularity           int          `json:"popularity"`
	Copyright            string       `json:"copyright"`
	URL                  string       `json:"url"`
	Explicit             bool         `json:"explicit"`
	Artists              TidalArtists `json:"artists"`
	Album                struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Cover string `json:"cover"`
	} `json:"album"`
}

// TidalIsrcTracks : The structure of the JSON in the response from TIDAL API when requesting an ISRC track
type TidalIsrcTracks struct {
	Limit              int `json:"limit"`
	Offset             int `json:"offset"`
	TotalNumberOfItems int `json:"totalNumberOfItems"`
	Items              []struct {
		ID                   int          `json:"id"`
		Title                string       `json:"title"`
		Duration             int          `json:"duration"`
		ReplayGain           float64      `json:"replayGain"`
		Peak                 float64      `json:"peak"`
		AllowStreaming       bool         `json:"allowStreaming"`
		StreamReady          bool         `json:"streamReady"`
		StreamStartDate      string       `json:"streamStartDate"`
		PremiumStreamingOnly bool         `json:"premiumStreamingOnly"`
		TrackNumber          int          `json:"trackNumber"`
		VolumeNumber         int          `json:"volumeNumber"`
		Version              interface{}  `json:"version"`
		Popularity           int          `json:"popularity"`
		Copyright            string       `json:"copyright"`
		URL                  string       `json:"url"`
		Isrc                 string       `json:"isrc"`
		Explicit             bool         `json:"explicit"`
		Artists              TidalArtists `json:"artists"`
		Album                struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Cover string `json:"cover"`
		} `json:"album"`
	} `json:"items"`
}
