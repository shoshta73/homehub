package metadata

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shoshta73/homehub/constants"
	"github.com/shoshta73/homehub/log"
)

var metadatafilepath string

func init() {
	_, err := os.Stat(constants.DATA_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating data directory")
			os.Mkdir(constants.DATA_DIR, 0755)
			log.Info("Data directory created")
		} else {
			log.Fatal(err)
		}
	}

	metadatafilepath = filepath.Join(constants.DATA_DIR, "metadata.json")
	_, err = os.Stat(metadatafilepath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating metadata file")
			f, err := os.Create(metadatafilepath)
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				log.Info("Closing metadata file")
				f.Close()
			}()
			log.Info("Metadata file created")

			log.Info("Marshalling default metadata")
			b, err := json.Marshal(defaultMetadata)
			if err != nil {
				log.Fatal(err)
			}
			log.Info("Marshalling default metadata done")

			log.Info("Writing default metadata")
			_, err = f.Write(b)
			if err != nil {
				log.Fatal(err)
			}

			_ = f.Sync()
			log.Info("Metadata file written")
		} else {
			log.Fatal(err)
		}
	}
}
