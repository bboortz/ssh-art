package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/calmh/randomart"
)

var r = rand.New(rand.NewSource(rand.Int63()))

func genRandomData(data []byte) {
	r = rand.New(rand.NewSource(r.Int63()))
	r.Read(data)
}

func updateRandomDataOneByte(data []byte) {
	r = rand.New(rand.NewSource(r.Int63()))
	n := r.Intn(16)
	token := make([]byte, 1)
	r.Read(token)
	data[n] = token[0]
}

func updateRandomDataOneBit(data []byte) {
	r = rand.New(rand.NewSource(r.Int63()))
	n := r.Intn(16)
	diff := r.Intn(1)

	if diff == 0 {
		data[n] = data[n] + 1
	} else {
		data[n] = data[n] - 1
	}
}

func main() {

	data := make([]byte, 16)
	/*
		data = []byte{0xde, 0xad, 0xbe, 0xef, 0x20, 0x19, 0x13, 0x37,
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	*/
	genRandomData(data)

	for {
		updateRandomDataOneBit(data)
		art := randomart.GenerateSubtitled(data, "SSH KEY", "RANDOMART")
		fmt.Print("\033[H\033[2J")
		fmt.Println(art)
		time.Sleep(100 * time.Millisecond)
	}
}
