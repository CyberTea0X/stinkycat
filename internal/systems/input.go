package systems

import (
	"slices"

	"github.com/CyberTea0X/stinkycat/internal/animation"
	c "github.com/CyberTea0X/stinkycat/internal/components"
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewInputSystem(alib *animation.AnimationLibrary) *ecs.System {
	return &ecs.System{
		RequiredComponents: []ecs.ComponentID{
			c.AnimationComponentID,
			c.InputBindingsComponentID,
		},
		UpdateFunc: NewInputFunc(alib),
	}
}

func NewInputFunc(alib *animation.AnimationLibrary) ecs.UpdateFunc {
	return func(entities []uint64, components map[ecs.ComponentID]map[uint64]any) {
		for _, entity := range entities {
			inputBindings := components[c.InputBindingsComponentID][entity].(*c.InputBindings)
			anim := components[c.AnimationComponentID][entity].(*c.Animation)

			// Проверка прыжка
			if rl.IsKeyPressed(inputBindings.JumpKey) && anim.CurrentAnimation != "jump" {
				if frames, exists := alib.GetAnimation("jump"); exists {
					anim.Frames = frames
					anim.CurrentFrame = 0
					anim.Timer = 0
					anim.CurrentAnimation = "jump"
				}
			}

			if rl.IsKeyReleased(inputBindings.JumpKey) && anim.CurrentAnimation == "jump" {
				if frames, exists := alib.GetAnimation("idle"); exists {
					anim.Frames = frames
					anim.CurrentFrame = 0
					anim.Timer = 0
					anim.CurrentAnimation = "idle"
				}
			}

			// Проверка ходьбы
			isWalking := slices.ContainsFunc(inputBindings.WalkKeys, rl.IsKeyDown)

			if isWalking && anim.CurrentAnimation != "walk" {
				if frames, exists := alib.GetAnimation("walk"); exists {
					anim.Frames = frames
					anim.CurrentFrame = 0
					anim.Timer = 0
					anim.CurrentAnimation = "walk"
				}
			} else if !isWalking && anim.CurrentAnimation == "walk" {
				if frames, exists := alib.GetAnimation("idle"); exists {
					anim.Frames = frames
					anim.CurrentFrame = 0
					anim.Timer = 0
					anim.CurrentAnimation = "idle"
				}
			}
		}
	}
}
