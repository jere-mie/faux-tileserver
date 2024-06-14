package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
)

//go:embed fonts/IBMPlexSans-Bold.ttf
var fontData []byte

func main() {
	app := fiber.New()

	app.Get("/:z/:x/:y.png", func(c *fiber.Ctx) error {
		xStr := c.Params("x")
		yStr := c.Params("y")
		zStr := c.Params("z")

		x, err := strconv.Atoi(xStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid X value")
		}

		y, err := strconv.Atoi(yStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Y value")
		}

		z, err := strconv.Atoi(zStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Z value")
		}

		const W = 256
		const H = 256

		// Write font data to a temporary file
		tmpFile, err := os.CreateTemp("", "font_*.ttf")
		if err != nil {
			log.Fatalf("could not create temporary file: %v", err)
		}
		defer os.Remove(tmpFile.Name()) // Clean up the temporary file
		if _, err := tmpFile.Write(fontData); err != nil {
			tmpFile.Close()
			log.Fatalf("could not write font data to temporary file: %v", err)
		}
		if err := tmpFile.Close(); err != nil {
			log.Fatalf("could not close temporary file: %v", err)
		}

		dc := gg.NewContext(W, H)
		dc.SetRGB(1, 1, 1)
		dc.Clear()

		// Load custom font from temporary file
		if err := dc.LoadFontFace(tmpFile.Name(), 40); err != nil {
			log.Fatalf("could not load font: %v", err)
		}

		dc.SetRGB(0, 0, 0)
		dc.DrawRectangle(0, 0, W, H)
		dc.Stroke()

		// Draw each value on its own line
		textX := fmt.Sprintf("X: %d", x)
		textY := fmt.Sprintf("Y: %d", y)
		textZ := fmt.Sprintf("Z: %d", z)
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(textX, W/2, H/3-20, 0.5, 0.5)
		dc.DrawStringAnchored(textY, W/2, H/2, 0.5, 0.5)
		dc.DrawStringAnchored(textZ, W/2, 2*H/3+20, 0.5, 0.5)

		c.Response().Header.Set("Content-Type", "image/png")
		return dc.EncodePNG(c.Response().BodyWriter())
	})

	log.Fatal(app.Listen(":3000"))
}
