package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Barang struct {
	Nama  string
	Stok  int
	Harga float64
}

const dataFile = "barang.txt"

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		command := args[0]
		switch command {
		case "tambah":
			nama := args[1]
			stok := parseInteger(args[2])
			harga := parseFloat(args[3])
			tambahBarang(nama, stok, harga)
		case "lihat":
			lihatDaftarBarang()
		case "cari":
			keyword := args[1]
			cariBarang(keyword)
		case "tentang":
			tentangAplikasi()
		default:
			fmt.Println("Perintah tidak valid.")
		}
	} else {
		runAplikasi()
	}
}

func initData() []Barang {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return []Barang{}
	}
	var result []Barang
	if err := json.Unmarshal(data, &result); err != nil {
		return []Barang{}
	}
	return result
}

func saveData(data []Barang) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Gagal menyimpan data.")
		return
	}
	if err := ioutil.WriteFile(dataFile, jsonData, 0644); err != nil {
		fmt.Println("Gagal menyimpan data.")
	}
}

func showMainMenu() {
	fmt.Println("===== Selamat datang di Toko PT Maju Kena Mundur Kena =====")
	fmt.Println("1. Tambah Barang Baru")
	fmt.Println("2. Lihat Daftar Barang")
	fmt.Println("3. Cari Barang")
	fmt.Println("4. Tentang Aplikasi")
	fmt.Println("0. Keluar")
}

func tambahBarang(nama string, stok int, harga float64) {
	data := initData()
	barang := Barang{nama, stok, harga}
	data = append(data, barang)
	saveData(data)
	fmt.Printf("Barang %s yang berjumlah %d dan dengan harga %.2f telah berhasil ditambahkan.\n", nama, stok, harga)
}

func lihatDaftarBarang() {
	data := initData()
	if len(data) == 0 {
		fmt.Println("Daftar barang kosong.")
	} else {
		fmt.Println("Daftar Barang:")
		for i, barang := range data {
			fmt.Printf("%d. Nama: %s, Stok: %d, Harga: %.2f\n", i+1, barang.Nama, barang.Stok, barang.Harga)
		}
	}
}

func cariBarang(keyword string) {
	data := initData()
	hasilCari := []Barang{}
	for _, barang := range data {
		if containsIgnoreCase(barang.Nama, keyword) {
			hasilCari = append(hasilCari, barang)
		}
	}
	if len(hasilCari) == 0 {
		fmt.Println("Barang tidak ditemukan.")
	} else {
		fmt.Println("Hasil Pencarian:")
		for i, barang := range hasilCari {
			fmt.Printf("%d. Nama: %s, Stok: %d, Harga: %.2f\n", i+1, barang.Nama, barang.Stok, barang.Harga)
		}
	}
}

func tentangAplikasi() {
	fmt.Println("===== Toko PT Maju Kena Mundur Kena =====")
	fmt.Println("Nama: Putu Widya Rusmananda Yasa.")
	fmt.Println("NIM: 2301020042.")
}

func keluarAplikasi() {
	fmt.Println("Aplikasi telah ditutup.")
	os.Exit(0)
}

func runAplikasi() {
	showMainMenu()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pilih menu (0-4): ")
	choice, _ := reader.ReadString('\n')
	choice = choice[:len(choice)-1]

	switch choice {
	case "1":
		fmt.Print("Nama barang: ")
		nama, _ := reader.ReadString('\n')
		nama = nama[:len(nama)-1]

		fmt.Print("Stok: ")
		stok := readInt(reader)

		fmt.Print("Harga: ")
		harga := readFloat(reader)

		tambahBarang(nama, stok, harga)
		runAplikasi()
	case "2":
		lihatDaftarBarang()
		runAplikasi()
	case "3":
		fmt.Print("Cari barang: ")
		keyword, _ := reader.ReadString('\n')
		keyword = keyword[:len(keyword)-1]

		cariBarang(keyword)
		runAplikasi()
	case "4":
		tentangAplikasi()
		runAplikasi()
	case "0":
		keluarAplikasi()
	default:
		fmt.Println("Pilihan tidak valid.")
		runAplikasi()
	}
}

func parseInteger(input string) int {
	var result int
	_, err := fmt.Sscanf(input, "%d", &result)
	if err != nil {
		return 0
	}
	return result
}

func parseFloat(input string) float64 {
	var result float64
	_, err := fmt.Sscanf(input, "%f", &result)
	if err != nil {
		return 0
	}
	return result
}

func containsIgnoreCase(s, substr string) bool {
	s, substr = strings.ToLower(s), strings.ToLower(substr)
	return strings.Contains(s, substr)
}

func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	return parseInteger(strings.TrimSpace(input))
}

func readFloat(reader *bufio.Reader) float64 {
	input, _ := reader.ReadString('\n')
	return parseFloat(strings.TrimSpace(input))
}
