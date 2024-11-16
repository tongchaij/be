package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type dataType [][]int

var DATA dataType
var maxRow int
var cache map[string]int

func getMaxPath(row int, col int) int {
	//fmt.Printf("getMaxPath(%d,%d)\n", row, col)
	cacheIdx := fmt.Sprintf("%d,%d", row, col)

	var maxData int
	cacheMaxData, exists := cache[cacheIdx]
	if exists { // calculated maxData exists in cache
		maxData = cacheMaxData
	} else { // find new maxData
		data := DATA[row][col] // data : current Node Data
		if row >= maxRow-1 {   // row is last row : maxData is current Node data
			maxData = data
		} else { // calculate new maxData
			leftData := getMaxPath(row+1, col)
			rightData := getMaxPath(row+1, col+1)

			if rightData > leftData {
				maxData = data + rightData
			} else {
				maxData = data + leftData
			}
		}
		cache[cacheIdx] = maxData
	}

	return maxData
}

func main() {

	fileName := "hard.json"
	fmt.Printf("File Name:%s\n", fileName)

	file, err := os.ReadFile("files/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &DATA)
	if err != nil {
		log.Fatal(err)
	}

	//	fmt.Println(DATA)

	maxRow = len(DATA)
	fmt.Printf("Max Row:%d\n", maxRow)

	cache = make(map[string]int)

	fmt.Printf("Max path:%d\n", getMaxPath(0, 0))
}
