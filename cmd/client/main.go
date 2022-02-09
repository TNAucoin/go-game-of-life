package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/TNAucoin/go-game-of-life/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	world            *world.World
	pixels           []byte
	canvasImage      *ebiten.Image
	gameStateRunning bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeySpace)) {
		g.gameStateRunning = false
	} else {
		g.gameStateRunning = true
	}
	if g.gameStateRunning {
		g.world.Update()
	} else {
		mx, my := ebiten.CursorPosition()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.world.Paint(mx, my)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	screen.DrawImage(g.canvasImage, nil)
	g.world.Draw(g.pixels)
	screen.ReplacePixels(g.pixels)
	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("%d, %d", mx, my)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		world:            world.NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10)),
		canvasImage:      ebiten.NewImage(screenWidth, screenHeight),
		gameStateRunning: false,
	}
	g.canvasImage.Fill(color.Black)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Go Game of Life")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
