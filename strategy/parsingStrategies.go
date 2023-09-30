package main

import (
	"encoding/json"
)

type ParsingStrategy interface {
	Parse([]byte) (Track, error)
}

func formatDuration(durationMs int) int {
	return durationMs / 1000
}

type Track struct {
	AlbumTitle    string `json:"album_title"`
	TrackName     string `json:"track_name"`
	Artist        string `json:"artist"`
	TrackDuration int    `json:"track_duration"`
}

func (t Track) String() string {
	trackJson, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return ""
	}
	return string(trackJson)
}

type Spotify struct{}

// https://developer.spotify.com/documentation/web-api/reference/get-track
type spotifyData struct {
	Album struct {
		Name string `json:"name"`
	} `json:"album"`
	Name       string `json:"name"`
	DurationMs int    `json:"duration_ms"`
	Artists    []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

func (s *Spotify) Parse(data []byte) (Track, error) {
	var spotify spotifyData
	if err := json.Unmarshal(data, &spotify); err != nil {
		return Track{}, err
	}

	return Track{
		TrackName:     spotify.Name,
		TrackDuration: formatDuration(spotify.DurationMs),
		AlbumTitle:    spotify.Album.Name,
		Artist:        spotify.Artists[0].Name,
	}, nil
}

type AppleMusic struct{}

// https://developer.apple.com/documentation/applemusicapi/songs/attributes
type appleMusicData struct {
	Track      string `json:"name"`
	DurationMs int    `json:"durationInMillis"`
	Album      string `json:"albumName"`
	Artist     string `json:"artistName"`
}

func (a *AppleMusic) Parse(data []byte) (Track, error) {
	var appleMusic appleMusicData
	if err := json.Unmarshal(data, &appleMusic); err != nil {
		return Track{}, err
	}

	return Track{
		TrackName:     appleMusic.Track,
		TrackDuration: formatDuration(appleMusic.DurationMs),
		AlbumTitle:    appleMusic.Album,
		Artist:        appleMusic.Artist,
	}, nil
}
