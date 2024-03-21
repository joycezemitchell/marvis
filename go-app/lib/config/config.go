package config

import (
    "context"
    "github.com/spf13/viper"
)

type Config interface{
    Load() (*viper.Viper, error)
}

type config struct{
    filename  string
    path string
}

    
func NewConfig(ctx context.Context, filename, path string) Config {
    return &config{
        filename: filename,
        path: path,
    }
}


func (c *config) Load() (*viper.Viper, error){ 
    // Initialize Viper
	v := viper.New()
	v.SetConfigName(c.filename)
	v.AddConfigPath(c.path)
	v.SetConfigType("yaml")
    v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		// log.Fatalf("Error reading config file: %s", err)
	}

    return v,nil
}

 

 
