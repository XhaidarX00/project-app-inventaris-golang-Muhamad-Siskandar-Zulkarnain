package repository

import (
	"main/model/investment"
)

func (r *Repository) GetTotalInvesmentsRepo(invesment *investment.Investment) error {
	query := `
	SELECT 
    SUM(price) AS total_investment,
    SUM(price * POWER(1 - depreciated_rate, 
        DATE_PART('month', AGE(DATE_TRUNC('month', CURRENT_DATE), DATE_TRUNC('month', purchase_date)))
    )) AS total_depreciated_value
	FROM inventory_items
	`

	err := r.DB.QueryRow(query).Scan(&invesment.TotalInvestment, &invesment.DepreciatedValue)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTotalInvesmentsByIdRepo(invesment *investment.ItemInvestment) error {
	query := `
	SELECT id, name, price, 
       price * POWER(1 - depreciated_rate, 
       DATE_PART('month', AGE(DATE_TRUNC('month', CURRENT_DATE), DATE_TRUNC('month', purchase_date)))) AS depreciated_value,
       depreciated_rate * 100 AS depreciated_rate
	FROM inventory_items
	WHERE id = $1
	`

	err := r.DB.QueryRow(query, invesment.ItemID).Scan(
		&invesment.ItemID,
		&invesment.Name,
		&invesment.InitialPrice,
		&invesment.DepreciatedValue,
		&invesment.DepreciationRate,
	)

	if err != nil {
		return err
	}

	return nil
}
