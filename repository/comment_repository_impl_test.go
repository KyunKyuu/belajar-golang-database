package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "Bruh12@gmail.com",
		Comment : "Test Comeent 12342",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T){
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	
	result, err := commentRepository.FindById(ctx, 1)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T){
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	
	result, err := commentRepository.FindAll(ctx)
	if err != nil{
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
	
}