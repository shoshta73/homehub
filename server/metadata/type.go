package metadata

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
)

var defaultMetadata = Metadata{
	Version:  1,
	HasAdmin: false,
}

type Metadata struct {
	Version  uint8 `json:"version"`
	HasAdmin bool  `json:"hasAdmin"`
}

func (md *Metadata) UpdateHasAdmin(hasAdmin bool) {
	md.HasAdmin = hasAdmin
}

func (md Metadata) Write() {
	b, err := json.Marshal(md)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(metadatafilepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Info("Closing metadata file")
		f.Close()
	}()

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}

}

func GetMetadata() *Metadata {
	b, err := os.ReadFile(metadatafilepath)
	if err != nil {
		log.Fatal(err)
	}

	m := &Metadata{}
	err = json.Unmarshal(b, m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
