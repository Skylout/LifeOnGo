package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Размер поля
const (
	width  = 50
	height = 20
)

//Поле – это двумерный массив логического типа
type Universe [][]bool

/*Инициализация поля – создание одномерного логического слайса размером в константу height, затем заполнение
каждой ячейки логическим слайсом размера width*/
func NewUniverse() (u Universe) {
	u = make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return
}

//Отображение поля
func (ourUniverse Universe) Show() {

	for i := 0; i < len(ourUniverse); i++ {
		for j := 0; j < len(ourUniverse[i]); j++ {
			if ourUniverse[i][j] == true {
				print("*")
			} else {
				print("_")
			}
		}
		print("\n")
	}
	println("/-----------------------/")
}

//инициализация поля
func (ourUniverse Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		ourUniverse.Set(rand.Intn(height), rand.Intn(width), true)
	}
}

//установка состояния для ячейки поля
func (ourUniverse Universe) Set(posX, posY int, state bool) {
	ourUniverse[posX][posY] = state
}

//проверка статуса ячейки. также учтена проверка краев
func (ourUniverse Universe) CheckCellStatus(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return ourUniverse[y][x]
}

//проверка соседних ячеек
func (ourUniverse Universe) CheckNeighbors(x, y int) (result int) {
	result = 0
	for w := -1; w <= 1; w++ {
		for h := -1; h <= 1; h++ {
			if !(w == 0 && h == 0) && ourUniverse.CheckCellStatus(x+h, y+w) {
				result++
			}
		}
	}
	return
}

//следующая итерация ячейки
func (ourUniverse Universe) NextIteration(x, y int) bool {
	numOfNeighbors := ourUniverse.CheckNeighbors(x, y)
	return numOfNeighbors == 3 || numOfNeighbors == 2 && ourUniverse.CheckCellStatus(x, y)
}

//сохранение текущего состояния поля и переход на следующий
func SaveCurrentStateOfTheUniverse(currentState, nextState Universe) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			nextState.Set(h, w, currentState.NextIteration(h, w))
		}
	}
}

func main() {
	kekUniverse := NewUniverse() //основное поле
	lelUniverse := NewUniverse() //поле резервной копии
	kekUniverse.Seed()

	for iteration := 0; iteration <= 300; iteration++ {
		fmt.Printf("Iteration №%d\n", iteration)
		SaveCurrentStateOfTheUniverse(kekUniverse, lelUniverse)
		kekUniverse.Show()
		time.Sleep(time.Second)
		kekUniverse, lelUniverse = lelUniverse, kekUniverse
	}
}
