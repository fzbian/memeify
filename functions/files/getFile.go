package files

import (
	"image"
	"log"
	"os"
)

func GetFile(nameFile string) image.Image {
	path := ""
	for _, v := range GlobalPaths {
		if v.Name == nameFile {
			path = v.Path
			break
		}
	}
	f, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return im
}
