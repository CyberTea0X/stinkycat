package internal

import "stinkycat/pkg/ecs"

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
