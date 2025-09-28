package main

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func NewBin(id string, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now().UTC(),
		Name:      name,
	}
}

func main() {
	var bin = NewBin("asjdbn", "bin", false)

	fmt.Println(bin)
}
