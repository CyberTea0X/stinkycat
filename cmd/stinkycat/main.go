package main

import (
	anim "github.com/CyberTea0X/stinkycat/internal/animation"
	c "github.com/CyberTea0X/stinkycat/internal/components"
	sys "github.com/CyberTea0X/stinkycat/internal/systems"
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Инициализация окна и текстуры
	rl.InitWindow(800, 600, "Platformer Roguelike")
	rl.SetTargetFPS(60)
	catTexture := rl.LoadTexture("assets/catsheet.png")

	// Создание регистра анимаций
	animLib := anim.NewAnimationLibrary()
	frameWidth, frameHeight := 32, 32

	// Добавление анимаций
	animLib.AddAnimation("idle", anim.GenerateAnimationFrames(catTexture, 0, 8, frameWidth, frameHeight))  // Строка 0, 6 кадров
	animLib.AddAnimation("walk", anim.GenerateAnimationFrames(catTexture, 4, 8, frameWidth, frameHeight))  // Строка 1, 6 кадров
	animLib.AddAnimation("jump", anim.GenerateAnimationFrames(catTexture, 18, 8, frameWidth, frameHeight)) // Строка 2, 4 кадра
	// Добавьте другие анимации при необходимости

	// Создание ECS мира
	w := ecs.NewWorld()
	w.AddSystem(sys.NewRenderSystem())
	w.AddSystem(sys.NewAnimationSystem())
	w.AddSystem(sys.NewInputSystem(animLib))

	// Создание сущности кота
	catEntity := w.NewEntity()
	w.AddComponent(catEntity, c.TransformComponentID, &c.Transform{
		X: 100, Y: 100, Width: 256, Height: 256,
	})

	// Установка начальной анимации
	initialFrames, _ := animLib.GetAnimation("idle")
	w.AddComponent(catEntity, c.SpriteComponentID, &c.Sprite{
		Texture: catTexture,
		Frame:   initialFrames[0],
	})

	w.AddComponent(catEntity, c.AnimationComponentID, &c.Animation{
		Frames:       initialFrames,
		FrameTime:    0.1,
		CurrentFrame: 0,
		Timer:        0,
	})

	w.AddComponent(catEntity, c.InputBindingsComponentID, &c.InputBindings{
		JumpKey:  rl.KeyW,
		WalkKeys: []int32{rl.KeyA, rl.KeyD},
	})

	// Основной цикл игры
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		w.Update()
		rl.EndDrawing()
	}

	// Очистка ресурсов
	rl.UnloadTexture(catTexture)
	rl.CloseWindow()
}
