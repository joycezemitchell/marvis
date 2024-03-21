package command

const (
    exit = "exit"
    def = "default"
    playlist = "playlist"
)

// List of all available commands
var (
    commands = map[string]string {
        exit:exit,
        playlist: playlist,
        def:def,
    }
)

func LoadCommands() map[string]string  {
    return commands
}

func Exit() string {
    return commands[exit]
}

func Def() string {
    return commands[def]
}

func Playlist() string {
    return commands[playlist]
}

