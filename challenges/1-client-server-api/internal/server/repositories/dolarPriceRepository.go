package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/johnldev/go-expert/chanllenge-1/internal/server/services"
)

func init() {

}

type DollarPriceRepository struct {
	db  *sql.DB
	ctx context.Context
}

func (r *DollarPriceRepository) Save(data services.PriceResult) error {
	ctx, cancel := context.WithTimeout(r.ctx, time.Millisecond*10)
	defer cancel()
	// insert
	stmt, err := r.db.PrepareContext(ctx, `INSERT INTO dolar_price ( id, code, codeIn, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, uuid.NewString(), data.Code, data.Codein, data.Name, data.High, data.Low, data.VarBid, data.PctChange, data.Bid, data.Ask, data.Timestamp, data.CreateDate)
	if err != nil {
		return err
	}

	return nil
}

func NewDolarPriceRepository(db *sql.DB, ctx context.Context) *DollarPriceRepository {
	return &DollarPriceRepository{
		db:  db,
		ctx: ctx,
	}
}
