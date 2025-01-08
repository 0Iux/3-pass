package bins

import (
	"encoding/json"
	"go_pass/storage"
	"time"

	"github.com/fatih/color"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id, name string, private bool) *Bin {
	return &Bin{
		Id:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}

}

type Db interface {
	ByteReader(string) ([]byte, error)
	ByteWriter([]byte)
}

type BinList struct {
	Bins []Bin `json:"bins"`
	Db   Db
}

func (binList *BinList) ToBytes() ([]byte, error) {
	data, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (binList *BinList) AddBin(bin *Bin) {
	binList.Bins = append(binList.Bins, *bin)
	data, err := binList.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	binList.Db.ByteWriter(data)
}

func NewBinLIst() *BinList {
	file, err := storage.ReadFromJson("storage.json")
	if err != nil {
		return &BinList{
			Bins: []Bin{},
		}
	}
	var vault BinList
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}
