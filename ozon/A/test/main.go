package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

var check int

const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
)

func main() {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var valuesBytes []byte
	valuesInt := make(map[int]struct{})
	valuesIntDub := make(map[int]struct{})
	i := 0
	sliceInt := make([]int, 0, 256)

	for i = range content {
		if content[i] == 10 {
			check, _ = strconv.Atoi(string(content[:i]))
			valuesBytes = content[i+1:]
			break
		}
	}
	content = content[:0]
	i = 0
	n := 0
	for i = range valuesBytes {
		if valuesBytes[i] == 32 {
			v, _ := strconv.Atoi(string(valuesBytes[n:i]))
			if _, ok := valuesInt[v]; ok {
				valuesIntDub[v] = struct{}{}
			} else {
				valuesInt[v] = struct{}{}
			}
			n = i + 1
		}
		if i == len(valuesBytes)-1 {
			if n == i {
				v, _ := strconv.Atoi(string(valuesBytes[i:]))
				if _, ok := valuesInt[v]; ok {
					valuesIntDub[v] = struct{}{}
				} else {
					valuesInt[v] = struct{}{}
				}
			}
			if n != i {
				v, _ := strconv.Atoi(string(valuesBytes[n:i]))
				if _, ok := valuesInt[v]; ok {
					valuesIntDub[v] = struct{}{}
				} else {
					valuesInt[v] = struct{}{}
				}
			}
		}
	}
	valuesBytes = valuesBytes[:0]
	for i = range valuesInt {
		sliceInt = append(sliceInt, i)
	}
	// for case 25 + 25 = 50
	for i = range valuesIntDub {
		sliceInt = append(sliceInt, i)
	}
	var wg sync.WaitGroup
	quit := make(chan struct{})
	go func() {
		wg.Add(1)
		go checkTwo(sliceInt, &wg, quit)
		wg.Wait()
		if !IsClosed(quit) {
			success(quit, false)
			close(quit)
		}
	}()
	<-quit
}

func checkTwo(n []int, wg *sync.WaitGroup, quit chan<- struct{}) {
	defer wg.Done()
	l := len(n)
	if l < 2 {
		return
	}
	right := l - 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range n {
			if i == right {
				continue
			}
			if (n[i] + n[right]) == check {
				success(quit, true)
				return
			}
		}
	}()
	wg.Add(1)
	go checkTwo(n[:right], wg, quit)
}

func success(quit chan<- struct{}, s bool) {
	var b byte = 48
	if s {
		b = 49
	}
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
		return
	}
	_, err = f.Write([]byte{b})
	if err != nil {
		log.Fatalf("failed write to file: %s", err)
		return
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("failed close file: %s", err)
		return
	}
	if s {
		close(quit)
	}
}

func IsClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
