package repository

import (
	"errors"
	"main/model/manage"
)

func (r *Repository) GetItemsReplacementRepo(data *[]manage.ReplacementItem) error {
	query := `
	SELECT 
    i.id, 
    i.name, 
    c.name AS category, 
    i.purchase_date, 
    (CURRENT_DATE - i.purchase_date) AS total_usage_days,
    CASE 
        WHEN (CURRENT_DATE - i.purchase_date) > 100 THEN true 
        ELSE false 
    END AS replacement_required
	FROM 
		inventory_items i
	JOIN 
		categories c ON c.id = i.category_id
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return errors.New("Data not found")
	}

	for rows.Next() {
		var data_ manage.ReplacementItem
		if err := rows.Scan(&data_.ID, &data_.Name, &data_.Category, &data_.PurchaseDate, &data_.TotalUsageDays, &data_.ReplacementRequired); err != nil {
			return err
		}

		*data = append(*data, data_)
	}

	return nil
}
