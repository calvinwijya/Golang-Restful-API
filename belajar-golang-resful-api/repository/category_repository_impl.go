package repository

import (
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRespositoryImplement struct {
}

func NewCategoryRepository() CategoryRespository {
	return &CategoryRespositoryImplement{}
}

func (repository *CategoryRespositoryImplement) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
func (repository *CategoryRespositoryImplement) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (repository *CategoryRespositoryImplement) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category = ? where = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
}
func (repository *CategoryRespositoryImplement) FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error) {
	SQL := "select id,name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, CategoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
	} else {
		return category, errors.New("category is not found")
	}
	return category, nil
}
func (repository *CategoryRespositoryImplement) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from category "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
