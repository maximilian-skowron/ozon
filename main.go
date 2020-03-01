package main

import (
	"github.com/Ozon-Project/ozon/cmd/ozon"
	"os"
)

func main() {
	if err := ozon.Execute(); err != nil {
		os.Exit(1)
	}
}
