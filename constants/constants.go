package constants

import (
	"os"

	"github.com/shoshta73/homehub/log"
)

const DATA_DIR string = "data"

func init() {
	_, err := os.Stat(DATA_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating data directory")

			err := os.Mkdir(DATA_DIR, 0755)
			if err != nil {
				log.Fatal(err)
			}

			log.Info("Data directory created")
		} else {
			log.Fatal(err)
		}
	}
}
