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