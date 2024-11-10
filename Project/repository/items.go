package repository

import (
	"database/sql"
	"main/model/manage"
	"main/model/response"
)

func (r *Repository) GetItemsPaginatedRepo(data *response.PaginationResponse) error {
	offset := (data.Page - 1) * data.Limit

	var totalItems int
	query := "SELECT COUNT(*) FROM inventory_items"
	err := r.DB.QueryRow(query).Scan(&totalItems)
	if err != nil {
		return err
	}

	queryMain := `
	SELECT 
		i.id, 
		i.name, 
		c.name AS category, 
		i.photo_url, 
		i.price, 
		i.purchase_date, 
		(CURRENT_DATE - i.purchase_date) AS total_usage_days
	FROM 
		inventory_items i
	JOIN 
		categories c ON c.id = i.category_id
	LIMIT $1 OFFSET $2`

	rows, err := r.DB.Query(queryMain, data.Limit, offset)
	if err != nil {
		return err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(sql.NullString)
	}

	var results []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return err
		}

		rowData := make(map[string]interface{})
		for i, col := range columns {
			val := values[i].(*sql.NullString)
			if val.Valid {
				rowData[col] = val.String
			} else {
				rowData[col] = nil
			}
		}
		results = append(results, rowData)
	}

	if results == nil {
		return err
	}

	// Calculate total pages
	totalPages := (totalItems + data.Limit - 1) / data.Limit

	data.Data = results
	data.StatusCode = 200
	data.TotalItems = totalItems
	data.TotalPages = totalPages
	data.Success = true

	return nil
}

func (r *Repository) GetInventoryItemsByIDRepo(item *manage.Item) error {
	query := `SELECT 
		i.id, 
		i.name, 
		c.name AS category, 
		i.photo_url, 
		i.price, 
		i.purchase_date,
		(CURRENT_DATE - i.purchase_date) AS total_usage_days
	FROM 
		inventory_items i
	JOIN 
		categories c ON c.id = i.category_id
	WHERE i.id = $1`

	err := r.DB.QueryRow(query, item.ID).Scan(&item.ID, &item.Name, &item.CategoryName, &item.PhotoURL, &item.Price, &item.PurchaseDate, &item.TotalUsageDays)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddInventoryItemRepo(item *manage.Item) error {
	query := `
		INSERT INTO inventory_items 
		(name, category_id, photo_url, price, purchase_date) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.DB.QueryRow(query, item.Name, item.Category, item.PhotoURL, item.Price, item.PurchaseDate).Scan(&item.ID)
	if err != nil {
		return err
	}

	err = r.GetInventoryItemsByIDRepo(item)
	if err != nil {
		return err
	}

	item.TotalUsageDays = 0
	return nil
}

func (r *Repository) UpdateInventoryItemByIdRepo(item *manage.Item) error {
	query := `UPDATE inventory_items 
			  SET name = $1, category_id = $2, photo_url = $3, price = $4, purchase_date = $5, updated_at = CURRENT_TIMESTAMP 
			  WHERE id = $6`

	_, err := r.DB.Exec(query, item.Name, item.Category, item.PhotoURL, item.Price, item.PurchaseDate, item.ID)
	if err != nil {
		return err
	}

	err = r.GetInventoryItemsByIDRepo(item)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteInventoryItemByIdRepo(id int) error {
	query := `DELETE FROM inventory_items WHERE id = $1`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
