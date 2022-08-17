package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func maxSum(arr []int) {
	players := arr[0]
	dice := arr[1]

	sumPoint := make([]int, 1)
	sumDice := make([]int, 1)
	newDice := make([]int, 1)
	for p := 1; p <= players; p++ {
		sumPoint = append(sumPoint, 0)
		sumDice = append(sumDice, dice)
		newDice = append(newDice, 0)
		fmt.Println("Player", p, "have", sumDice[p-1], "dice and point", sumPoint[p-1])
	}

	var playerWin int
	var diceWin int
	for turn := 0; turn < 10; turn++ {
		turn++

		fmt.Println("Turn", turn)
		for p := 1; p <= players; p++ {
			if sumDice[p-1] > 0 {
				fmt.Println("player", p)
				var strDice string
				for rollDice := 0; rollDice < sumDice[p-1]; rollDice++ {
					rand.Seed(time.Now().UnixNano())
					getDice := (rand.Intn(6) + 1)
					if strDice == "" {
						strDice = strconv.Itoa(getDice)
					} else {
						strDice = strDice + ", " + strconv.Itoa(getDice)
					}

					if getDice == 6 {
						newDice[p-1] = newDice[p-1] - 1
						sumPoint[p-1] = sumPoint[p-1] + 1
					}

					if getDice == 1 {
						newDice[p-1] = newDice[p-1] - 1
						for a := p; a < players; a++ {
							if sumDice[a] > 0 {
								newDice[a] = newDice[a] + 1
							}
						}

						for b := 0; b < p-1; b++ {
							if sumDice[b] > 0 {
								newDice[b] = newDice[b] + 1
							}
						}
					}
				}
				fmt.Println(newDice)
			}
		}

		var result int
		for p := 1; p <= players; p++ {
			sumDice[p-1] = sumDice[p-1] + sumDice[p-1]
			newDice[p-1] = 0
			if diceWin < sumPoint[p-1] {
				playerWin = p
				diceWin = sumPoint[p-1]
			}

			fmt.Println("player", p, "have", sumDice[p-1], "point", sumPoint[p-1])
			if sumDice[p-1] == 0 {
				result++
			}

		}
		if result == players-1 {
			return
		}
	}

	fmt.Println("player win", playerWin, "with point", diceWin)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int
	for i := 0; i < 2; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 32)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	maxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
