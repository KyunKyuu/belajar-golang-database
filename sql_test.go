package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "INSERT INTO customer(id,name) VALUES('teguh','teguh')"
	_,err := db.ExecContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil ditambahkan")
}

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "SELECT id,name FROM customer"
	rows,err := db.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
	}

}