package repository

import (
	"context"
	"fmt"
	"testing"
	traininggolangdatabase "training-golang-database"
	"training-golang-database/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(traininggolangdatabase.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "hartadi@gmail.com",
		Comment: "Hello World",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(traininggolangdatabase.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(traininggolangdatabase.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
