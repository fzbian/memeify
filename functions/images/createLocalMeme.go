package images

//TODO: change all functions to scope of services

/*func CustomMeme(top string, bottom string, url string) string {
	nameFile, im := files.GetImageURL(url)
	bounds := im.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	max := math.Round(float64(width) / 10)

	dc := gg.NewContextForImage(im)

	if err := dc.LoadFontFace("fonts/impact.ttf", max); err != nil {
		panic(err)
	}

	positionX := float64(width / 2)
	positionTopY := float64(height / 6)
	positionBottomY := float64(5 * height / 6)

	dc.SetRGB(0, 0, 0)
	n := 6
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				continue
			}
			x := positionX + float64(dx)
			ytop := positionTopY + float64(dy)
			ybottom := positionBottomY + float64(dy)
			dc.DrawStringAnchored(strings.ToUpper(top), x, ytop, 0.5, 0)
			dc.DrawStringAnchored(strings.ToUpper(bottom), x, ybottom, 0.5, 1)
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(strings.ToUpper(top), positionX, positionTopY, 0.5, 0)
	dc.DrawStringAnchored(strings.ToUpper(bottom), positionX, positionBottomY, 0.5, 1)

	buff := new(bytes.Buffer)
	jpeg.Encode(buff, dc.Image(), nil)
	os.WriteFile(nameFile, buff.Bytes(), 0644)
	return nameFile
}*/
