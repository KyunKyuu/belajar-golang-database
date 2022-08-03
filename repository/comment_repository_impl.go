package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB:db}
}

func (repository *commentRepositoryImpl)Insert(ctx context.Context, comment entity.Comment)(entity.Comment, error){
	sql := "INSERT INTO comments(email, comment) VALUES(?,?)"
	
	result, err := repository.DB.ExecContext(ctx, sql, comment.Email, comment.Comment)
	if err != nil{
		panic(err)
	}
	id,err := result.LastInsertId()
	if err != nil{
		panic(err)
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl)FindById(ctx context.Context, id int32)(entity.Comment, error){
	sql := "SELECT id,email,comment FROM comments WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sql, id)
	comment := entity.Comment{}
	if err != nil{
		return comment, err
	}
	defer rows.Close()

	if rows.Next(){
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		return comment, errors.New("ID" + strconv.Itoa(int(id)) + "Tidak ditemukan")
	}

}

func (repository *commentRepositoryImpl)FindAll(ctx context.Context)([]entity.Comment, error){
	sql := "SELECT id,email,comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sql)
	
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next(){
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}