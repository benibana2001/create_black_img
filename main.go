package main

import "github.com/benibana2001/create_black_img/BlackImg"

func main() {
	var(
		maxX = 200
		maxY = 200
	)
	c := BlackImg.NewCreator()
	c.SetRectangle(maxX, maxY)
	c.CreateBlackImage()
	c.WritePng("test.png")
	c.WriteJpeg("test.jpg")
}
