package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var level string
	fmt.Println("Pilih level kesulitan (mudah/sedang/sulit):")
	fmt.Scanln(&level)

	var target, kesempatan int

	switch level {
	case "mudah":
		target = rand.Intn(10) + 1
		kesempatan = 5
	case "sedang":
		target = rand.Intn(50) + 1
		kesempatan = 7
	case "sulit":
		target = rand.Intn(100) + 1
		kesempatan = 10
	default:
		fmt.Println("Level tidak valid. Pilih antara mudah, sedang, atau sulit.")
		return
	}

	fmt.Printf("Tebak angka antara 1 dan %d. Kamu punya %d kesempatan.\n", target*10, kesempatan)
	for i := 0; i < kesempatan; i++ {
		var tebakan int
		fmt.Printf("Kesempatan %d: ", i+1)
		fmt.Scanln(&tebakan)
		if tebakan < target {
			fmt.Println("Terlalu rendah!")
		} else if tebakan > target {
			fmt.Println("Terlalu tinggi!")
		} else {
			fmt.Printf("Selamat! Kamu menebak dengan benar dalam %d kesempatan!\n", i+1)
			return
		}
	}
}
