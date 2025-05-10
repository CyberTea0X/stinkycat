package animation

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// AnimationLibrary хранит анимации по именам
type AnimationLibrary struct {
	animations map[string][]rl.Rectangle
}

// NewAnimationLibrary создает новый регистр анимаций
func NewAnimationLibrary() *AnimationLibrary {
	return &AnimationLibrary{
		animations: make(map[string][]rl.Rectangle),
	}
}

// AddAnimation добавляет анимацию в регистр
func (al *AnimationLibrary) AddAnimation(name string, frames []rl.Rectangle) {
	al.animations[name] = frames
}

// GetAnimation возвращает кадры анимации по имени
func (al *AnimationLibrary) GetAnimation(name string) ([]rl.Rectangle, bool) {
	frames, exists := al.animations[name]
	return frames, exists
}
