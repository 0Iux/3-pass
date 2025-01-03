package main

import "go_pass/bins"

func main() {
	binList := bins.NewBinLIst()
	new_bin := bins.NewBin("1", "why", true)
	binList.AddBin(new_bin)
}
