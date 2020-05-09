package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "zihun.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

const title = "征婚"

var text = []string{
	"\"年纪不小了，赶紧随便找个人过残生吧。\"",
	"--这是我全家的意思--",
	"生命本来就没有什么意义",
	"现在想要混吃等死也由不得你",
	"总有些人打着为你好的旗帜逼迫你去做些事情",
	"有些事情我不得不做",
	"因为满足他们的需求也是这少数有点意义的事情了",
	"========= 在线招亲，招到为止===========",
	"只有2个要求，活的，女的",
	"",
}

func drawFile() {
	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Draw the background and the guidelines.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	const imgW, imgH = 640, 480
	rgba := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	h := font.HintingNone
	switch *hinting {
	case "full":
		h = font.HintingFull
	}
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    *size,
			DPI:     *dpi,
			Hinting: h,
		}),
	}
	y := 10 + int(math.Ceil(*size**dpi/72))
	dy := int(math.Ceil(*size * *spacing * *dpi / 72))
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(title)) / 2,
		Y: fixed.I(y),
	}
	d.DrawString(title)
	y += dy
	for _, s := range text {
		d.Dot = fixed.P(12, y)
		d.DrawString(s)
		y += dy
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}

func covertToBase64() {
	b, err := ioutil.ReadFile("out.png")
	if err != nil {
		log.Fatal(err)
	}

	s := base64.StdEncoding.EncodeToString(b)
	log.Println(s)
}

func main() {
	drawFile()
	covertToBase64()
}
