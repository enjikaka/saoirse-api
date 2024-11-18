package main

// DeezerSearchAlbum : JSON structure from Deezer API when searching for album
type DeezerSearchAlbum struct {
	Data []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Artist struct {
			Name string `json:"name"`
		} `json:"artist"`
		Type string `json:"type"`
	} `json:"data"`
	Total int `json:"total"`
}

// DeezerSearchTrack : JSON structure from Deezer API when searching for track
type DeezerSearchTrack struct {
	Data []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Artist struct {
			Name string `json:"name"`
		} `json:"artist"`
		Album struct {
			Title string `json:"title"`
		} `json:"album"`
		Type string `json:"type"`
	} `json:"data"`
	Total int `json:"total"`
}

// DeezerTrack : JSON structure from Deezer API when getting a track
type DeezerTrack struct {
	ID                 int      `json:"id"`
	Readable           bool     `json:"readable"`
	Title              string   `json:"title"`
	TitleShort         string   `json:"title_short"`
	TitleVersion       string   `json:"title_version"`
	Isrc               string   `json:"isrc"`
	Link               string   `json:"link"`
	Share              string   `json:"share"`
	Duration           int      `json:"duration"`
	TrackPosition      int      `json:"track_position"`
	DiskNumber         int      `json:"disk_number"`
	Rank               int      `json:"rank"`
	ReleaseDate        string   `json:"release_date"`
	ExplicitLyrics     bool     `json:"explicit_lyrics"`
	Preview            string   `json:"preview"`
	Bpm                float64  `json:"bpm"`
	Gain               int      `json:"gain"`
	AvailableCountries []string `json:"available_countries"`
	Contributors       []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Link          string `json:"link"`
		Share         string `json:"share"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`
		Radio         bool   `json:"radio"`
		Tracklist     string `json:"tracklist"`
		Type          string `json:"type"`
		Role          string `json:"role"`
	} `json:"contributors"`
	Artist struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Link          string `json:"link"`
		Share         string `json:"share"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`
		Radio         bool   `json:"radio"`
		Tracklist     string `json:"tracklist"`
		Type          string `json:"type"`
	} `json:"artist"`
	Album struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		Cover       string `json:"cover"`
		CoverSmall  string `json:"cover_small"`
		CoverMedium string `json:"cover_medium"`
		CoverBig    string `json:"cover_big"`
		CoverXl     string `json:"cover_xl"`
		ReleaseDate string `json:"release_date"`
		Tracklist   string `json:"tracklist"`
		Type        string `json:"type"`
	} `json:"album"`
	Type string `json:"type"`
}

// DeezerAlbum : Part of the response when requesting an album in the Deezer API
type DeezerAlbum struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Upc    string `json:"upc"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
}

type DeezerUPCAlbum struct {
	ID int `json:"id"`
}
