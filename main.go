//version for developer
//learning on youtube
package main

import (
	"fmt"

	"github.com/silabig1294/discord-bigbot/bot"
	"github.com/silabig1294/discord-bigbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
