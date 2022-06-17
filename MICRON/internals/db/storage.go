package db

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	metrics "./MICRON/internals/metric_model"
	"fmt"
)

type Storage struct {
	Data *pgxpool.Pool
}

type Saver interface{
	Add(storage *Storage)
}

func checker(err error) {
	if err != nil {
		log.Errorln(err)
	}
}

func Metric_storage(pgx *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.Data = pgx
	return storage
}

func (storage *Storage) Add(cont context.Context, metrics metrics.Metric) {
	sql := "INSERT INTO metrics(name, value) VALUES ($1, $2)"
	_, err := storage.Data.Exec(cont, sql, metrics.Name, metrics.Value)
	checker(err)
}

func (storage *Storage) List(cont context.Context, Filter string) []metrics.Metric {
	args := make([]interface{}, 0)
	var metrics []metrics.Metric
	sql := "SELECT id, timestamp, name, value FROM metrics"
	if Filter != "" {
		sql += " WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", Filter))
	}
	err := pgxscan.Select(cont, storage.Data, &metrics, sql, args...)

	checker(err)

	return metrics
}
