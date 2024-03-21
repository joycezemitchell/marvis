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
    "bufio"
    "os"
    "strings"
)

type Voice interface{
    WaitForCommand() (*Message, error) 
    WaitForCommandShell() (*Message, error) 
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
    cfg := v.config
    var rMsg = &Message{}

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
    rMsg.Request = msg

    // Extract command type from request
    v.GetCommand(rMsg)

    return rMsg, nil
}

// This function is for test only. Accept request from shell
func (v *voice) WaitForCommandShell() (*Message, error) {
    var rMsg = &Message{}
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Command-Line Chat")
	fmt.Println("---------------------")

    text, _ := reader.ReadString('\n') 
    text = strings.Replace(text, "\n", "", -1)

    log.Println(text)
    rMsg.Request = text

    // Extract command type from request
    v.GetCommand(rMsg)

    log.Println(rMsg)

    return rMsg,nil
}

func (v *voice) GetCommand(msg *Message) {
    cfg := v.config 
    m := msg.Request

    switch {
    case common.CheckWordsExist(m, cfg.GetString("voice.exitword")):
        // Exit the application
        msg.Command = command.Exit()
        msg.Ok = false
        return
    case common.CheckWordsExist(m, cfg.GetString("voice.playlistword")):
        // Create a playlist
        msg.Command = command.Playlist()
        msg.Ok = true
        return
    case common.CheckWordsExist(m, cfg.GetString("voice.triggerword")):
        // Default command for assistant
        msg.Command =  command.Def()
        msg.Ok = true
        return
    }
} 
            
