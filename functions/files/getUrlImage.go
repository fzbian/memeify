package files

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	ABC = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomName(lenght int) string {
	b := make([]rune, lenght)

	for i := range b {
		b[i] = ABC[rand.Intn(len(ABC))]
	}
	return string(b)
}

func GetImageURL(url string) (string, image.Image) {
	response, _ := http.Get(url)
	m, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	buff := new(bytes.Buffer)
	jpeg.Encode(buff, m, nil)

	nameFile := RandomName(15) + ".png"
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile, m
}
