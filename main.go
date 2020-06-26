package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/tskdsb/DigitalWar/types"
)

var world = types.NewMap(8 * 16)
var a int

// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {
	// Write your game's logical update.
	{

	}

	if ebiten.IsDrawingSkipped() {
		// When the game is running slowly, the rendering result
		// will not be adopted.
		return nil
	}

	// Write your game's rendering.
	{
		image, err := ebiten.NewImage(1024, 512, ebiten.FilterDefault)
		if err != nil {
			return err
		}
		a++
		if err := image.Fill(color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: uint8(a % 256),
		}); err != nil {
			return err
		}
		if err := screen.DrawImage(image, nil); err != nil {
			return err
		}
	}

	return nil
}
func main() {
	// Init map
	{
		for i := 0; i < 8*8; i++ {
			world.AddBlock(types.NewBlock(i/8, i%8))
		}
	}

	s := ebiten.DeviceScaleFactor()
	// Call ebiten.Run to start your game loop.
	if err := ebiten.Run(update, 1024, 512, 1/s, "数字战争"); err != nil {
		log.Fatal(err)
	}
}
