package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем канал для передачи данных
	ch := make(chan string)

	// Запускаем горутину, которая отправляет данные в канал через 2 секунды
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Данные из канала"
	}()

	// Используем select для ожидания данных из канала
	fmt.Println("Ожидание данных из канала...")

	ctxP := context.Background()
	ctx, cancel := context.WithTimeout(ctxP, 3*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case msg := <-ch:
				fmt.Println("Получено:", msg)
				return
			case <-ctx.Done():
				fmt.Println("Тайм-аут: данные не получены вовремя")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("Секция default: данные еще не готовы, ждем")
			}
		}
	}(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	// Даем время для завершения горутины
	time.Sleep(3 * time.Second)
}
