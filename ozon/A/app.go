package main

import (
    //"fmt"
    "os"
	"bufio"
	"log"
	"strconv"
)

func main(){
    input, err := os.Open("input-201.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var numbers[]int
	for scanner.Scan() {
		var myInt, _ = strconv.Atoi(scanner.Text())
		numbers = append(numbers, myInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	uniques := unique(numbers)

	output, err := os.Create("input-201.a.txt")
	myWriter := bufio.NewWriter(output)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer output.Close()

	for _, number := range uniques {
		myWriter.WriteString(strconv.Itoa(number))    // запись строки
		myWriter.WriteString("\n")   // перевод строки
	}
	myWriter.Flush()
}

func unique(ar []int) []int{
	var uniques []int
	for i := 0; i < len(ar); i++ {
		is_removed := false
		for j := 0; j < len(uniques); j++ {
			if equals(uniques, j, ar[i]) {
				uniques = remove(uniques, j)
				is_removed = true
				break
			}
		}
		if !is_removed {
			uniques = append(uniques, ar[i])
		}
	}
	return uniques
}

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func equals(s []int, j int, value int) bool {
    return s[j] == value
}