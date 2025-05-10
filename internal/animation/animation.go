package animation

import rl "github.com/gen2brain/raylib-go/raylib"

// GenerateAnimationFrames возвращает слайс прямоугольников для анимации
func GenerateAnimationFrames(texture rl.Texture2D, row int, framesCount int, frameWidth, frameHeight int) []rl.Rectangle {
	var frames []rl.Rectangle
	y := float32(row) * float32(frameHeight)

	for i := 0; i < framesCount; i++ {
		x := float32(i) * float32(frameWidth)
		frames = append(frames, rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  float32(frameWidth),
			Height: float32(frameHeight),
		})
	}
	return frames
}
