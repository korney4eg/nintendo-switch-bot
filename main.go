package main

import (
	"fmt"

	"github.com/korney4eg/nintendo-switch-bot/internal/telegram"
)

func main() {
	fmt.Println("Starting main loop")
	telegram.MainLoop()
}
