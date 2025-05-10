package components

import (
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TransformComponentID ecs.ComponentID = iota
	SpriteComponentID
	AnimationComponentID
	InputBindingsComponentID
)

type InputBindings struct {
	JumpKey  int32   // Клавиша для прыжка
	WalkKeys []int32 // Клавиши для ходьбы
}

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

type Animation struct {
	Frames           []rl.Rectangle // Все кадры анимации
	CurrentFrame     int            // Текущий кадр
	FrameTime        float32        // Время между кадрами
	Timer            float32        // Таймер для смены кадров
	CurrentAnimation string         // Имя текущей анимации
}
