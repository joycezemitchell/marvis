package gpt

import (
    "context"
)

type GPT interface{
    CreatePlaylist(msg *Message)
    Ask(msg *Message) 
}


type gpt struct{}

type Message struct {
    Request string
}
    
func NewGPT(ctx context.Context) GPT {
    return &gpt{}
}


func (g *gpt) CreatePlaylist(msg *Message) { 

}

 
func (g *gpt) Ask(msg *Message) { 

}

 
