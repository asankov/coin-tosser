package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// Heads is the position indicating the heads side of the coin
	Heads Position = "HEADS"
	// Tails is the position indicating the tails side of the coin
	Tails Position = "TAILS"

	numCoins          = 100_000
	machineIterations = 500_000
)

// Position is the type that shows the position of the coin
type Position string

// Coin is the struct that represents a coin
type Coin struct {
	Position
}

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	coins := createCoins(r)

	h, t, dist := countCoins(coins)
	fmt.Printf("1. Heads: %d, Tails: %d, Dist (H/D): %f\n", h, t, dist)

	iterate(coins, r)
	// iterateInOrder(coins, r)

	h, t, dist = countCoins(coins)
	fmt.Printf("2. Heads: %d, Tails: %d, Dist (H/D): %f \n", h, t, dist)
}

func createCoins(r *rand.Rand) []*Coin {
	coins := make([]*Coin, numCoins)

	for i := 0; i < numCoins; i++ {
		var position Position
		// at the beginning it has equal chances to be heads, or tails
		if r.Intn(2) == 1 {
			position = Heads
		} else {
			position = Tails
		}

		coins[i] = &Coin{Position: position}
	}

	return coins
}

func iterate(coins []*Coin, r *rand.Rand) {
	for i := 0; i < machineIterations; i++ {
		// get a random coin from the pile
		rr := r.Intn(numCoins)
		coin := coins[rr]

		if coin.Position == Tails {
			// if it is tails, it becomes heads
			coin.Position = Heads
		} else {
			// if it is heads, it has equal chances to land as tails, or stay heads
			if r.Intn(2) == 0 {
				coin.Position = Tails
			}
		}
	}
}

func iterateInOrder(coins []*Coin, r *rand.Rand) {
	for _, coin := range coins {
		// get the coin in order
		if coin.Position == Tails {
			// if it is tails, it becomes heads
			coin.Position = Heads
		} else {
			// if it is heads, it has equal chances to land as tails, or stay heads
			if r.Intn(2) == 0 {
				coin.Position = Tails
			}
		}
	}
}

func countCoins(coins []*Coin) (heads, tails int, distribution float64) {
	for _, c := range coins {
		if c.Position == Heads {
			heads++
		} else {
			tails++
		}
	}

	return heads, tails, float64(heads) / float64(heads+tails) * 100
}
