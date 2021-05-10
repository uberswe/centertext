package centertext

import (
	"errors"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
)

func OnImage(img image.Image, f truetype.Font, text string) (image.Image, error) {
	b := img.Bounds().Max
	size := 10.0
	rgba := image.NewRGBA(image.Rect(0, 0, b.X, b.Y))
	draw.Draw(rgba, img.Bounds(), img, image.ZP, draw.Src)

	opts := truetype.Options{}
	opts.Size = size
	face := truetype.NewFace(&f, &opts)

	tw := 0
	th := 0
	for _, glyph := range text {
		gb, ga, ok := face.GlyphBounds(glyph)
		if ok != true {
			return rgba, errors.New("could not calculate advance width of glyph")
		}
		tw += int(float64(ga) / 64)
		if th < (int(float64(gb.Max.Y)/64) - int(float64(gb.Min.Y)/64)) {
			th = int(float64(gb.Max.Y)/64) - int(float64(gb.Min.Y)/64)
		}
	}
	factor := calculateFactor(float64(b.X), float64(tw), 10)
	size = size * factor
	opts.Size = size
	face = truetype.NewFace(&f, &opts)

	tw = 0
	th = 0
	for _, glyph := range text {
		gb, ga, ok := face.GlyphBounds(glyph)
		if ok != true {
			return img, errors.New("could not calculate advance width of glyph")
		}
		tw += int(float64(ga) / 64)
		if th < (int(float64(gb.Max.Y)/64) - int(float64(gb.Min.Y)/64)) {
			th = int(float64(gb.Max.Y)/64) - int(float64(gb.Min.Y)/64)
		}
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(&f)
	c.SetFontSize(size)
	c.SetClip(img.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.Black)
	c.SetHinting(font.HintingNone)

	pt := calculatePt(b.X, tw, th)
	_, err := c.DrawString(text, pt)
	if err != nil {
		return rgba, err
	}

	return rgba, nil
}

func calculatePt(containerWidth int, textLength int, textHeight int) fixed.Point26_6 {
	l := containerWidth - textLength
	h := containerWidth - textHeight
	return freetype.Pt(l/2, (h/2)+textHeight)
}

func calculateFactor(containerWidth float64, textLength float64, margin float64) float64 {
	textLength += margin
	s := containerWidth - textLength
	if s < 0 {
		return containerWidth / textLength
	}
	return s / textLength
}
