package main

import "fmt"

const NMAX int = 2023

type Item struct {
	ID       int
	Name     string
	Quantity int
	Category string
}

type Transaction struct {
	Type             string
	ItemID           int
	Quantity         int
	Day, Month, Year int
}

var inventory [NMAX]Item
var itemCount int
var transactions [NMAX]Transaction
var transactionCount int

func main() {
	var check bool = false
	for check != true {
		var choice int
		fmt.Println("*** ------------------------------------------- ***")
		fmt.Println("***           Aplikasi Inventory Barang         ***")
		fmt.Println("***                  Created by                 ***")
		fmt.Println("***       Muh Aqil Rajab H dan Muhammad Faiz    ***")
		fmt.Println("***          Algoritma Pemrograman 2023         ***")
		fmt.Println("*** ------------------------------------------- ***")
		fmt.Println("Aplikasi Inventori Barang")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Catatan Transaksi")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Lihat Barang")
		fmt.Println("0. Keluar")
		fmt.Print("Masukan Opsi Anda: ")
		fmt.Scanln(&choice)
		if choice == 1 {
			addItem()
		} else if choice == 2 {
			editItem()
		} else if choice == 3 {
			deleteItem()
		} else if choice == 4 {
			recordTransaction()
		} else if choice == 5 {
			searchItems()
		} else if choice == 6 {
			printItems()
		} else if choice == 0 {
			fmt.Println("Keluar...")
			check = true
		} else {
			fmt.Println("Masukan salah!")
		}
	}
}

func addItem() {
	if itemCount >= NMAX {
		fmt.Println("Inventori penuh. Tidak bisa menambahkan barang lagi.")
	} else {
		var newItem Item
		fmt.Println("Masukan Detail barang:")
		fmt.Print("ID: ")
		fmt.Scanln(&newItem.ID)
		fmt.Print("Nama: ")
		fmt.Scanln(&newItem.Name)
		fmt.Print("Jumlah Barang: ")
		fmt.Scanln(&newItem.Quantity)
		fmt.Print("Kategori: ")
		fmt.Scanln(&newItem.Category)
		inventory[itemCount] = newItem
		itemCount++
		fmt.Println("Barang berhasil ditambahkan!")
	}
}

func editItem() {
	if itemCount == 0 {
		fmt.Println("Inventori kosong. Tidak ada barang yang bisa diubah.")
	} else {
		fmt.Println("Pilih barang yang akan diubah:")
		for i := 0; i < itemCount; i++ {
			fmt.Printf("%d. ID: %d, Name: %s, Quantity: %d, Category: %s\n", i+1, inventory[i].ID, inventory[i].Name, inventory[i].Quantity, inventory[i].Category)
		}
		var itemNumber int
		fmt.Print("Masukkan nomor barang yang akan diubah: ")
		fmt.Scanln(&itemNumber)
		if itemNumber < 1 || itemNumber > itemCount {
			fmt.Println("Nomor barang tidak valid!")
		} else {
			var newItem Item
			fmt.Println("Masukan detail barang baru:")
			fmt.Print("Nama: ")
			fmt.Scanln(&newItem.Name)
			fmt.Print("Jumlah barang: ")
			fmt.Scanln(&newItem.Quantity)
			fmt.Print("Kategori: ")
			fmt.Scanln(&newItem.Category)
			inventory[itemNumber-1] = newItem
			fmt.Println("Barang berhasil diperbarui!")
		}
	}
}

func deleteItem() {
	if itemCount == 0 {
		fmt.Println("Inventori kosong. Tidak ada barang yang bisa dihapus.")
	} else {
		fmt.Println("Pilih barang yang akan dihapus:")
		for i := 0; i < itemCount; i++ {
			fmt.Printf("%d. ID: %d, Nama: %s, Jumlah barang: %d, Kategori: %s\n", i+1, inventory[i].ID, inventory[i].Name, inventory[i].Quantity, inventory[i].Category)
		}
		var itemNumber int
		fmt.Print("Masukan nomor barang yang akan dihapus: ")
		fmt.Scanln(&itemNumber)
		if itemNumber < 1 || itemNumber > itemCount {
			fmt.Println("Nomor barang tidak valid!")
		} else {
			for i := itemNumber - 1; i < itemCount-1; i++ {
				inventory[i] = inventory[i+1]
			}
			itemCount--
			fmt.Println("Barang berhasil dihapus!")
		}
	}
}

func recordTransaction() {
	if itemCount == 0 {
		fmt.Println("Inventori kosong. Tidak ada barang untuk dicatat transaksinya.")
	} else {
		fmt.Println("Pilih barang untuk transaksi:")
		for i := 0; i < itemCount; i++ {
			fmt.Printf("%d. ID: %d, Nama: %s, Jumlah barang: %d, Kategori: %s\n", i+1, inventory[i].ID, inventory[i].Name, inventory[i].Quantity, inventory[i].Category)
		}
		var itemNumber int
		fmt.Print("Masukan nomor barang: ")
		fmt.Scanln(&itemNumber)
		if itemNumber < 1 || itemNumber > itemCount {
			fmt.Println("Nomor barang tidak valid!")
		} else {
			var newTransaction Transaction
			fmt.Println("Masukan detail transaksi:")
			fmt.Print("Ketik (in/out): ")
			fmt.Scanln(&newTransaction.Type)
			fmt.Print("Jumlah Barang: ")
			fmt.Scanln(&newTransaction.Quantity)
			fmt.Print("Hari: ")
			fmt.Scanln(&newTransaction.Day)
			fmt.Print("Bulan: ")
			fmt.Scanln(&newTransaction.Month)
			fmt.Print("Tahun: ")
			fmt.Scanln(&newTransaction.Year)
			newTransaction.ItemID = inventory[itemNumber-1].ID
			transactions[transactionCount] = newTransaction
			transactionCount++
			// Update the item quantity in the inventory
			itemIndex := itemNumber - 1
			if newTransaction.Type == "in" {
				inventory[itemIndex].Quantity += newTransaction.Quantity
			} else if newTransaction.Type == "out" {
				inventory[itemIndex].Quantity -= newTransaction.Quantity
			}
			fmt.Println("Catatan transaksi berhasil diperbarui!")
		}
	}
}

func searchItems() {
	if itemCount == 0 {
		fmt.Println("Inventori kosong. Tidak ada barang yang bisa dicari.")
	} else {
		fmt.Print("Masukkan kata kunci pencarian: ")
		var keyword string
		fmt.Scanln(&keyword)
		var found bool
		for i := 0; i < itemCount; i++ {
			if inventory[i].Name == keyword || inventory[i].Category == keyword {
				fmt.Printf("ID: %d, Nama: %s, Jumlah barang: %d, Kategori: %s\n", inventory[i].ID, inventory[i].Name, inventory[i].Quantity, inventory[i].Category)
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada barang yang ditemukan.")
		}
	}
}

func printItems() {
	if itemCount == 0 {
		fmt.Println("Inventori kosong.")
	} else {
		fmt.Println("Pilih kategori sorting:")
		fmt.Println("1. Jumlah barang")
		fmt.Println("2. Tanggal transaksi")
		var sortChoice int
		fmt.Print("Masukan pilihan anda: ")
		fmt.Scanln(&sortChoice)
		if sortChoice == 1 {
			sortByQuantity()
			fmt.Println("Barang berhasil disorting berdasarkan jumlah barang:")
		} else if sortChoice == 2 {
			sortTransactionsByDate(transactions[:transactionCount])
			fmt.Println("Barang berhasil disorting berdasarkan tanggal transaksi:")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
		printSortedItems(sortChoice)
	}
}

func printSortedItems(sortChoice int) {
	if sortChoice == 1 {
		for i := 0; i < itemCount; i++ {
			fmt.Printf("ID: %d, Nama: %s, Jumlah barang: %d, Kategori: %s\n", inventory[i].ID, inventory[i].Name, inventory[i].Quantity, inventory[i].Category)
		}
	} else {
		for i := 0; i < transactionCount; i++ {
			transaction := transactions[i]
			item := findItemByID(transaction.ItemID)
			fmt.Printf("Tipe transaksi: %s, Barang: %s, Jumlah barang: %d, Tanggal: %d/%d/%d\n", transaction.Type, item.Name, inventory[i].Quantity, transaction.Day, transaction.Month, transaction.Year)
		}
	}
}

// Algoritma Selection Sort (Descending)
func sortByQuantity() {
	for i := 0; i < itemCount-1; i++ {
		minIndex := i
		for j := i + 1; j < itemCount; j++ {
			if inventory[j].Quantity > inventory[minIndex].Quantity {
				minIndex = j
			}
		}
		swapItems(i, minIndex)
	}
}

// Algoritma Insertion Sort (Ascending)
func sortTransactionsByDate(transactions []Transaction) {
	for i := 1; i < transactionCount; i++ {
		key := transactions[i]
		j := i - 1
		for j >= 0 && compareDates(key, transactions[j]) {
			transactions[j+1] = transactions[j]
			j--
		}
		transactions[j+1] = key
	}
}

func compareDates(t1, t2 Transaction) bool {
	if t1.Year < t2.Year {
		return true
	} else if t1.Year > t2.Year {
		return false
	} else if t1.Month < t2.Month {
		return true
	} else if t1.Month > t2.Month {
		return false
	} else if t1.Day < t2.Day {
		return true
	} else if t1.Day > t2.Day {
		return false
	} else {
		return false
	}
}

func swapItems(i, j int) {
	temp := inventory[i]
	inventory[i] = inventory[j]
	inventory[j] = temp
}

func findItemByID(id int) Item {
	for i := 0; i < itemCount; i++ {
		if inventory[i].ID == id {
			return inventory[i]
		}
	}
	return Item{}
}
