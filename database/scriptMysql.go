package database

import (
	"context"
	"fmt"
)

func displayAllDataFromDatabase() {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM barang ORDER BY id ASC"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, stok int
		var name string
		err = rows.Scan(&id, &name, &stok)
		if err != nil {
			panic(err)
		}
		fmt.Println("--------------------------------------------")
		fmt.Printf("| Id: %d\t| Name: %s\t| Stok: %d |\n", id, name, stok)
	}
	fmt.Println("--------------------------------------------")
}

func insertData(name string, stok int) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO barang(name, stok) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, name, stok)
	if err != nil {
		panic(err)
	}
	insertId , err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Succes Insert Barang (id: %d)\n", insertId)
}

func updateStok(stok, id int) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "UPDATE barang SET stok= ? WHERE id= ? "
	_, err := db.ExecContext(ctx, script, stok, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Update Stok")
}

func updateNamaBarang(name string, id int) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "UPDATE barang SET name= ? WHERE id= ? "
	_, err := db.ExecContext(ctx, script, name, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Update Nama Barang")
}

func deleteById(id int) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "DELETE FROM barang WHERE id = ?"
	_, err := db.ExecContext(ctx, script, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes Hapus Barang")
}


func GetDisplayAllDataFromDatabase() {
	displayAllDataFromDatabase()
}

func SetInsertData( name string, stok int){
	insertData(name, stok)
}

func SetUpdateData(stok, id int){
	updateStok(stok, id)
}

func SetUpdateNamaBarang(name string, id int){
	updateNamaBarang(name, id)
}

func SetDeleteData(id int){
	deleteById(id)
}