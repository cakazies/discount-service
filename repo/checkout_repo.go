package repo

import (
	"context"
	"discount-service/infra"
	"discount-service/interfaces"
	"discount-service/models"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var trx *sqlx.Tx

type checkoutRepo struct {
	db *infra.DB
}

func NewCheckoutRepo(db *infra.DB) interfaces.ICheckoutRepo {
	return &checkoutRepo{
		db: db,
	}
}

func (cr *checkoutRepo) BeginTrx(ctx context.Context) (err error) {
	trx, err = cr.db.Write.Beginx()
	return err
}

func (cr *checkoutRepo) RoolBackTrx(ctx context.Context) (err error) {
	err = trx.Rollback()
	return err
}

func (cr *checkoutRepo) CommitTrx(ctx context.Context) (err error) {
	err = trx.Commit()
	return err
}

func (cr *checkoutRepo) RepoGetItemsList(ctx context.Context, sku []string) (output []models.Items, err error) {

	query := `SELECT id, sku, name, price, currency, qty
				FROM items
				WHERE deleted_at IS NULL
				AND sku = ANY($1)`

	err = cr.db.Read.SelectContext(ctx, &output, query, pq.Array(sku))
	return output, err
}

func (cr *checkoutRepo) RepoPromotionActiveList(ctx context.Context) (output []models.PromotionItems, err error) {

	query := `SELECT id, item_id, item_sku, min_qty, free_item, discount, is_cashback, direct_cashback, max_count, detail, is_active, start_date, end_date
				FROM promotion_items
				WHERE deleted_at IS NULL
				AND is_active
				AND NOW() BETWEEN start_date and end_date`

	err = cr.db.Read.SelectContext(ctx, &output, query)
	return output, err
}

func (cr *checkoutRepo) RepoInsertOrder(ctx context.Context, data []models.Orders) (err error) {
	now := time.Now().UTC()
	if len(data) == 0 {
		return fmt.Errorf("data not found")
	}

	var dataQuery []interface{}
	var valuesQuery []string

	query := `INSERT INTO orders (user_id, item_id, item_sku, item_name, item_price, item_qty, discount, total, created_at) VALUES`
	params := `(?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, v := range data {
		dataQuery = append(dataQuery, v.UserID)
		dataQuery = append(dataQuery, v.ItemID)
		dataQuery = append(dataQuery, v.ItemSku)
		dataQuery = append(dataQuery, v.ItemName)
		dataQuery = append(dataQuery, v.ItemPrice)
		dataQuery = append(dataQuery, v.ItemQty)
		dataQuery = append(dataQuery, v.Discount)
		dataQuery = append(dataQuery, v.Total)
		dataQuery = append(dataQuery, now)

		valuesQuery = append(valuesQuery, params)
	}

	addValuesQuery := strings.Join(valuesQuery, ",")

	query = query + " " + addValuesQuery + ";"

	queryExec, args, err := sqlx.In(query, dataQuery...)
	if err != nil {
		return err
	}

	queryExec = cr.db.Write.Rebind(queryExec)

	if trx == nil {
		_, err = cr.db.Write.ExecContext(ctx, queryExec, args...)
	} else {
		_, err = trx.ExecContext(ctx, queryExec, args...)
	}

	return err
}

// RepoPromotionActiveList ...
func (cr *checkoutRepo) RepoUpdateItems(ctx context.Context, data models.Orders) (err error) {
	query := `UPDATE items SET qty = (qty - $1) WHERE id = $2`

	if trx == nil {
		_, err = cr.db.Write.ExecContext(ctx, query, data.ItemQty, data.ItemID)
	} else {
		_, err = trx.ExecContext(ctx, query, data.ItemQty, data.ItemID)
	}
	return err
}
