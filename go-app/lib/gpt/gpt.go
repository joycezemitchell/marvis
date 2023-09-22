package gpt

import (
    "context"
    "github.com/spf13/viper"
)

type GPT interface{
    CreatePlaylist(msg *Message)
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


func (g *gpt) CreatePlaylist(msg *Message) { 

}

 
func (g *gpt) Ask(msg *Message) { 

}

 
