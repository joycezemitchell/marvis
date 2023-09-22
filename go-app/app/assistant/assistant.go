package assistant

import (
    "context"
    "log"
    "marvis/lib/voice"
    "marvis/utils/command"
    "marvis/lib/gpt"
)

type Assistant interface {
    Start()
}

type assist struct{
    voice voice.Voice
    gpt gpt.GPT
}

func NewAssistant(ctx context.Context, v voice.Voice, g gpt.GPT) Assistant {
    return &assist{
        voice: v,
        gpt: g,
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
        vc, err := a.voice.WaitForCommand()
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
                Request : vc.Command,
            }

            // Check what kind of command has given
            switch(vc.Command) {
                // If playlist, call gpt that creates playlist
                // Call spotify and play tracks
                case command.Playlist():
                    log.Println("Create a plalist")
                    a.gpt.CreatePlaylist(gptMessage)
        
                // If anything else call gpt for general purpose inquiry 
                default:
                    log.Println("default command")
                    a.gpt.Ask(gptMessage)
            }
        } 
    }
}







