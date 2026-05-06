package main

import (
	"time"

	"github.com/google/uuid"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	BinList []Bin
}

func main() {

}

func createBin(name string, private bool) *Bin {
	bin := Bin{
		id:        uuid.New().String(),
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	return &bin
}
