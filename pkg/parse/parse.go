package parse

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func ParsingFile(path string) [][]string {
	var sliceData [][]string
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		data := strings.Split(fileScanner.Text(), ";")
		sliceData = append(sliceData, data)
	}
	if err = fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	return sliceData
}

func ParseFileBit(path string) (uint8, []bool) {
	dataSlice, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %s", err)
	}
	lenData := len(dataSlice)
	var resultSum uint8
	resbul := make([]bool, lenData)
	for i := lenData - 1; i >= 0; i-- {
		if dataSlice[i] == 49 {
			a := math.Pow(2, float64(i))
			resultSum += 1 * uint8(a)
			resbul[i] = true
		} else {
			resbul[i] = false
		}
	}
	return resultSum, resbul
}
