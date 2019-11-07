package main

import (
	"fmt"
	"strconv"
	"sync"
)

var myBalance = &balance{amount: 50, currency: "EUR"}

type balance struct {
	amount   float64
	currency string
	// mu       sync.Mutex
	// RWMutex est plus interrasant dans cas car il permet d'avoir
	// plusieur reader en meme temps mais 1 seul writer
	mu sync.RWMutex
}

func (b *balance) Add(amount float64) {
	b.mu.Lock()
	b.amount += amount
	b.mu.Unlock()
}

func (b *balance) Display() string {
	// on appel RLock au lieu de Lock afin de permettre plusieurs reader
	b.mu.RLock()
	// b.mu.Lock()
	// cette instruction sera utilisé apres le return
	// ca permet d'executer le strconv avec les valeur de b et unlock
	// le mutex apres l'utilisation de ces variables
	// ainsi on previens le risque de race condition
	defer b.mu.RUnlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}

func main() {
	fmt.Println(myBalance.Display())
}
