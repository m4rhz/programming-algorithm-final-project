package main

import "fmt"

type Jamaah struct {
	ID           int
	Nama         string
	Umur         int
	JenisKelamin string
}

type PaketUmroh struct {
	ID        int
	Nama      string
	Biaya     int
	Fasilitas string
}

const NMAX int = 2023

var TabJamaah [NMAX]Jamaah
var JumlahJamaah int

var TabPaketUmroh [NMAX]PaketUmroh
var JumlahPaketUmroh int

func tambahJamaah(j *Jamaah) {
	if JumlahJamaah < NMAX {
		fmt.Print("Masukkan ID Jamaah: ")
		fmt.Scanln(&j.ID)
		fmt.Print("Masukkan Nama Jamaah: ")
		fmt.Scanln(&j.Nama)
		fmt.Print("Masukkan Umur Jamaah: ")
		fmt.Scanln(&j.Umur)
		fmt.Print("Masukkan Jenis Kelamin Jamaah: ")
		fmt.Scanln(&j.JenisKelamin)
		j.ID = JumlahJamaah + 1
		TabJamaah[JumlahJamaah] = *j
		JumlahJamaah++
		fmt.Println("Jamaah berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas jamaah sudah penuh.")
	}
}

func tambahPaket(p *PaketUmroh) {
	if JumlahPaketUmroh < NMAX {
		fmt.Print("Masukkan ID Paket Umroh: ")
		fmt.Scanln(&p.ID)
		fmt.Print("Masukkan Nama Paket Umroh: ")
		fmt.Scanln(&p.Nama)
		fmt.Print("Masukkan Biaya Paket Umroh: ")
		fmt.Scanln(&p.Biaya)
		fmt.Print("Masukkan Fasilitas Paket Umroh: ")
		fmt.Scanln(&p.Fasilitas)
		p.ID = JumlahPaketUmroh + 1
		TabPaketUmroh[JumlahPaketUmroh] = *p
		JumlahPaketUmroh++
		fmt.Println("Paket Umroh berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas paket umroh sudah penuh.")
	}
}

func ubahJamaah(id int, j Jamaah) {
	var index int
	var stop bool

	index = -1
	for i := 0; i < JumlahJamaah && !stop; i++ {
		if TabJamaah[i].ID == id {
			index = i
			stop = true
		}
	}
	if index != -1 {
		fmt.Print("Masukkan Nama Jamaah baru: ")
		fmt.Scanln(&j.Nama)
		fmt.Print("Masukkan Umur Jamaah baru: ")
		fmt.Scanln(&j.Umur)
		fmt.Print("Masukkan Jenis Kelamin Jamaah baru: ")
		fmt.Scanln(&j.JenisKelamin)

		TabJamaah[index].Nama = j.Nama
		TabJamaah[index].Umur = j.Umur
		TabJamaah[index].JenisKelamin = j.JenisKelamin

		fmt.Println("Data Jamaah berhasil diubah.")
	} else {
		fmt.Println("Jamaah dengan ID tersebut tidak ditemukan.")
	}
}

func ubahPaket(id int, p PaketUmroh) {
	var index int
	var stop bool

	index = -1
	for i := 0; i < JumlahPaketUmroh && !stop; i++ {
		if TabPaketUmroh[i].ID == id {
			index = i
			stop = true
		}
	}
	if index != -1 {
		fmt.Print("Masukkan Nama Paket Umroh baru: ")
		fmt.Scanln(&p.Nama)
		fmt.Print("Masukkan Biaya Paket Umroh baru: ")
		fmt.Scanln(&p.Biaya)
		fmt.Print("Masukkan Fasilitas Paket Umroh baru: ")
		fmt.Scanln(&p.Fasilitas)

		TabPaketUmroh[index].Nama = p.Nama
		TabPaketUmroh[index].Biaya = p.Biaya
		TabPaketUmroh[index].Fasilitas = p.Fasilitas

		fmt.Println("Data Paket Umroh berhasil diubah.")
	} else {
		fmt.Println("Paket Umroh dengan ID tersebut tidak ditemukan.")
	}
}

func hapusJamaah(id int) {
	var index int
	var stop bool

	index = -1
	for i := 0; i < JumlahJamaah && !stop; i++ {
		if TabJamaah[i].ID == id {
			index = i
			stop = true
		}
	}
	if index != -1 {
		for i := index; i < JumlahJamaah-1; i++ {
			TabJamaah[i] = TabJamaah[i+1]
		}
		JumlahJamaah--
	}
}

func hapusPaket(id int) {
	var index int
	var stop bool

	index = -1
	for i := 0; i < JumlahPaketUmroh && !stop; i++ {
		if TabPaketUmroh[i].ID == id {
			index = i
			stop = true
		}
	}
	if index != -1 {
		for i := index; i < JumlahPaketUmroh-1; i++ {
			TabPaketUmroh[i] = TabPaketUmroh[i+1]
		}
		JumlahPaketUmroh--
	}
}

// Sequential Search
func cariJamaah(kriteria string) [NMAX]Jamaah {
	var hasil [NMAX]Jamaah
	var count int

	if kriteria == "ID" {
		var nilai int
		fmt.Print("Masukkan nilai ID yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahJamaah; i++ {
			if TabJamaah[i].ID == nilai {
				hasil[count] = TabJamaah[i]
				count++
			}
		}
	} else if kriteria == "Nama" {
		var nilai string
		fmt.Print("Masukkan nilai Nama yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahJamaah; i++ {
			if TabJamaah[i].Nama == nilai {
				hasil[count] = TabJamaah[i]
				count++
			}
		}
	} else if kriteria == "Umur" {
		var nilai int
		fmt.Print("Masukkan nilai Umur yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahJamaah; i++ {
			if TabJamaah[i].Umur == nilai {
				hasil[count] = TabJamaah[i]
				count++
			}
		}
	} else if kriteria == "JenisKelamin" {
		var nilai string
		fmt.Print("Masukkan nilai Jenis Kelamin yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahJamaah; i++ {
			if TabJamaah[i].JenisKelamin == nilai {
				hasil[count] = TabJamaah[i]
				count++
			}
		}
	}

	var result [NMAX]Jamaah
	for i := 0; i < count; i++ {
		result[i] = hasil[i]
	}

	return result
}

// Sequential Search
func cariPaket(kriteria string) [NMAX]PaketUmroh {
	var hasil [NMAX]PaketUmroh
	var count int

	if kriteria == "Nama" {
		var nilai string
		fmt.Print("Masukkan nilai Nama yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahPaketUmroh; i++ {
			if TabPaketUmroh[i].Nama == nilai {
				hasil[count] = TabPaketUmroh[i]
				count++
			}
		}
	} else if kriteria == "Biaya" {
		var nilai int
		fmt.Print("Masukkan nilai Biaya yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahPaketUmroh; i++ {
			if TabPaketUmroh[i].Biaya == nilai {
				hasil[count] = TabPaketUmroh[i]
				count++
			}
		}
	} else if kriteria == "Fasilitas" {
		var nilai string
		fmt.Print("Masukkan nilai nama fasilitas yang ingin dicari: ")
		fmt.Scanln(&nilai)
		for i := 0; i < JumlahPaketUmroh; i++ {
			if TabPaketUmroh[i].Fasilitas == nilai {
				hasil[count] = TabPaketUmroh[i]
				count++
			}
		}
	}

	if count > 0 {
		fmt.Println("Paket-paket umroh yang sesuai dengan kriteria:")
		for i := 0; i < count; i++ {
			fmt.Printf("Nama: %s\n", hasil[i].Nama)
			fmt.Printf("Biaya: %d\n", hasil[i].Biaya)
			fmt.Printf("Fasilitas: %s\n", hasil[i].Fasilitas)
			fmt.Println()
		}
	} else {
		fmt.Println("Tidak ditemukan paket umroh yang sesuai dengan kriteria.")
	}

	var result [NMAX]PaketUmroh
	for i := 0; i < count; i++ {
		result[i] = hasil[i]
	}

	return result
}

func menu() {
	fmt.Println()
	fmt.Println("Pilih opsi:")
	fmt.Println("1. Tambah Jamaah")
	fmt.Println("2. Tambah Paket Umroh")
	fmt.Println("3. Ubah Jamaah")
	fmt.Println("4. Ubah Paket Umroh")
	fmt.Println("5. Hapus Jamaah")
	fmt.Println("6. Hapus Paket Umroh")
	fmt.Println("7. Cari Jamaah")
	fmt.Println("8. Cari Paket Umroh")
	fmt.Println("9. Tampilkan Jamaah (Selection Sort (Ascending))")
	fmt.Println("10. Tampilkan Paket Umroh")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
}

func tampilkanJamaah() {
	fmt.Println("Data Jamaah:")

	// Implementasi Selection Sort secara ascending untuk mengurutkan umur jamaah
	for i := 0; i < JumlahJamaah-1; i++ {
		minIndex := i
		for j := i + 1; j < JumlahJamaah; j++ {
			if TabJamaah[j].Umur < TabJamaah[minIndex].Umur {
				minIndex = j
			}
		}
		// Menukar posisi jamaah
		TabJamaah[i], TabJamaah[minIndex] = TabJamaah[minIndex], TabJamaah[i]
	}

	// Menampilkan data jamaah
	for i := 0; i < JumlahJamaah; i++ {
		fmt.Printf("ID: %d\n", TabJamaah[i].ID)
		fmt.Printf("Nama: %s\n", TabJamaah[i].Nama)
		fmt.Printf("Umur: %d\n", TabJamaah[i].Umur)
		fmt.Printf("Jenis Kelamin: %s\n", TabJamaah[i].JenisKelamin)
		fmt.Println("-------------------")
	}
}

func tampilkanPaketUmroh() {
	fmt.Println("Data Paket Umroh:")

	// Implementasi Insertion Sort secara descending untuk mengurutkan biaya paket umroh
	for i := 1; i < JumlahPaketUmroh; i++ {
		key := TabPaketUmroh[i]
		j := i - 1
		for j >= 0 && TabPaketUmroh[j].Biaya < key.Biaya {
			TabPaketUmroh[j+1] = TabPaketUmroh[j]
			j--
		}
		TabPaketUmroh[j+1] = key
	}

	// Menampilkan data paket umroh
	for i := 0; i < JumlahPaketUmroh; i++ {
		fmt.Printf("ID: %d\n", TabPaketUmroh[i].ID)
		fmt.Printf("Nama: %s\n", TabPaketUmroh[i].Nama)
		fmt.Printf("Biaya: %d\n", TabPaketUmroh[i].Biaya)
		fmt.Printf("Fasilitas: %s\n", TabPaketUmroh[i].Fasilitas)
		fmt.Println("-------------------")
	}
}

func main() {
	var pilihan int
	var check bool = false

	for check != true {
		menu()
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			var jamaah Jamaah
			tambahJamaah(&jamaah)
		} else if pilihan == 2 {
			var paket PaketUmroh
			tambahPaket(&paket)
		} else if pilihan == 3 {
			var id int
			fmt.Print("Masukkan ID Jamaah yang ingin diubah: ")
			fmt.Scanln(&id)
			var jamaah Jamaah
			ubahJamaah(id, jamaah)
		} else if pilihan == 4 {
			var id int
			fmt.Print("Masukkan ID Paket Umroh yang ingin diubah: ")
			fmt.Scanln(&id)
			var paket PaketUmroh
			ubahPaket(id, paket)
		} else if pilihan == 5 {
			var id int
			fmt.Print("Masukkan ID Jamaah yang ingin dihapus: ")
			fmt.Scanln(&id)
			hapusJamaah(id)
			fmt.Println("Jamaah berhasil dihapus.")
		} else if pilihan == 6 {
			var id int
			fmt.Print("Masukkan ID Paket Umroh yang ingin dihapus: ")
			fmt.Scanln(&id)
			hapusPaket(id)
			fmt.Println("Paket Umroh berhasil dihapus.")
		} else if pilihan == 7 {
			var kriteria string
			fmt.Print("Masukkan kriteria pencarian Jamaah (ID/Nama/Umur/JenisKelamin): ")
			fmt.Scanln(&kriteria)
			result := cariJamaah(kriteria)
			fmt.Println("Hasil pencarian:")
			for i := 0; i < NMAX; i++ {
				if result[i].ID != 0 {
					fmt.Println(result[i])
				}
			}
		} else if pilihan == 8 {
			var kriteria string
			fmt.Print("Masukkan kriteria pencarian Paket Umroh (Nama/Biaya/Fasilitas): ")
			fmt.Scanln(&kriteria)
			result := cariPaket(kriteria)
			fmt.Println("Hasil pencarian:")
			for i := 0; i < NMAX; i++ {
				if result[i].ID != 0 {
					fmt.Println(result[i])
				}
			}
		} else if pilihan == 9 {
			tampilkanJamaah()
		} else if pilihan == 10 {
			tampilkanPaketUmroh()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih. Program selesai.")
			check = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih opsi yang benar.")
		}
		fmt.Println()
	}
}
