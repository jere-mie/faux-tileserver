package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"net/http"

	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/IBMPlexSans-Bold.ttf
var fontData []byte

//go:embed static/*
var embeddedAssets embed.FS

func main() {
	// Define the port flag with a default value
	port := flag.String("port", "3000", "Port to run the server on")
	flag.Parse()

	app := fiber.New()

	// Disable CORS restrictions
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})

	// Serve the embedded assets directory
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embeddedAssets),
		PathPrefix: "static",
		// Browse:     true,
	}))

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

	// Start the server on the specified port
	log.Fatal(app.Listen(":" + *port))
}
