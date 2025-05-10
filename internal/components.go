package internal

import (
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TransformComponentID ecs.ComponentID = iota
	SpriteComponentID
)

// Transform — компонент для позиции и размера
type Transform struct {
	X, Y          float32
	Width, Height float32
}

// Sprite — компонент для текстуры и параметров отрисовки
type Sprite struct {
	Texture rl.Texture2D
	Frame   rl.Rectangle // Регион спрайта в текстуре
}
