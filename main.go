package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/akashvbabu/testbeat/beater"
)

func main() {
	err := beat.Run("testbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
