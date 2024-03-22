package voice

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"marvis/utils/command"
	"marvis/utils/common"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Voice interface {
	WaitForCommand() (*Message, error)
	WaitForCommandShell() (*Message, error)
	Talk() error
}

type voice struct {
	config *viper.Viper
}

type Message struct {
	Command string
	Request string
	Ok      bool
}

func NewVoice(ctx context.Context, config *viper.Viper) Voice {
	return &voice{
		config: config,
	}
}

// Windows only - For now
func (v *voice) WaitForCommand() (*Message, error) {
	cfg := v.config
	var rMsg = &Message{}

	// Ask for a command to execute
	// Convert voice command into text

	// Windows, need to install
	// py -m pip install SpeechRecognition
	// py -m pip install pyaudio

	speechfile := cfg.GetString("voice.speechfile")
	cmd := exec.Command(
		"py", // change to python3 for linux. TODO: Need to make this dynamic
		"speechtotextV3.py",
		fmt.Sprintf("--filename=%s", speechfile),
	)
	cmd.Dir = "assets"
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to run the script:", err)
		return nil, err
	}

	// Read the file
	fullSpeechfilePath := fmt.Sprintf("assets/%s", speechfile)
	message, err := ioutil.ReadFile(fullSpeechfilePath)
	if err != nil {
		log.Printf("Failed to read the file %s: %v", fullSpeechfilePath, err)
		return nil, err
	}
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

	return rMsg, nil
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
	default: // case common.CheckWordsExist(m, cfg.GetString("voice.triggerword")):
		// Default command for assistant
		msg.Command = command.Def()
		msg.Ok = true
		return
	}
}

func (v *voice) Talk() error {
	// Check if using windows or linux
	switch runtime.GOOS {
	case "windows":
		v.WindowsTalk()
		log.Println("Running on Windows")
	case "linux", "darwin", "freebsd", "netbsd", "openbsd":
		log.Println("Running on a Unix-like OS")
	default:
		log.Printf("Running on an unsupported OS: %s\n", runtime.GOOS)
		return errors.New("Running on an unsupported OS: %s\\n\", runtime.GOOS")
	}
	return nil
}

func (v *voice) WindowsTalk() error {
	cmdStr := `py tts.py text-to-speech.txt say.mp3 && ffplay -v 0 -nodisp -autoexit say.mp3`
	cmd := exec.Command("cmd", "/C", cmdStr)
	cmd.Dir = "assets"
	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println("Command executed successfully")
	return nil
}

func (v *voice) LinuxTalk() {

}
