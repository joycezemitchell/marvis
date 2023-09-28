package assistant

import (
    "context"
    "log"
    "marvis/lib/voice"
    "marvis/utils/command"
    "marvis/lib/gpt"
    "marvis/lib/spotify"
    "encoding/json"
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

    // Run forever until it stop
    for {
        // Ask for a command to execute
        // Convert command into voice
        // Save command into a file
        // Read the file
        // Process file and extract the command
        // vc, err := a.voice.WaitForCommand()
        vc, err := a.voice.WaitForCommandShell()

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
                    log.Println("Create a plalist")
                    a.CreatePlaylist(gptMessage)  
                // If anything else call gpt for general purpose inquiry 
                default:
                    log.Println("default command")
                    a.gpt.Ask(gptMessage)
            }
        } 
    }
}

func (a *assist) CreatePlaylist(gptMessage *gpt.Message) {
    var playlist spotify.Playlist

    // Ask gpt to create the playlist
    plBytes := a.gpt.CreatePlaylist(gptMessage)

    // Convert the result to spotify playlist model
    err := json.Unmarshal(plBytes, &playlist)
    if err != nil {
        log.Println(err)
    }
    log.Printf("%+v\n", playlist)

    // Call spotify to create the playlist
    a.spotify.CreatePlaylist(&playlist)

    // Play the tracks, maybe
}





