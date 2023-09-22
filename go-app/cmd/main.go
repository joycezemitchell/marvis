package main

import(
    "marvis/app/assistant"
    "marvis/lib/voice"
    "marvis/lib/gpt"
    "marvis/lib/config"
    "context"
    "log"
)


const (
   path = "config"
   fileName = "config"
)

func main(){
    ctx := context.Background()
    
    // Load configuration    
    configObj := config.NewConfig(ctx, fileName, path)
    cfg, err := configObj.Load()
    if err != nil {
        log.Println("Error loading configuration file ", err)
    }

    // Start voice assistant 
    gptObj :=  gpt.NewGPT(ctx, cfg)
    voiceObj := voice.NewVoice(ctx, cfg) 
    a := assistant.NewAssistant(ctx, voiceObj, gptObj)
    a.Start()

    // Start localhost server
    // go server.Start()
}
