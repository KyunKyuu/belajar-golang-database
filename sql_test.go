package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
	"time"
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

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows,err := db.QueryContext(ctx, sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		fmt.Println("Email : ", email)
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		fmt.Println("Birth Date : ", birthDate)
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", createdAt)
	}

}

func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sql := "SELECT username FROM user WHERE username ='"+ username +"' AND password = '"+ password +"' LIMIT 1"
	rows,err := db.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
			fmt.Println("Login Berhasil")
	}else {
		fmt.Println("Login Gagal")
	}

}

func TestSqlInjectionSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sql := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows,err := db.QueryContext(ctx, sql, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
			fmt.Println("Login Berhasil")
	}else {
		fmt.Println("Login Gagal")
	}

}

func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "guhhs@gmail.com"
	comment := "test komen"

	sql := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result,err := db.ExecContext(ctx, sql, email, comment)

	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last ID nya adalah", lastId)
}
