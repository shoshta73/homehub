package user

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"sync"

	"github.com/charmbracelet/log"

	"github.com/shoshta73/homehub/constants"
)

const defaultSize = 1024

var mapMutex sync.Mutex = sync.Mutex{}
var creatingMap map[int64]bool = map[int64]bool{}

func GenerateIdenticon(u User) error {
	mapMutex.Lock()
	creatingMap[u.Id] = true
	mapMutex.Unlock()

	defer func() {
		mapMutex.Lock()
		creatingMap[u.Id] = false
		mapMutex.Unlock()
	}()

	log.Info("Generating identicon for", "user", u.Username)
	const gridSize = 128
	cellSize := defaultSize / gridSize

	img := image.NewRGBA(image.Rect(0, 0, defaultSize, defaultSize))

	fillColor := parseColorFromHash(u.Avatar[:], 0)
	log.Info("Fill color", "fillColor", fillColor)

	for x := 0; x < gridSize; x++ {
		for y := 0; y <= gridSize; y++ {
			idx := (x*gridSize + y) % len(u.Avatar)
			if idx >= len(u.Avatar) {
				idx = idx % len(u.Avatar)
			}
			if u.Avatar[idx]%2 == 0 {
				fillRectangle(img, x*cellSize, y*cellSize, cellSize, cellSize, fillColor)
			} else {
				fillRectangle(img, x*cellSize, y*cellSize, cellSize, cellSize, parseColorFromHash(u.Avatar[:], idx))
			}
		}
	}

	_, err := os.Stat(filepath.Join(constants.DATA_DIR, "identicons"))
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating identicons directory")
			os.Mkdir(filepath.Join(constants.DATA_DIR, "identicons"), 0755)
			log.Info("Identicons directory created")
		} else {
			return err
		}
	}

	log.Info("Creating identicon", "username", u.Username)
	outFile, err := os.Create(filepath.Join(constants.DATA_DIR, "identicons", u.Username+".png"))
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func parseColorFromHash(hashStr string, ofset int) color.Color {
	r := uint8(hashStr[uint8(hashStr[(ofset+0)%len(hashStr)])])
	g := uint8(hashStr[uint8(hashStr[(ofset+1)%len(hashStr)])])
	b := uint8(hashStr[uint8(hashStr[(ofset+2)%len(hashStr)])])
	a := uint8(hashStr[uint8(hashStr[(ofset+3)%len(hashStr)])])
	return color.RGBA{r, g, b, a}
}

func fillRectangle(img *image.RGBA, x, y, width, height int, col color.Color) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			img.Set(i, j, col)
		}
	}
}
