package systems

import (
	c "github.com/CyberTea0X/stinkycat/internal/components"
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewAnimationSystem() *ecs.System {
	return &ecs.System{
		RequiredComponents: []ecs.ComponentID{c.AnimationComponentID},
		UpdateFunc:         UpdateAnimations,
	}
}

func UpdateAnimations(entities []uint64, components map[ecs.ComponentID]map[uint64]any) {
	animations := components[c.AnimationComponentID]
	delta := rl.GetFrameTime()

	for _, entity := range entities {
		anim := animations[entity].(*c.Animation)
		anim.Timer += delta

		if anim.Timer >= anim.FrameTime {
			anim.CurrentFrame++
			anim.Timer = 0

			if anim.CurrentFrame >= len(anim.Frames) {
				anim.CurrentFrame = 0 // Повтор анимации
			}
		}
	}
}
