package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	s "github.com/connordennison/soundcloud-rpc/structs"
	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
)

var (
	connectedToDiscordRPC = false
)

func getCurrentTrack(creds s.SoundCloudApiCreds) s.PlayHistoryTrack {
	// first we make the request to the soundcloud api (https://api-v2.soundcloud.com/me/play-history/tracks?limit=1) with the client id and auth token as query params
	resp, err := http.Get("https://api-v2.soundcloud.com/me/play-history/tracks?limit=1&client_id=" + creds.ClientId + "&auth_token=" + creds.AuthToken)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// then we decode the response into a PlayHistory struct
	var playHistory s.PlayHistory
	err = json.NewDecoder(resp.Body).Decode(&playHistory)
	if err != nil {
		panic(err)
	}

	// then we return the first track in the PlayHistory struct
	return playHistory.Collection[0].PlayHistoryTrack
}

func getUserInfo(creds s.SoundCloudApiCreds) s.User {
	resp, err := http.Get("https://api-v2.soundcloud.com/me?client_id=" + creds.ClientId + "&auth_token=" + creds.AuthToken)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user s.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user
}

func getCredentials() s.SoundCloudApiCreds {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var creds s.SoundCloudApiCreds
	err = json.NewDecoder(file).Decode(&creds)
	if err != nil {
		panic(err)
	}

	return creds
}

func connectRPC() {
	err := client.Login("1121969757122478090")
	if err != nil {
		panic(err)
	}
	connectedToDiscordRPC = true
	log.Println("Connected to Discord RPC")
}

func updatePresence(playHistTrack s.PlayHistoryTrack, user s.User) {
	startTime := time.Unix(0, playHistTrack.PlayedAt*int64(time.Millisecond))
	duration := time.Duration(playHistTrack.Track.FullDuration) * time.Millisecond
	endTime := startTime.Add(duration)
	currentTime := time.Now()

	if currentTime.After(endTime) {
		client.Logout()
		connectedToDiscordRPC = false
		return
	}

	if !connectedToDiscordRPC {
		connectRPC()
	}

	err := client.SetActivity(client.Activity{
		State:      playHistTrack.Track.User.Username,
		Details:    playHistTrack.Track.Title,
		LargeImage: playHistTrack.Track.ArtworkUrl,
		LargeText:  playHistTrack.Track.Title,
		SmallImage: user.AvatarUrl,
		SmallText:  fmt.Sprintf("Logged in as %s", user.Username),
		Buttons: []*client.Button{
			{
				Label: "Listen on SoundCloud",
				Url:   playHistTrack.Track.PermalinkUrl,
			},
		},
		Timestamps: &client.Timestamps{
			Start: &startTime,
			End:   &endTime,
		},
	})
	if err != nil {
		// in this scenario we don't want to panic as it'd be useless to halt the whole program over one failed rpc update
		// instead we just print the error to the console
		log.Println(err)
	}
}

func getIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	if runtime.GOOS == "windows" {
		systray.SetIcon(getIcon("assets/icon.ico"))
	} else {
		systray.SetIcon(getIcon("assets/icon.png"))
	}
	systray.SetTitle("SoundCloud RPC")
	systray.SetTooltip("SoundCloud RPC")
	systray.AddMenuItem("SoundCloud RPC by cnnd", "SoundCloud RPC by cnnd").Disable()
	systray.AddSeparator()
	bQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-bQuit.ClickedCh
		systray.Quit()
	}()

	creds := getCredentials()
	user := getUserInfo(creds) // we only need to run this once as it's rare for the userinfo to change within one session
	connectRPC()

	for {
		playHistTrack := getCurrentTrack(creds)
		updatePresence(playHistTrack, user)
		time.Sleep(time.Second * 8)
	}
}

func onExit() {
	client.Logout()
}
