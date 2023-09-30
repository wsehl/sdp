package main

import (
	"fmt"
)

func main() {
	// Create service with Spotify strategy initially
	service := NewParsingService(&Spotify{})

	spotifyData := []byte(`{
	    "album": {
	        "name": "Take Me Back to Eden"
	    },
	    "name": "The Summoning",
	    "duration_ms": 395000,
	    "artists": [{
	        "name": "Sleep Token"
	    }]
	}`)

	track, _ := service.ParseTrack(spotifyData)
	fmt.Println(track)

	// Switch strategy to AppleMusic
	service.SetStrategy(&AppleMusic{})

	appleMusicData := []byte(`{
        "name": "Silvera",
        "durationInMillis": 212000,
        "albumName": "Magma",
		"artistName": "Gojira"
  	}`)

	track, _ = service.ParseTrack(appleMusicData)
	fmt.Println(track)
}
