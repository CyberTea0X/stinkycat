package systems

import (
	c "github.com/CyberTea0X/stinkycat/internal/components"
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewRenderSystem() *ecs.System {
	return &ecs.System{
		RequiredComponents: []ecs.ComponentID{c.TransformComponentID, c.SpriteComponentID},
		UpdateFunc:         Render,
	}
}

func Render(entities []uint64, components map[ecs.ComponentID]map[uint64]any) {
	for _, entity := range entities {
		transform := components[c.TransformComponentID][entity].(*c.Transform)
		sprite := components[c.SpriteComponentID][entity].(*c.Sprite)
		animation := components[c.AnimationComponentID][entity].(*c.Animation)

		currentFrame := animation.Frames[animation.CurrentFrame]

		rl.DrawTexturePro(
			sprite.Texture,
			currentFrame,
			rl.Rectangle{
				X:      transform.X,
				Y:      transform.Y,
				Width:  transform.Width,
				Height: transform.Height,
			},
			rl.Vector2{X: 0, Y: 0},
			0,
			rl.White,
		)
	}
}
