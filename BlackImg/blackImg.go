package BlackImg

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type Creator struct {
	img *image.NRGBA
	r image.Rectangle
}

func NewCreator() *Creator{
	return &Creator{}
}

func (c *Creator) SetRectangle(maxX int, maxY int) {
	min := image.Point{X: 0, Y: 0}
	max := image.Point{X: maxX, Y: maxY}
	r := image.Rectangle{Min: min, Max: max}
	c.r = r
	fmt.Printf("サイズ: %v の画像を作成します。\n", c.r.Size())
}

func (c *Creator) CreateBlackImage() {
	// NRGBA を作成
	img := image.NewNRGBA(c.r)

	// 描画
	for h := c.r.Min.Y; h < c.r.Max.Y; h++ {
		for v := c.r.Min.X; v < c.r.Max.X; v++ {
			img.Set(v, h, color.RGBA{
				R: uint8(0),
				G: uint8(0),
				B: uint8(0),
				A: uint8(255),
			})
		}
	}
	c.img = img
}

func (c *Creator) WritePng(name string) {
	// io.Writer 作成
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// エンコード
	if err := png.Encode(f, c.img); err != nil {
		log.Fatal(err)
	}
}

func (c *Creator) WriteJpeg(name string) {
	// io.Writer 作成
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if err := jpeg.Encode(f, c.img, &jpeg.Options{100}); err != nil {
		fmt.Println("error:jpeg\n",err)
	}

}
