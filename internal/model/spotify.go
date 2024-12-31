package model

type SpotifyCurrentlyPlaying struct {
	Item                 SpotifyTrack    `json:"item"`
	CurrentlyPlayingType string          `json:"currently_playing_type"`
	IsPlaying            bool            `json:"is_playing"`
	ProgressMs           int             `json:"progress_ms"`
	Timestamp            int64           `json:"timestamp"`
	Context              *SpotifyContext `json:"context"`
}

type SpotifyTrack struct {
	Name         string              `json:"name"`
	Album        SpotifyAlbum        `json:"album"`
	Artists      []SpotifyArtist     `json:"artists"`
	DurationMs   int                 `json:"duration_ms"`
	Explicit     bool                `json:"explicit"`
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Id           string              `json:"id"`
	Uri          string              `json:"uri"`
}

type SpotifyAlbum struct {
	Name         string              `json:"name"`
	Artists      []SpotifyArtist     `json:"artists"`
	Images       []SpotifyImage      `json:"images"`
	ReleaseDate  string              `json:"release_date"`
	TotalTracks  int                 `json:"total_tracks"`
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Id           string              `json:"id"`
	Uri          string              `json:"uri"`
}

type SpotifyArtist struct {
	Name         string              `json:"name"`
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Id           string              `json:"id"`
	Uri          string              `json:"uri"`
}

type SpotifyImage struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyExternalUrls struct {
	Spotify string `json:"spotify"`
}

type SpotifyContext struct {
	ExternalUrls SpotifyExternalUrls `json:"external_urls"`
	Uri          string              `json:"uri"`
	Type         string              `json:"type"`
	Id           string              `json:"id"`
}

type SpotifyErrorResponse struct {
	Error SpotifyErrorDetail `json:"error"`
}

type SpotifyErrorDetail struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

type SpotifyCurrentlyPlayingTrackUpdate struct {
	ProgressMs int   `json:"progress_ms"`
	Timestamp  int64 `json:"timestamp"`
}
