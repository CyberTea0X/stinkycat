package main

import (
	"log"

	"github.com/CyberTea0X/stinkycat/internal"
	"github.com/CyberTea0X/stinkycat/pkg/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Инициализация окна
	rl.InitWindow(800, 600, "Platformer Roguelike")
	rl.SetTargetFPS(60)
	log.Println("here")
	// Загрузка текстуры кота
	catTexture := rl.LoadTexture("assets/catsheet.png")

	// Создание мира ECS
	world := ecs.NewWorld()

	// Создание сущности кота
	catEntity := world.NewEntity()

	// Добавление компонента Transform
	world.AddComponent(catEntity, internal.TransformComponentID, &internal.Transform{
		X: 100, Y: 100,
		Width: 64, Height: 64,
	})

	// Добавление компонента Sprite (первый кадр текстуры)
	world.AddComponent(catEntity, internal.SpriteComponentID, &internal.Sprite{
		Texture: catTexture,
		Frame:   rl.Rectangle{X: 0, Y: 0, Width: 32, Height: 32},
	})
	world.AddSystem(internal.NewRenderSystem())

	// Основной цикл игры
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Обновление систем (рендеринг)
		world.Update()

		rl.EndDrawing()
	}

	// Очистка ресурсов
	rl.UnloadTexture(catTexture)
	rl.CloseWindow()
}
