package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	perms := permutations([]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'})
	const numJobs = 8
	ansChan := make(chan string, 1)
	jobsChan := make(chan int, numJobs)
	fmt.Println("number of possible permutations", len(perms))
	go func() {
		for _, possible := range perms {
			jobsChan <- 1
			go task(possible, ansChan, jobsChan)
		}
	}()

	select {
	case ans := <-ansChan:
		fmt.Println("zero2hero+" + ans)
	case <-ctx.Done():
		fmt.Println("can't find solution in 5 minutes...")
		return
	}

}

func task(possibleInput []byte, ans chan<- string, jobsChan chan int) {
	sha256.New()
	sum := sha256.Sum256(append([]byte("zero2hero+"), possibleInput...))
	if sum[0] == 0 && sum[1] == 0 {
		fmt.Println(sum)
		fmt.Println([]byte(possibleInput))
		ans <- string(possibleInput)
	}
	<-jobsChan
}

func permutations(arr []byte) [][]byte {
	var helper func([]byte, int)
	res := [][]byte{}

	helper = func(arr []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
