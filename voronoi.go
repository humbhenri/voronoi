package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var (
	// flags
	width        = flag.Int("w", 640, "width")
	height       = flag.Int("h", 480, "height")
	filename     = flag.String("f", "out", "output file name")
	numcentroids = flag.Int("c", 10, "number of centroids")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: voronoi [flags]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func createImage() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, *width, *height))
}

type point struct {
	x, y  int
	color color.Color
}

func makeRndColor() color.Color {
	r, g, b := uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
	return color.RGBA{r, g, b, 255}
}

func (p *point) distance(q *point) uint {
	return uint(math.Sqrt(math.Pow(float64(p.x-q.x), 2.0) + math.Pow(float64(p.y-q.y), 2.0)))
}

func nearestCentroidColor(p point, centroids []*point) color.Color {
	min := ^uint(0)
	var color color.Color
	for _, c := range centroids {
		if d := p.distance(c); d < min {
			min = d
			color = c.color
		}
	}
	return color
}

func createCentroids(num int, width int, height int) []*point {
	centroids := make([]*point, num)
	for i := 0; i < num; i++ {
		centroids[i] = &point{
			x:     rand.Intn(width),
			y:     rand.Intn(height),
			color: makeRndColor(),
		}
	}
	return centroids
}

func colorize(img *image.RGBA, centroids []*point) {
	p := point{}
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			p.x = x
			p.y = y
			img.Set(x, y, nearestCentroidColor(p, centroids))
		}
	}
}

func saveImgToFile(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	png.Encode(file, img)
	log.Println("Image saved to file " + filename)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	centroids := createCentroids(*numcentroids, *width, *height)
	img := createImage()
	colorize(img, centroids)
	saveImgToFile(img, *filename)
	os.Exit(0)
}
