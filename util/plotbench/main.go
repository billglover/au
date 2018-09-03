package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func main() {
	title := flag.String("title", "Time Complexity", "title of the resulting chart")
	width := flag.Int("width", 640, "width of the resulting plot")
	height := flag.Int("height", 480, "height of the resulting plot")
	flag.Parse()

	X, Y, _ := scanLines(os.Stdin, os.Stdout)
	plotBenchmark(X, Y, *title, *width, *height, "out.png")
}

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}

func plotBenchmark(X, Y []float64, title string, width, height int, file string) {
	XYs := xy{x: X, y: Y}

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)
	plotter.DefaultLineStyle.Color = color.RGBA{B: 128, A: 255}
	plotter.DefaultGlyphStyle.Color = color.RGBA{R: 255, A: 255}
	plotter.DefaultGridLineStyle = draw.LineStyle{
		Color: color.Gray{128},
		Width: vg.Points(0.25),
	}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	p.Title.Text = title
	font, _ := vg.MakeFont("Courier", 10)
	p.Title.Font = font
	p.Title.Padding = 10

	font, _ = vg.MakeFont("Courier", 8)
	p.X.Label.Font = font
	p.X.Tick.Label.Font = font
	p.X.Label.Text = "len(A)"
	p.Y.Label.Font = font
	p.Y.Tick.Label.Font = font
	p.Y.Label.Text = "ns/op"
	p.X.Label.YAlign = draw.YAlignment(-1.5)
	p.Y.Label.YAlign = draw.YAlignment(1.5)

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(XYs)
	if err != nil {
		log.Panic(err)
	}

	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	img := vgimg.New(vg.Length(width), vg.Length(height))
	dc := draw.New(img)

	dc = draw.Crop(dc, 20, -20, 20, -10)
	p.Draw(dc)

	w, err := os.Create(file)
	if err != nil {
		log.Panic(err)
	}
	png := vgimg.PngCanvas{Canvas: img}
	if _, err = png.WriteTo(w); err != nil {
		log.Panic(err)
	}

	if err = w.Close(); err != nil {
		log.Panic(err)
	}
}

func scanLines(in io.Reader, out io.Writer) ([]float64, []float64, error) {
	X := []float64{}
	Y := []float64{}

	r := io.TeeReader(in, out)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		x, y, ok := parseXY(scanner.Text())
		if ok == false {
			continue
		}

		X = append(X, x)
		Y = append(Y, y)
	}
	fmt.Fprintln(out)

	return X, Y, nil
}

func parseXY(s string) (x, y float64, ok bool) {
	var re = regexp.MustCompile(`(?U)^\w+_(\d+)-.*([\d.]+)\sns/op.*`)
	m := re.FindStringSubmatch(s)
	if m == nil {
		return 0, 0, false
	}

	x, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return 0, 0, false
	}
	y, err = strconv.ParseFloat(m[2], 64)
	if err != nil {
		return 0, 0, false
	}

	return x, y, true
}
