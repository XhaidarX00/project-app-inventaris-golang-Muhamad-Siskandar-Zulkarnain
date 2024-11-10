package repository

import (
	"errors"
	"main/model/category"
)

func (r *Repository) GetCategoriesRepo(category_ *[]category.Category) error {
	query := "SELECT id, name, description FROM categories ORDER BY id"
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return errors.New("Data not found")
	}

	for rows.Next() {
		var categoryr category.Category
		if err := rows.Scan(&categoryr.ID, &categoryr.Name, &categoryr.Description); err != nil {
			return err
		}

		*category_ = append(*category_, categoryr)
	}

	return nil
}

func (r *Repository) AddCategoryRepo(category *category.Category) error {
	// Query untuk menambahkan kategori baru
	query := `
        INSERT INTO categories 
        (name, description)
        VALUES ($1, $2) RETURNING id
    `

	var categoryID int
	err := r.DB.QueryRow(query, category.Name, category.Description).Scan(&categoryID)
	if err != nil {
		return err
	}

	query = `SELECT id, name, description FROM categories WHERE id = $1`
	err = r.DB.QueryRow(query, categoryID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCategoryByIdRepo(category *category.Category) error {
	query := `SELECT id, name, description FROM categories WHERE id = $1`
	err := r.DB.QueryRow(query, category.ID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) PutCategoryByIdRepo(category *category.Category) error {
	query := `UPDATE categories SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`

	_, err := r.DB.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	r.GetCategoryByIdRepo(category)

	return nil
}

func (r *Repository) DeletCategoryByIdRepo(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
