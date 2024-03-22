package spotify

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"marvis/utils/common"
	"net/http"
	"net/url"
	"strings"
)

type Spotify interface {
	CreatePlaylist(*Playlist)
	Search(artist, song string) string
	PlayTracks(playTracks PlayTracks) error
}

type spotify struct {
	config *viper.Viper
}

type Playlist struct {
	Tracks []Track `json:"playlist",omitempty`
}

type Track struct {
	Title  string `json:"title",omitempty`
	Artist string `json:"artist",omitemtpy`
}

type SpotifyResp struct {
	// SpotipyError  `json:"error"`
	Err struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}

type TrackUri struct {
	Tracks struct {
		Items []struct {
			URI string `json"uri"`
		} `json:"items"`
	} `json:"tracks"`
}

type PlayTracks struct {
	Track []string `json:"uris"`
}

var (
	authToken   = ""
	spotifyResp SpotifyResp
)

func NewSpotify(ctx context.Context, config *viper.Viper) Spotify {
	return &spotify{
		config: config,
	}
}

func (s *spotify) CreatePlaylist(pl *Playlist) {
	log.Println("Creating a playlist")
	var playTracks PlayTracks
	// Search for songs and tracks ids
	for _, track := range pl.Tracks {
		log.Println(track.Title, track.Artist)
		tr := s.Search(track.Artist, track.Title)
		log.Println(tr)
		playTracks.Track = append(playTracks.Track, tr)
	}

	// Create the playlist
	// log.Println(playTracks.Track)

	// Play all the tracks
	// s.PlayTracks(playTracks)
	return
}

func (s *spotify) Search(artist, song string) string {
	log.Println("Search for ", artist, song)
	var trackUri TrackUri
	spotifyURL := "https://api.spotify.com/v1/search"
	query := fmt.Sprintf(song, artist)
	query = url.QueryEscape(query)
	query = strings.ReplaceAll(query, "+", "%20")

	// Add the parameter names back into the URL
	reqURL := fmt.Sprintf("%s?q=%s&type=track&limit=1", spotifyURL, query)
	// log.Println(reqURL)

	// Call request url function
	resp, err := s.CallUrl(reqURL, "GET", new(bytes.Buffer))
	if err != nil {
		log.Println(err)
	}

	// log.Println(string(resp))

	err = json.Unmarshal(resp, &trackUri)
	if err != nil {
		log.Println(err)
	}

	if len(trackUri.Tracks.Items) > 0 {
		// ilog.Println(trackUri.Tracks.Items[0].URI)
		return trackUri.Tracks.Items[0].URI
	}

	return ""
}

func (s *spotify) PlayTracks(playTracks PlayTracks) error {
	log.Println("Playing tracks")

	reqURL := "https://api.spotify.com/v1/me/player/play?device_id=928c0746e8539bbc48ce127e1486bb2a81c505c7"

	payload, err := json.Marshal(playTracks)
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := s.CallUrl(reqURL, "PUT", bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
		return err
	}

	// log.Println(string(resp))
	// It means something went wrong
	if string(resp) != "" {
		json.Unmarshal(resp, &spotifyResp)
		x := spotifyResp.Err.Message
		log.Println("Error: ", x)
		return errors.New(fmt.Sprintf("Error:%v", x))
	}

	return nil
}

func (s *spotify) CallUrl(reqURL, method string, payload *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest(method, reqURL, payload)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	// log.Println("Search authtoken")
	// log.Println(authToken)
	// log.Println(reqURL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode == 401 || resp.StatusCode == 400 {
		// Refresh authtoken
		authToken = s.GenerateAuthToken()

		// Do the request again
		req.Header.Set("Authorization", "Bearer "+authToken)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return body, nil
}

func (s *spotify) GenerateAuthToken() string {
	// Setup config values
	s.config.SetEnvPrefix("spotify")
	clientID := s.config.GetString("client_id")
	clientSecret := s.config.GetString("client_secret")
	refreshToken := s.config.GetString("refresh_token")

	authHeader := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	log.Println(clientID)
	log.Println(clientSecret)
	log.Println(refreshToken)
	log.Println(authHeader)

	// create request body
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	// HTTP request
	req, err := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token",
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", fmt.Sprintf("__Host-device_id=%s; sp_tr=false'", refreshToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		var rb map[string]interface{}
		json.Unmarshal(body, &rb)
		accessToken := rb["access_token"].(string)
		// log.Println(accessToken)
		return accessToken
	}

	log.Println(resp.StatusCode)
	log.Println(string(body))
	return ""
}

// We only need this one time
func (s *spotify) GenerateAuthURL() string {
	// Setup config values
	s.config.SetEnvPrefix("spotify")
	clientID := s.config.GetString("client_id")
	redirectURI := s.config.GetString("redirect_uri")

	log.Println(clientID)
	log.Println(redirectURI)

	// Generate random string
	state := common.GenerateRandomString(16)
	scope := "user-read-private user-read-email"

	spotifyURL := "https://accounts.spotify.com/authorize"
	queryParams := url.Values{
		"response_type": {"code"},
		"client_id":     {clientID},
		"scope":         {scope},
		"redirect_uri":  {redirectURI},
		"state":         {state},
	}

	return spotifyURL + "?" + queryParams.Encode()
}
