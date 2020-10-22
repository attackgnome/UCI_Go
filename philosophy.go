package main

import (
	"fmt"
	_ "fmt"
	"sync"
)

//ChopS chopstick type
type ChopS struct {
	sync.Mutex
}

//Philo pholosopher type
type Philo struct {
	leftCS, rightCS *ChopS
	label           int
	count           int
}

var wg sync.WaitGroup

func main() {

	//Initialization
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i, 3}
	}

	//Set up channel for the host method
	c := make(chan int, 2)

	//Start the Dining
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(c)
		go host(c)
	}
	wg.Wait()

}

//eat is the method for the philosopher to eat
func (p Philo) eat(c chan<- int) {
	for i := p.count; i > 0; i-- { //philosopher eats 3 times
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.label)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Println("finishing eating ", p.label)
		c <- 1 //send to channel that  philospher is done eating

	}
	wg.Done()
}

//host uses channels to keep the number of philosophers eating at 2 or less
func host(c <-chan int) {

	for {
		first := <-c
		_ = first
		second := <-c
		_ = second
	}
}
