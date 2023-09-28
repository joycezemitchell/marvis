package gpt

import (
    "context"
    "github.com/spf13/viper"
    "log"
   	openai "github.com/sashabaranov/go-openai"
)

type GPT interface{
    CreatePlaylist(msg *Message) []byte
    Ask(msg *Message) 
}

type gpt struct{
    config *viper.Viper
}

type Message struct {
    Request string
}
    
func NewGPT(ctx context.Context, config *viper.Viper) GPT {
    return &gpt{
        config: config,
    }
}

func (g *gpt) CreatePlaylist(msg *Message) []byte { 
    resp := g.GPTRequest(msg.Request, "Create a playlist in the following json format {playlist:[{'title':'xxx', artist:'xxx'}]. Just return the json and not anything else") 
    reply := resp.Choices[0].Message.Content
    log.Println(reply)
    return []byte(reply)
}
 
func (g *gpt) Ask(msg *Message) { 

}

func (g *gpt) GPTRequest(msg, system string) *openai.ChatCompletionResponse {
    g.config.SetEnvPrefix("gpt")
    key :=  g.config.GetString("key")
    log.Println(key)

    client := openai.NewClient(key)
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT3Dot5Turbo,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleSystem,
                    Content: system,
                },
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: msg,
                },
            },
        },
    )  
    if err != nil {
        log.Println(err)
    }
   
    return &resp
}



