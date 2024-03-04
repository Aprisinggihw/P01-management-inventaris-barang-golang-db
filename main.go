package main

import (
	"bufio"
	"P01-management-inventaris-barang-golang-db/database"
	"fmt"
	"os"
	"strings"
)

func main() {
	var newId int
	var newStok int
	var pilih int
	input := bufio.NewReader(os.Stdin)
	yes := true
	for yes {
		fmt.Println("\n\n------------------------------------")
		fmt.Println("\tMENU MANAGEMENT BARANG")
		fmt.Println("------------------------------------")
		fmt.Println("1. TAMPILKAN BARANG")
		fmt.Println("2. TAMBAH BARANG")
		fmt.Println("3. HAPUS BARANG")
		fmt.Println("4. UPDATE STOK BARANG")
		fmt.Println("5. UPDATE NAMA BARANG")
		fmt.Println("6. KELUAR")
		fmt.Println("------------------------------------")
		fmt.Print("Masukan Pilihan: ")
		fmt.Scanln(&pilih)
		fmt.Print("-----------------------------------\n\n\n")
		switch pilih {
		case 1:
			database.GetDisplayAllDataFromDatabase()
		case 2:
			fmt.Println("\n\n========================================")
			fmt.Println("\tTAMBAH DATA BARANG")
			fmt.Print("nama: ")
			newName, err := input.ReadString('\n') //dibaca sampai newline
			if err != nil {
				fmt.Println("error reading newName", err)
			}
			newName = strings.TrimSpace(newName)
			fmt.Print("stok: ")
			fmt.Scanln(&newStok)
			fmt.Print("========================================\n\n\n")
			database.SetInsertData(newName, newStok)
		case 3:
			var deleteById string
			fmt.Println("\n\n---------------------------------")
			fmt.Print("Masukan Id Barang Yang Akan Dihapus: ")
			fmt.Scanln(&deleteById)
			fmt.Print("-------------------------------------\n\n\n")
			database.SetDeleteData(newId)
		case 4:
			fmt.Println("\n\n---------------------------------")
			fmt.Print("Masukan Id Barang Yang Akan Diupdate Stoknya: ")
			fmt.Scanln(&newId)
			fmt.Print("Stok Baru: ")
			fmt.Scanln(&newStok)
			database.SetUpdateData(newStok, newId)
			fmt.Print("-------------------------------------\n\n\n")
		case 5:
			fmt.Println("\n\n---------------------------------")
			fmt.Print("Masukan Id Barang Yang Akan Diupdate Namanya: ")
			fmt.Scanln(&newId)
			fmt.Print("Nama Baru: ")
			newName, err := input.ReadString('\n') //dibaca sampai newline
			if err != nil {
				fmt.Println("error reading newName", err)
			}
			newName = strings.TrimSpace(newName)
			database.SetUpdateNamaBarang(newName, newId)
			fmt.Print("-------------------------------------\n\n\n")
		case 6:
			fmt.Println("================================")
			fmt.Println("\t\tBYE BYE")
			fmt.Println("================================")
			yes = false
		}
	}

}
