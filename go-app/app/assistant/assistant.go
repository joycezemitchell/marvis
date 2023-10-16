package assistant

import (
    "context"
    "log"
    "marvis/lib/voice"
    "marvis/utils/command"
    "marvis/lib/gpt"
    "marvis/lib/spotify"
    "encoding/json"
    "flag"
)

type Assistant interface {
    Start()
}

type assist struct{
    voice voice.Voice
    gpt gpt.GPT
    spotify spotify.Spotify
}

func NewAssistant(ctx context.Context, v voice.Voice, g gpt.GPT, s spotify.Spotify) Assistant {
    return &assist{
        voice: v,
        gpt: g,
        spotify: s,
    }
}

func (a *assist) Start() {
    log.Println("Starting assistant app...")
    isVoice := flag.Bool("isVoice", true, "Use either keyboard or voice command")
    var vc *voice.Message 
    var err error

    flag.Parse()
    if *isVoice {
        log.Println("Running using voice command")
    } else {
        log.Println("Running using keyboard")
    }

    // Run forever until it stop
    for {
        log.Println("Ready for new command")
        // Ask for a command to execute
        // Convert command into voice
        // Save command into a file
        // Read the file
        // Process file and extract the command
        if *isVoice {
            // Running using voice command
            vc, err = a.voice.WaitForCommand()
        }else{
            // Running using keyboard
            vc, err = a.voice.WaitForCommandShell()
        } 

        if  err != nil {
            // Log error
            log.Println(err)
        }

        // Exit application
        if vc.Command == command.Exit() {
            log.Println("exit")
            break;
        }     

        // Run the assistant if command has the trigger word
        if vc.Ok {
            gptMessage := &gpt.Message{
                Request : vc.Request,
            }

            // Check what kind of command has given
            switch(vc.Command) {
                // If playlist, call gpt that creates playlist
                // Call spotify and play tracks
                case command.Playlist():
                    log.Println("Create a playlist")
                    a.CreatePlaylist(gptMessage)  
                    break
                // If anything else call gpt for general purpose inquiry 
                default:
                    log.Println("default command")
                    a.gpt.Ask(gptMessage)
                    break
            }
        } 
    }
}

func (a *assist) CreatePlaylist(gptMessage *gpt.Message) {
    var pl spotify.Playlist
    var playTracks spotify.PlayTracks

    // Ask gpt to create the playlist
    plBytes := a.gpt.CreatePlaylist(gptMessage)

    // Convert the result to spotify playlist model
    err := json.Unmarshal(plBytes, &pl)
    if err != nil {
        log.Println(err)
    }

    // Search for songs and tracks ids
    for _, track := range pl.Tracks {
        log.Println(track.Title, track.Artist)
        tr := a.spotify.Search(track.Artist, track.Title) 
        log.Println(tr)
        playTracks.Track = append(playTracks.Track, tr)
    }
    
    // Play all tracks
    a.spotify.PlayTracks(playTracks)

    // Create the actual playlist in the background

    // log.Printf("%+v\n", playlist)
    // Call spotify to create the playlist
    // a.spotify.CreatePlaylist(&playlist)
}





