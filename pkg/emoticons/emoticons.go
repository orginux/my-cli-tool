package emoticons

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/orginux/my-cli-tool/pkg/config"
)

type App struct {
	Config *config.Config
}

func (a *App) Emote(name, dest string) {
	emoticon := a.Config.Emoticon[name]

	switch dest {
	case "clipboard":
		clipboard.WriteAll(emoticon)
		fmt.Println("Coppied: ", emoticon)
	default:
		fmt.Println(emoticon)
	}
}

func New() (*App, error) {
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	return &App{Config: c}, nil
}
