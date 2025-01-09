package bins

import (
	"errors"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	bins []Bin
}

func NewBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	newBin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin, nil
}

func (bl *BinList) AddBin(bin Bin) {
	bl.bins = append(bl.bins, bin)
}
