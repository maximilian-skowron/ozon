package main

import (
	"os"
	"github.com/Ozon-Project/ozon/cmd/ozon"
)

func main() {
	if err := ozon.Execute(); err != nil {
		os.Exit(1)
	}
}
