package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))

}

func HighestRank(nums []int) int {
	tempMap := make(map[int]int)

	for _, v := range nums {
		tempMap[v] = tempMap[v] + 1
	}

	needNum := 0
	numCount := 0

	fmt.Println(tempMap)

	for k, v := range tempMap {

		if v > numCount {
			needNum = k
			numCount = v
		}

		if v == numCount {
			if k > needNum {
				needNum = k
				numCount = v
			}

		}

	}

	return needNum

}

func HighAndLow(in string) string {

	min, max := 0, 0

	splitedStr := strings.Split(in, " ")

	for i, v := range splitedStr {

		n, _ := strconv.Atoi(v)

		if i == 0 {
			min = n
			max = n
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

	}

	return fmt.Sprintf("%d %d", max, min)

}

func FindOdd(seq []int) int {

	tempMap := make(map[int]int)

	for _, v := range seq {
		tempMap[v] = tempMap[v] + 1

	}

	for k, v := range tempMap {
		if v%2 == 0 {
			continue
		}

		return k

	}

	return 0

}

func Accum(s string) string {
	fStr := ""

	for i, v := range s {
		sv := string(v)

		fStr += strings.ToUpper(sv)

		if i != 0 {
			fStr += strings.Repeat(strings.ToLower(sv), i)

		}

		if i == len(s)-1 {
			break
		}

		fStr += "-"

	}

	return fStr
}

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func WorkersRunner() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()

	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)

	}
	fmt.Println(time.Since(now))
}

func processParallel(in <-chan int, out chan<- int, workersCount int) {

	go func() {
		time.Sleep(5 * time.Second)
		close(out)
	}()

	for range workersCount {
		go func() {
			worker(in, out)
		}()
	}

}

func worker(in <-chan int, out chan<- int) {

	for val := range in {

		select {
		case out <- processData(val):
		default:
			fmt.Println("DEFAULT")
			return

		}

	}

}

func Capitalize(st string, arr []int) string {

	strArr := strings.Split(st, "")

	for _, v := range arr {
		if len(st) < v-1 {
			break
		}

		strArr[v] = strings.ToUpper(strArr[v])

	}
	return strings.Join(strArr, "")
}

func predictableTimeWork() {

	done := make(chan struct{})

	go func() {
		randomTimeWork()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Task completed")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout ")
	}

}

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)

}

func LogWriter(ch <-chan int) {

	for v := range ch {

		fmt.Println(v)

	}

}

func Writer() <-chan int {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := range 5 {
			ch <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range 5 {
			ch <- i + 5
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch

}

func Writer2(intCount int) <-chan int {

	ch := make(chan int)

	go func() {
		for v := range intCount {
			ch <- v
		}
		close(ch)
	}()

	return ch

}
func Doubler(ch <-chan int) <-chan int {

	newCh := make(chan int)

	go func() {
		for v := range ch {
			newCh <- v * 2
		}

		close(newCh)

	}()

	return newCh

}
func Reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
