package postgresql

/*

This file together with the model, has all the needed methods to interact with the epoch_metrics table of the database

*/

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

var (
	CREATE_BLOCK_ARRIVAL_TABLE = `
		CREATE TABLE IF NOT EXISTS t_block_metrics(
			f_slot INT,
			f_label VARCHAR(100),
			f_timestamp TIME,
			CONSTRAINT PK_SlotAddr PRIMARY KEY (f_slot,f_label));`

	InsertNewBlock = `
		INSERT INTO t_score_metrics (	
			f_slot, 
			f_label, 
			f_timestamp)
		VALUES ($1, $2, $3);`
)

// in case the table did not exist
func (p *PostgresDBService) createBlockMetricsTable(ctx context.Context, pool *pgxpool.Pool) error {
	// create the tables
	_, err := pool.Exec(ctx, CREATE_BLOCK_ARRIVAL_TABLE)
	if err != nil {
		return errors.Wrap(err, "error creating score metrics table")
	}
	return nil
}

func (p *PostgresDBService) InsertNewBlock(slot int, label string, timestamp time.Time) error {

	_, err := p.psqlPool.Exec(p.ctx, InsertNewBlock,
		slot,
		label,
		timestamp)

	if err != nil {
		return errors.Wrap(err, "error inserting row in score metrics table")
	}
	return nil
}
