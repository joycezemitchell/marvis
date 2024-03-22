package gpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"log"
	"marvis/utils/common"
)

type GPT interface {
	CreatePlaylist(msg *Message) []byte
	Ask(msg *Message)
}

type gpt struct {
	config   *viper.Viper
	messages []openai.ChatCompletionMessage // Add a slice to store conversation messages
}

type Message struct {
	Request string
}

func NewGPT(ctx context.Context, config *viper.Viper) GPT {
	return &gpt{
		config:   config,
		messages: make([]openai.ChatCompletionMessage, 0), // Initialize the slice
	}
}

func (g *gpt) CreatePlaylist(msg *Message) []byte {
	log.Println("GPT: creating the playlist")
	resp := g.GPTRequest(msg.Request, "Create a playlist in the following json format {playlist:[{'title':'xxx', artist:'xxx'}]. Just return the json and not anything else")
	reply := resp.Choices[0].Message.Content
	log.Println("GPT: Playlist has now been generated. Waiting for Spotify...")
	return []byte(reply)
}

func (g *gpt) Ask(msg *Message) {
	resp := g.GPTRequest(msg.Request, "")
	reply := resp.Choices[0].Message.Content

	cfg := g.config
	speechfile := cfg.GetString("ask.speechfile")
	log.Println(speechfile)

	// Save gpt reply to a text file
	common.SaveToTextFile(reply, speechfile)
	log.Println(reply)
}

func (g *gpt) GPTRequest(msg, system string) *openai.ChatCompletionResponse {
	g.config.SetEnvPrefix("gpt")
	key := g.config.GetString("key")

	client := openai.NewClient(key)

	g.messages = append(g.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	if system != "" {
		g.messages = append(g.messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: system,
		})
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4,
			Messages: g.messages,
		},
	)
	if err != nil {
		log.Println(err)
		return nil
	}

	g.messages = append(g.messages, resp.Choices[0].Message)

	return &resp
}
