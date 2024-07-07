package backup

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	imageWidth   = 150
	imageHeight  = 200

	buttonX = 10
	buttonY = 10
	buttonW = 100
	buttonH = 50

	moveSpeed = 20.0
)

var (
	dir = "./game/pic"
)

type Game struct {
	images      []*ebiten.Image
	expanded    bool
	cardPos     [3][2]float64
	cardTargets [3][2]float64
	cardIndex   int
	moveStarted time.Time
}

func NewGame() (*Game, error) {
	images := make([]*ebiten.Image, 3)
	for i := 1; i <= 3; i++ {
		path := filepath.Join(dir, strconv.Itoa(i)+".jpg")
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return nil, err
		}
		images[i-1] = img
	}
	return &Game{
		images: images,
		cardPos: [3][2]float64{
			{0, ScreenHeight}, {0, ScreenHeight}, {0, ScreenHeight},
		},
		cardTargets: [3][2]float64{
			{ScreenWidth/2 - 100 - imageWidth/2, ScreenHeight/2 - imageHeight/2},
			{ScreenWidth/2 - imageWidth/2, ScreenHeight/2 - imageHeight/2},
			{ScreenWidth/2 + 100 - imageWidth/2, ScreenHeight/2 - imageHeight/2},
		},
		cardIndex:   -1,
		moveStarted: time.Now(),
	}, nil
}

func (g *Game) Update() error {
	// Check for button click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= buttonX && x <= buttonX+buttonW && y >= buttonY && y <= buttonY+buttonH {
			// Reset cards if button clicked
			g.cardPos = [3][2]float64{
				{0, ScreenHeight}, {0, ScreenHeight}, {0, ScreenHeight},
			}
			g.cardIndex = 0
			g.moveStarted = time.Now()
		}
	}

	// Move cards one by one
	if g.cardIndex >= 0 && g.cardIndex < 3 {
		targetX := g.cardTargets[g.cardIndex][0]
		targetY := g.cardTargets[g.cardIndex][1]
		currentX := g.cardPos[g.cardIndex][0]
		currentY := g.cardPos[g.cardIndex][1]

		dx := targetX - currentX
		dy := targetY - currentY
		dist := moveSpeed

		if dx*dx+dy*dy <= dist*dist {
			g.cardPos[g.cardIndex][0] = targetX
			g.cardPos[g.cardIndex][1] = targetY
			g.cardIndex++
			if g.cardIndex < 3 {
				g.moveStarted = time.Now()
			}
		} else {
			angle := math.Atan2(dy, dx)
			g.cardPos[g.cardIndex][0] += dist * math.Cos(angle)
			g.cardPos[g.cardIndex][1] += dist * math.Sin(angle)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Draw button
	buttonImage := ebiten.NewImage(buttonW, buttonH)
	buttonColor := color.RGBA{0, 255, 0, 255}
	if g.expanded {
		buttonColor = color.RGBA{255, 0, 0, 255}
	}
	buttonImage.Fill(buttonColor)
	opButton := &ebiten.DrawImageOptions{}
	opButton.GeoM.Translate(float64(buttonX), float64(buttonY))
	screen.DrawImage(buttonImage, opButton)

	// Get cursor position
	cursorX, cursorY := ebiten.CursorPosition()

	// Determine if any card should be scaled
	scaleIndex := -1
	for i := 0; i < len(g.images); i++ {
		// Check if the current card is hovered
		hover := float64(cursorX) >= g.cardPos[i][0] && float64(cursorX) <= g.cardPos[i][0]+imageWidth &&
			float64(cursorY) >= g.cardPos[i][1] && float64(cursorY) <= g.cardPos[i][1]+imageHeight

		// Check if this card is overlapping with the next one
		if i < len(g.images)-1 && g.cardPos[i][0]+imageWidth > g.cardPos[i+1][0] {
			// If the current card is hovered, mark it for scaling
			if hover {
				scaleIndex = i
				break
			}
		} else if hover {
			// If there is no overlap, mark the hovered card for scaling
			scaleIndex = i
		}
	}

	// Draw cards in reversed order (left covering right)
	for i := len(g.images) - 1; i >= 0; i-- {
		img := g.images[i]
		op := &ebiten.DrawImageOptions{}

		// Apply scaling if this is the card to be scaled
		if i == scaleIndex {
			op.GeoM.Scale(1.2, 1.2)
			op.GeoM.Translate(-imageWidth*0.1, -imageHeight*0.1-50)
		}

		op.GeoM.Translate(g.cardPos[i][0], g.cardPos[i][1])
		screen.DrawImage(img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
