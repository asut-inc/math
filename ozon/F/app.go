package main

import (
    //"fmt"
    "os"
	"bufio"
	"log"
	"strconv"
	"strings"
	//"time"
	"sort"
)

func main(){
    input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024*64)

	var target int
	var numbers[]int
	var i = 0;
	for scanner.Scan() {
		if i == 1 {
			arr := strings.Split(scanner.Text(), " ")
			for j := 0; j < len(arr); j++ {
				var otherInt, _ = strconv.Atoi(arr[j])
				numbers = append(numbers, otherInt)
			}
		}
		if i == 0 {
			var myInt, _ = strconv.Atoi(scanner.Text())
			target = myInt
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//timer1 := time.Now()
	result := solve(numbers, target)
	//timer2 := time.Now()
	//fmt.Println("Script Time:", timer2.Sub(timer1));
 
	output, err := os.Create("output.txt")
	myWriter := bufio.NewWriter(output)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer output.Close()

	myWriter.WriteString(strconv.Itoa(result))
	myWriter.WriteString("\n")
	
	myWriter.Flush()
}

func solve(ar []int, target int) int{
	sort.Ints(ar)

	k := sort.Search(len(ar), func(k int) bool { return ar[k] >= target })

	if k > 0 && k < len(ar) {
		for i := 0; i <= k; i++ {
			if ar[i] >= target {
				continue
			}
			for j := k; j >= 0; j-- {
				if i != j && ar[i] + ar[j] == target {
					return 1
				}
			}
		}
	} else {
		return 0
	}

	return 0
}