package main

import ( //import yang berguna untuk membaca file,mengelola data dalam bentuk json,mengelola string dan berkomunikasi dengan os atau melalui input dan output
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dataFile = "barang.txt"

type Barang struct {
	Nama  string
	Stok  int
	Harga float64
}

func main() {
	args := os.Args[1:] //menyimpan semua argumen selain argumen pertama ke dalam slice
	if len(args) > 0 {  //ketika panjang argumen lebih dari 0
		command := args[0] //menyimpan argumen pertama pada slice argumen
		switch command {
		case "tambah": //apabila pada argumen user terdapat kata "tambah" pada argumen pertama yang diikuti oleh data barang pada argumen kedua, ketiga, dan keempat maka akan menjalankan aksi tambahData.
			nama := args[1]
			stok := args[2]
			harga := args[3]
			tambahBarang(nama, stok, harga)
		case "lihat": //apabila pada argumen user terdapat kata "tambah" pada argumen pertama,  maka akan menjalankan aksi lihatData.
			lihatDaftarBarang()
		case "cari": //apabila pada argumen user terdapat kata "tambah" pada argumen pertama dan diikuti oleh nama barang pada argumen kedua, maka akan menjalankan aksi cariBarang.
			keyword := args[1]
			cariBarang(keyword)
		case "tentang": //apabila pada argumen user terdapat kata "tentang" pada argumen pertama,  maka akan menjalankan aksi tentangAplikasi.
			tentangAplikasi()
		default: //apabila pada argumen user tidak terdapat kata kunci diatas pada argumen pertama, maka akan memunculkan pesan dibawah
			fmt.Println("Perintah tidak valid.")
		}
	} else {
		runAplikasi()
	}
}

func initData() []Barang { //membaca file, melakukan translate dari format json menjadi format biasa(yang bisa dibaca dan ditampilkan dengan baik oleh program) serta mengembalikan hasilnya ke dalam slice dengan tipe data barang(jika berhasil)
	var data []Barang              //menginialisasi
	file, err := os.Open(dataFile) //mencoba untuk mengakses file, apabila gagal maka fungsi akan memberikan nilai return slice kosong
	if err != nil {
		return data
	}
	defer file.Close() //memastikan file agar ditutup setelah selesai digunakan

	scanner := bufio.NewScanner(file) //membaca file dari baris ke baris
	for scanner.Scan() {
		line := scanner.Text()                       //hasil dari membaca text akan disimpan fi variabel line
		var barang Barang                            //membuat variable barang dengan tipe data barang
		err := json.Unmarshal([]byte(line), &barang) //convert json
		if err == nil {                              //jika tidak ada eror maka data dari file akan ditambahkan ke slice data
			data = append(data, barang)
		}
	}

	return data
}

func saveData(data []Barang) { //menyimpan data pada file barang.txt
	file, err := os.Create(dataFile) //membuka file barang.txt
	if err != nil {                  //apabila membuka data itu gagal, maka akan muncul pop up error gagal menyimpan data
		fmt.Println("Gagal menyimpan data.")
		return
	}
	defer file.Close() //memastikan file agar ditutup setelah selesai digunakan

	for _, barang := range data { //mengiterasi semua barang yang ada dalam slice data
		dataJSON, err := json.Marshal(barang) //translate data dalam slice tersebut kedalam format json
		if err == nil {
			file.WriteString(string(dataJSON) + "\n") //data json yang dihasilkan oleh objek barang akan ditulis dalam bentuk string ditambahkan dengan karakter newline (\n) untuk pemisahan setiap data barang.
		}
	}
}

func showMainMenu() { //menampilkan halaman utama menu pada saat program dijalankan
	fmt.Println("===== Selamat datang di Toko PT Maju Kena Mundur Kena =====")
	fmt.Println("1. Tambah Barang Baru")
	fmt.Println("2. Lihat Daftar Barang")
	fmt.Println("3. Cari Barang")
	fmt.Println("4. Tentang Aplikasi")
	fmt.Println("5. Keluar")
}

func tambahBarang(nama, stok, harga string) { //fungsi untuk menambahkan data ke dalam file
	data := initData() // Fungsi initData() mengembalikan data barang dalam bentuk slice dari tipe Barang, dan data ini disimpan dalam variabel data.
	stokInt := 0       // mengubah nilai stok dan harga, yang awalnya dalam format string, menjadi tipe data yang sesuai.
	hargaFloat := 0.0

	fmt.Sscan(stok, &stokInt) // digunakan untuk membaca dan mem-parsing string stok dan harga ke dalam variabel stokInt dan hargaFloat, masing-masing.
	fmt.Sscan(harga, &hargaFloat)

	barang := Barang{Nama: nama, Stok: stokInt, Harga: hargaFloat} //membuat objek barang baru
	data = append(data, barang)                                    //data dari objek barang akan di tambahkan ke slice data
	saveData(data)                                                 //menyimpan data ke dalam file barang.txt sesuai dengan

	fmt.Printf("Barang %s yang berjumlah %d dan dengan harga %.2f telah berhasil ditambahkan.\n", nama, stokInt, hargaFloat)
}

func lihatDaftarBarang() { //fungsi untuk melihat data di dalam file untuk ditampilkan pada program
	data := initData()  // Fungsi initData() mengembalikan data barang dalam bentuk slice dari tipe Barang, dan data ini disimpan dalam variabel data.
	if len(data) == 0 { //apabila jumlah data dari variable data itu 0, maka akan memunculkan notifikasi data kosong
		fmt.Println("Daftar barang kosong.")
	} else {
		fmt.Println("Daftar Barang:")
		for i, barang := range data {
			fmt.Printf("%d. Nama: %s, Stok: %d, Harga: %.2f\n", i+1, barang.Nama, barang.Stok, barang.Harga)
		}
	}
}

func cariBarang(keyword string) {
	data := initData()     // Fungsi initData() mengembalikan data barang dalam bentuk slice dari tipe Barang, dan data ini disimpan dalam variabel data.
	var hasilCari []Barang //membuat slice kosong dengan tipe data barang

	for _, barang := range data { //mengambil setiap data barang yang ada dalam variable data
		if strings.Contains(barang.Nama, keyword) { //apabila nama barang sesuai dengan data yang ada dalam variable data. maka data dari barang tersebut akan dipindahkan ke slice hasilcari
			hasilCari = append(hasilCari, barang)
		}
	}

	if len(hasilCari) == 0 { //apabila slice hasilcari bernilai 0 (tidak ada data dalam slice tersebut) maka akan muncul pesan barang tidak ditemukan
		fmt.Println("Barang tidak ditemukan.")
	} else {
		fmt.Println("Hasil Pencarian:")
		for i, barang := range hasilCari {
			fmt.Printf("%d. Nama: %s, Stok: %d, Harga: %.2f\n", i+1, barang.Nama, barang.Stok, barang.Harga)
		}
	}
}

func tentangAplikasi() { //memunculkan siapa yang membuat program
	fmt.Println("===== Toko PT Maju Kena Mundur Kena =====")
	fmt.Println("Nama: Putu Widya Rusmananda Yasa")
	fmt.Println("NIM: 2301020042")
	fmt.Println("Kelas: Pagi 2")
}

func keluarAplikasi() { //menutup berjalannya program/aplikasi
	fmt.Println("Aplikasi telah ditutup.")
	os.Exit(0)
}

func runAplikasi() {
	showMainMenu()                  //menampilkan menu utama
	fmt.Print("Pilih menu (1-5): ") //memunculkan pilihan

	var choice string
	fmt.Scanln(&choice) //menyimpan pilihan dari user

	switch choice {
	case "1": //apabila user memilih pilihan 1 maka akan dijalankan aksi tambahBarang sesuai dengan data yang diinputkan oleh user
		var nama, stok, harga string
		fmt.Print("Nama barang: ")
		fmt.Scanln(&nama)
		fmt.Print("Stok: ")
		fmt.Scanln(&stok)
		fmt.Print("Harga: ")
		fmt.Scanln(&harga)
		tambahBarang(nama, stok, harga)
		runAplikasi()
	case "2": //apabila user memilih pilihan 2 maka akan dijalankan aksi lihatBarang
		lihatDaftarBarang()
		runAplikasi()
	case "3": //apabila user memilih pilihan 3 maka akan dijalankan aksi cariBarang sesuai dengan nama barang yang diinputkan oleh user
		var keyword string
		fmt.Print("Cari barang: ")
		fmt.Scanln(&keyword)
		cariBarang(keyword)
		runAplikasi()
	case "4": //apabila user memilih pilihan 4 maka akan dijalankan aksi tentangAplikasi
		tentangAplikasi()
		runAplikasi()
	case "5": //apabila user memilih pilihan 5 maka akan dijalankan aksi keluarAplikasi
		keluarAplikasi()
	default: //apabila user memilih pilihan selain angka 1-5 maka akan dijalankan aksi keluarAplikasi
		fmt.Println("Pilihan tidak valid.")
		runAplikasi()
	}
}
