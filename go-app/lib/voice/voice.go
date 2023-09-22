package voice

import (
    "context"
    "github.com/spf13/viper"
    "log"
    "os/exec"
   	"io/ioutil"
    "fmt"
    "marvis/utils/command"
    "marvis/utils/common"
)

type Voice interface{
    WaitForCommand() (*Message, error) 
}

type voice struct{
    config *viper.Viper
}

type Message struct {
    Command string
    Request string
    Ok bool   
}

func NewVoice(ctx context.Context, config *viper.Viper) Voice {
    return &voice{
        config: config,
    }
}

func (v *voice) WaitForCommand() (*Message, error) {
    var rCommand string
    var rOk bool
    cfg := v.config 


    // Ask for a command to execute
    // Convert voice command into text
    speechfile := cfg.GetString("voice.speechfile")
    cmd := exec.Command(
        "python3", 
        "assets/speechtotextV3.py", 
        fmt.Sprintf("--filename=%s", speechfile),
    )
    err := cmd.Run()
    if err != nil {
        log.Println("Failed to run the script:", err)
        return nil, err
    }

    
    // Read the file
    message, err := ioutil.ReadFile(speechfile)
    msg := string(message)

    log.Println(msg)

    switch {
    case common.CheckWordsExist(msg, cfg.GetString("voice.exitword")):
        // Exit the application
        rCommand  = command.Exit()
        rOk = false    
        break;
    case common.CheckWordsExist(msg, cfg.GetString("voice.playlistword")):
        // Create a playlist
        rCommand = command.Playlist()
        rOk = true
        break;
    case common.CheckWordsExist(msg, cfg.GetString("voice.triggerword")):
        // Default command for assistant
        rCommand = command.Def()
        rOk = true
        break;
    }

    return &Message{
        Command: rCommand,
        Request: msg,
        Ok: rOk,
    }, nil
}


 
            
