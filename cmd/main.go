package main

import (
	"fmt"
	"strings"
)

func Splitter(str string) [][]string {
	splitFunc := func(r rune) bool {
		return strings.ContainsRune(" 	", r)
	}

	temp := strings.Split(str, "\n")
	slice := make([][]string, len(temp))
	for ind := range temp {
		slice[ind] = strings.FieldsFunc(temp[ind], splitFunc)
	}
	return slice
}

func CounterOfWords(slice [][]string) {
	for ind := range slice {
		fmt.Printf("number of words in the %d row: %d \n", ind, len(slice[ind]))
	}
}
func indexesOf(element int, data []int) []int {
	var indexes []int
	for k, v := range data {
		if element == v {
			indexes = append(indexes, k)
		}
	}
	return indexes
}

func CounterOfUniqueWords(slice [][]string) (int, int, int, []int) {
	var uniqueWordsPerRowCount []int
	wordsCount := make(map[string]int)

	for ind := range slice {
		for _, word := range slice[ind] {
			_, match := wordsCount[word]
			if match {
				wordsCount[word] += 1
			} else {
				wordsCount[word] = 1
			}
		}
		uniqueWordsPerRowCount = append(uniqueWordsPerRowCount, len(wordsCount))
		wordsCount = make(map[string]int)
	}

	if len(uniqueWordsPerRowCount) < 3 {
		fmt.Println("Invalid input")
		return -1, -1, -1, uniqueWordsPerRowCount
	}

	first, second, third := 0, 0, 0
	for ind := range uniqueWordsPerRowCount {
		if uniqueWordsPerRowCount[ind] > first {
			third = second
			second = first
			first = uniqueWordsPerRowCount[ind]
		} else if uniqueWordsPerRowCount[ind] > second {
			third = second
			second = uniqueWordsPerRowCount[ind]
		} else if uniqueWordsPerRowCount[ind] > third {
			third = uniqueWordsPerRowCount[ind]
		}

	}
	return first, second, third, uniqueWordsPerRowCount

}

func PrinterOfIndexesAndCounts(first int, second int, third int, count []int) {
	indexesOfFirst := indexesOf(first, count)
	indexesOfSecond := indexesOf(second, count)
	indexesOfThird := indexesOf(third, count)
	fmt.Println("\nRATING OF ROWS:")
	if first == second && second == third {
		for i := 0; i < len(indexesOfFirst); i++ {
			fmt.Printf("number of unique words in the %d row: %d \n", indexesOfFirst[i], first)
		}
	} else if first == second {
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfFirst[0], first)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfFirst[1], second)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfThird[0], third)
	} else if second == third {
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfFirst[0], first)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfSecond[0], second)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfSecond[1], third)
	} else {
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfFirst[0], first)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfSecond[0], second)
		fmt.Printf("number of unique words in the %d row: %d \n", indexesOfThird[0], third)
	}
}

func main() {
	//It doesn't mean that I can't use fmt.Scan and other scanning functions :)
	driver := Splitter("111 222 333 444 \n555 666 777 888\n888 99976 12356 65432 764567")
	CounterOfWords(driver)
	first, second, third, countUnique := CounterOfUniqueWords(driver)
	PrinterOfIndexesAndCounts(first, second, third, countUnique)
}
