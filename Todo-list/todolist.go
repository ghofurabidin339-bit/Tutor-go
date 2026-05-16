package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var todoList []string
	var command string
	fmt.Println("Selamat datang di aplikasi Todo List!")
	for {
		fmt.Println("\nPerintah yang tersedia: add, list, remove, exit")
		fmt.Print("Masukkan perintah: ")
		fmt.Scanln(&command)
		switch command {
		case "add":
			var item string
			fmt.Print("Masukkan item yang ingin ditambahkan: ")
			fmt.Scanln(&item)
			todoList = append(todoList, item)
			fmt.Println("Item berhasil ditambahkan!")
		case "list":
			fmt.Println("Todo List:")
			for i, item := range todoList {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		case "remove":
			var index int
			fmt.Print("Masukkan nomor item yang ingin dihapus: ")
			fmt.Scanln(&index)
			if index >= 1 && index <= len(todoList) {
				todoList = append(todoList[:index-1], todoList[index:]...)
				fmt.Println("Item berhasil dihapus!")
			} else {
				fmt.Println("Nomor item tidak valid!")
			}
		case "exit":
			fmt.Println("Terima kasih telah menggunakan aplikasi Todo List!")
			return
		default:
			fmt.Println("Perintah tidak valid!")
		}
	}
}
