package apps

import (
    "log"
)

type CLI struct{}

func (c *CLI) Run() {
	log.Print("Starting CLI")
}
