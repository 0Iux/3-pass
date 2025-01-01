package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id, name string, private bool) *Bin {
	return &Bin{
		id:        id,
		name:      name,
		private:   private,
		createdAt: time.Now(),
	}

}

type BinList struct {
	bins []Bin
}

func (binList BinList) addBin(bin *Bin) {
	binList.bins = append(binList.bins, *bin)
}
