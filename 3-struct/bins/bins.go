package bins

import (
	"errors"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("Неверный id")
	}

	if name == "" {
		return nil, errors.New("Неверное имя")
	}

	newBin := &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	return newBin, nil
}

func (vault *BinList) AddBin(bin Bin) {
	vault.Bins = append(vault.Bins, bin)
}