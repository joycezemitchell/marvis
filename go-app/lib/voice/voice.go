package voice

import (
    "context"
)

type Voice interface{
    WaitForCommand() (*Message, error) 
}


type voice struct{}

type Message struct {
    Command string
    Ok bool   
}


func NewVoice(ctx context.Context) Voice {
    return &voice{}
}


func (v *voice) WaitForCommand() (*Message, error) {
    // Ask for a command to execute

    // Convert command into voice
    
    // Save command into a file
    
    // Read the file
    
    // Process file and extract the command

    return &Message{}, nil
}

 
            
