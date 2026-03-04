package main

import (
	"fmt"
	"sync"
	"time"
)

// ChopS representa un palillo
type ChopS struct{ sync.Mutex }

// Philo representa a un filósofo
type Philo struct {
	leftChop, rightChop *ChopS
	id                  int
}

// Host controla el permiso para comer (máximo 2 filósofos)
func host(request chan bool, done chan bool) {
	tokens := 2
	for {
		select {
		case <-request:
			for tokens == 0 {
				<-done
				tokens++
			}
			tokens--
		case <-done:
			tokens++
		}
	}
}

func (p Philo) eat(wg *sync.WaitGroup, request chan bool, done chan bool) {
	defer wg.Done()
	for i := 0; i < 3; i++ { // Cada filósofo come solo 3 veces
		
		// 1. Pedir permiso al anfitrión
		request <- true

		// 2. Coger palillos en cualquier orden
		p.leftChop.Lock()
		p.rightChop.Lock()

		// 3. Empezar a comer
		fmt.Printf("empezando a comer %d\n", p.id)
		
		// Simulamos el tiempo de comer
		time.Sleep(time.Millisecond * 100)

		// 4. Terminar de comer
		fmt.Printf("terminando de comer %d\n", p.id)
		
		p.rightChop.Unlock()
		p.leftChop.Unlock()

		// 5. Avisar al anfitrión que terminó
		done <- true
	}
}

func main() {
	// Inicializar palillos y filósofos
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	// Canales para el anfitrión
	request := make(chan bool)
	done := make(chan bool)
	go host(request, done)

	var wg sync.WaitGroup
	wg.Add(5)

	// Lanzar las 5 goroutines de los filósofos
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, request, done)
	}

	wg.Wait()
	fmt.Println("Todos los filósofos han terminado de comer 3 veces.")
}