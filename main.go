package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	words := []string{"chipi", "chapa", "dubi", "daba"}
	correct := []int{0, 1, 2, 3}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	shuffle := func(slice []int) {
		for i := len(slice) - 1; i > 0; i-- {
			j := rng.Intn(i + 1)
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	index := make([]int, len(correct))
	copy(index, correct)

	for {
		shuffle(index)
		for _, i := range index {
			fmt.Printf("%s %s ", words[i], words[i])
		}
		if slices.Equal(index, correct) {
			break
		}
	}
	fmt.Println("mágico mi dubi dubi boom boom boom boom")
}
