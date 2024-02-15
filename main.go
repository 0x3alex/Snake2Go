package main

import (
	"fmt"
	"log"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type (
	obj struct {
		x, y int
	}
)

var (
	snake     []obj
	apple     obj
	dir       obj
	dim       = 1000
	rectSize  = 50
	coinSound rl.Sound
)

func updateBodyPos() {
	for i := len(snake) - 1; i >= 1; i-- {
		snake[i].x = snake[i-1].x
		snake[i].y = snake[i-1].y
	}
	snake[0].x += dir.x * rectSize
	snake[0].y += dir.y * rectSize
}

func handleKey() {
	switch {
	case rl.IsKeyDown(rl.KeyA) && dir.x == 0:
		dir.y = 0
		dir.x = -1
		break
	case rl.IsKeyDown(rl.KeyD) && dir.x == 0:
		dir.y = 0
		dir.x = 1
		break
	case rl.IsKeyDown(rl.KeyW) && dir.y == 0:
		dir.y = -1
		dir.x = 0
		break
	case rl.IsKeyDown(rl.KeyS) && dir.y == 0:
		dir.y = 1
		dir.x = 0
		break
	}

}

func checkOutOfBounds() bool {
	head := snake[0]
	return head.x < 0 || head.x >= dim || head.y < rectSize || head.y >= dim
}

func checkForSelfHit() bool {
	head := snake[0]
	for i := 1; i < len(snake); i++ {
		if snake[i].x == head.x && snake[i].y == head.y {
			return true
		}
	}
	return false
}

func checkForAppleHit() {
	head := snake[0]
	if head.x == apple.x && head.y == apple.y {
		handleApple()
		rl.PlaySound(coinSound)
		snake = append(snake, obj{
			x: -1,
			y: -1,
		})
	}
}

func handleApple() {
	x := rand.Intn(dim-rectSize-0) + 0
	y := rand.Intn(dim-rectSize-30) + 30
	for x%50 != 0 {
		x = rand.Intn(dim-rectSize-0) + 0
	}
	for y%50 != 0 {
		y = rand.Intn(dim-rectSize-30) + 30
	}

	log.Printf("Apple at: %d, %d\n", x, y)
	apple.x = x
	apple.y = y
}

func endScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	txt := fmt.Sprintf("Score: %d, Press Enter to exit", len(snake))
	rl.DrawText(txt, 0, 0, 30, rl.White)
	rl.EndDrawing()
}

func main() {
	rl.InitWindow(int32(dim), int32(dim), "Snake2Go")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	defer func() {
		for !rl.IsKeyPressed(rl.KeyEnter) && !rl.WindowShouldClose() {
			endScreen()
		}
	}()

	coinSound = rl.LoadSoundFromWave(rl.LoadWave("sound.wav"))

	dir.x = 1
	handleApple()
	snake = append(snake, obj{
		x: int(dim / 2),
		y: int(dim / 2),
	})
	rl.SetTargetFPS(10)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		handleKey()
		if checkOutOfBounds() || checkForSelfHit() {
			return
		}
		checkForAppleHit()
		updateBodyPos()
		rl.DrawRectangle(int32(apple.x), int32(apple.y), int32(rectSize), int32(rectSize), rl.Red)

		for _, v := range snake {
			rl.DrawRectangle(int32(v.x), int32(v.y), int32(rectSize), int32(rectSize), rl.Green)
		}

		rl.DrawRectangle(0, 0, int32(dim), int32(rectSize), rl.White)
		rl.DrawText(fmt.Sprintf("Score: %d", len(snake)), 0, 0, 30, rl.Black)

		rl.EndDrawing()
	}
}
