package internal

import (
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewRenderSystem() *ecs.System {
	return &ecs.System{
		RequiredComponents: []ecs.ComponentID{TransformComponentID, SpriteComponentID},
		UpdateFunc:         Render,
	}
}

func Render(entities []uint64, components map[ecs.ComponentID]map[uint64]any) {
	for _, entity := range entities {
		transform := components[TransformComponentID][entity].(*Transform)
		sprite := components[SpriteComponentID][entity].(*Sprite)

		// Создаем прямоугольник для отрисовки с учетом размера из Transform
		destRect := rl.Rectangle{
			X:      transform.X,
			Y:      transform.Y,
			Width:  transform.Width,  // Используем размер из Transform
			Height: transform.Height, // Используем размер из Transform
		}

		// Рисуем спрайт с масштабированием
		rl.DrawTexturePro(
			sprite.Texture,
			sprite.Frame,
			destRect,
			rl.Vector2{X: 0, Y: 0}, // Начало координат в левом верхнем углу
			0,                      // Угол поворота
			rl.White,
		)
	}
}
