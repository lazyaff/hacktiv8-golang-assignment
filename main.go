package main

import (
	"fmt"
	"os"
	"strconv"
)

// define struct
type Siswa struct {
	nama string
	alamat string
	pekerjaan string
	alasan string
}

func find(absen int, siswa []Siswa) {
	index := absen - 1

	// validasi keberadaan data
	if index < 0 || index >= len(siswa) {
		fmt.Printf("Siswa dengan nomor absen %d tidak ditemukan.\n", absen)
		return
	}

	// ambil data dan print hasilnya
	data := siswa[index]
	fmt.Printf("Siswa dengan nomor absen %d adalah %s.\nIa adalah seorang %s yang berasal dari %s.\nAlasan ia mengikuti kelas golang adalah karena %s.\n", absen, data.nama, data.pekerjaan, data.alamat, data.alasan)
}


func main(){
	// inisialisasi data
	var siswa = []Siswa{
		{
			nama:      "Rina",
			alamat:    "Semarang",
			pekerjaan: "Designer",
			alasan:    "memiliki passion dalam bidang back end",
		},
		{
			nama:      "Budi",
			alamat:    "Malang",
			pekerjaan: "System Administrator",
			alasan:    "memiliki keahlian di bidang pemrograman",
		},
		{
			nama:      "Ahmad",
			alamat:    "Medan",
			pekerjaan: "Dosen",
			alasan:    "golang itu keren",
		},
		{
			nama:      "Zain",
			alamat:    "Jogja",
			pekerjaan: "Programmer",
			alasan:    "ingin belajar bahasa baru",
		},
		{
			nama:      "Zaki",
			alamat:    "Surabaya",
			pekerjaan: "Mahasiswa",
			alasan:    "ingin belajar bahasa pemrograman populer",
		},
		{
			nama:      "Siti",
			alamat:    "Surabaya",
			pekerjaan: "Programmer",
			alasan:    "ikut ikutan teman",
		},
	}

	// validasi apakah ada input
	if len(os.Args) != 2 {
		fmt.Println("Masukkan nomor absen")
		os.Exit(0)
	}

	// ambil input sebagai nomor absen
	absen, err := strconv.Atoi(os.Args[1])

	// validasi apakah nomor absen valid
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat")
		os.Exit(0)
	}

	// ambil data sesuai nomor absen
	find(absen, siswa)
	
}