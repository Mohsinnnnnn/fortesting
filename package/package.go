package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type block struct {
	data     string
	hashs    string
	previous *block
}

func main() {
	myfirstblock := block{data: "Genesisblock"}
	fmt.Println(myfirstblock.data)
	h1 := sha256.New()
	h1.Write([]byte(myfirstblock.data))
	mysecondblock := block{data: "Second Block", hashs: hex.EncodeToString(h1.Sum(nil)), previous: &myfirstblock}
	fmt.Println(mysecondblock)
	h2 := sha256.New()
	h2.Write([]byte(mysecondblock.data))
	mythirdblock := block{data: "Third Block", hashs: hex.EncodeToString(h2.Sum(nil)), previous: &mysecondblock}
	fmt.Println(mythirdblock)
	//h3 := sha256.New()
	//h3.Write([]byte(mythirdblock.data))
}
