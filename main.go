package main

import (
	"gitlab.com/davecremins/safe-deposit-box/apps"
)

func main() {
	app := apps.CLI{}
	app.Run()
}
