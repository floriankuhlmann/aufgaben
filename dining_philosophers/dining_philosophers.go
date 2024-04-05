package main

import (
	"fmt"
	"sync"
	"time"
)

// Implement the dining philosopher’s problem with the following constraints/modifications.

type Host struct {
	permission chan int
}

type ChopS struct{ sync.Mutex }

type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

var thedinner sync.WaitGroup

func main() {

	// There should be 5 philosophers sharing chopsticks,
	// with one chopstick between each adjacent pair of philosophers.

	// init the chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// init the philosophers
	// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		// Each philosopher is numbered, 1 through 5.
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	// The host allows no more than 2 philosophers to eat concurrently.
	h := Host{make(chan int, 2)}

	// eat now
	for i := 0; i < 5; i++ {
		go philos[i].eat(&h)
	}

	// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	go permit(&h)

	thedinner.Wait()
}

func permit(h *Host) {
	defer thedinner.Done()
	for x := range h.permission {
		x++ // prevent warning
	}
}

func (p Philo) eat(h *Host) {
	defer thedinner.Done()

	// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for i := 0; i < 3; i++ {
		h.permission <- p.number
		p.leftCS.Lock()
		p.rightCS.Lock()

		// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
		// on a line by itself, where <number> is the number of the philosopher.
		fmt.Printf("starting to eat %d  \n", p.number)
		time.Sleep(time.Second)

		// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
		// on a line by itself, where <number> is the number of the philosopher.
		fmt.Printf("finished eating %d  \n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}
