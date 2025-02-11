package main

import (
	"bufio"
	"fmt"
	"game/entity/game"
	"os"
	"strings"
)

/*
код писать в этом файле
наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

var g *game.Game

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/

	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	//var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	// defer out.Flush()

	initGame()

	for {
		command, _ := in.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "выход" {
			break
		}
		if command == "идти домой" {
			command = "идти коридор"
		}
		fmt.Println(handleCommand(command))
	}

}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все комнаты
		если что-то было - оно корректно перезатирается
	*/
	g = game.New(game.CreateWorld(), game.MakeLookAroundHandlers(), game.MakeAfterMoveHandlers())
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/

	return g.HandleCommand(command)
}
