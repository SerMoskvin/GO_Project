package main

import (
	backend "BG_site/Backend"
)

func main() {
	backend.InitDB()
	defer backend.CloseDB()
	backend.HandleRequest()
}
