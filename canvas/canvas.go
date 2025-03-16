package canvas

import (
	"bufio"
	"fmt"
	"os"
	"ray-tracer-challenge/geometry"
	"strings"
)

type Canvas struct {
	Pixels [][]geometry.Tuple
	Height int
	Width  int
}

func NewCanvas(width, height int, bg ...geometry.Tuple) *Canvas {
	var bgColor geometry.Tuple
	if len(bg) > 0 {
		bgColor = bg[0]
	} else {
		bgColor = geometry.NewColor(0, 0, 0)
	}

	pixels := make([][]geometry.Tuple, height)
	for i := range pixels {
		pixels[i] = make([]geometry.Tuple, width)
		for j := range pixels[i] {
			pixels[i][j] = bgColor
		}
	}
	return &Canvas{
		Pixels: pixels,
		Height: height,
		Width:  width,
	}
}

func (c *Canvas) WritePixel(row, col int, color geometry.Tuple) {
	c.Pixels[row][col] = color
}

func (c *Canvas) PixelAt(row, col int) geometry.Tuple {
	return c.Pixels[row][col]
}

func (c *Canvas) ToPPM() string {
	fmt.Println("Converting to PPM")
	var sb strings.Builder
	_, err := fmt.Fprintf(&sb, "P3\n%d %d\n255\n", c.Width, c.Height)
	if err != nil {
		return ""
	}

	for i := 0; i < c.Height; i++ {
		var currentString strings.Builder
		for j := 0; j < c.Width; j++ {
			color := c.PixelAt(i, j).ToColorString()

			if currentString.Len()+len(color)+1 > 70 {
				sb.WriteString(currentString.String())
				sb.WriteString("\n")
				currentString.Reset()
				currentString.WriteString(color + " ")
			} else {
				currentString.WriteString(color + " ")
			}
		}
		sb.WriteString(currentString.String() + "\n")
	}

	return sb.String()
}

func (c *Canvas) SaveImage(path string) error {
	fmt.Println("Saving image...")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(c.ToPPM())
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
