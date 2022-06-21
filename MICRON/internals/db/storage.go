package db

import (
	"context"
	"fmt"

	metrics "./metric_model"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const (
	query = "SELECT id, timestamp, name, value FROM metrics"
	response = "INSERT INTO metrics(name, value) VALUES ($1, $2)"
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

func NewStorage(pgx *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.Data = pgx
	return storage
}

func (storage *Storage) Add(cont context.Context, metrics metrics.Metric) {
	_, err := storage.Data.Exec(cont, response, metrics.Name, metrics.Value)
	checker(err)
}

func (storage *Storage) List(cont context.Context, filters *metrics.Filters) []metrics.Metrics {
	var mhp *MHelper
	args := make([]interface{}, 0)
	str := query
	var metrics []metrics.Metrics
	switch {
	case mhp.MH.namehelper():
		str += " WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", filters.Name))
	case mhp.MH.beginhelper() && mhp.MH.endhelper(): 
		filters.Duration = filters.End.Sub(filters.Begin)
		str += " WHERE duration > $1"
		args = append(args, filters.Duration)
	}
	err := pgxscan.Select(cont, storage.Data, &metrics, str, args...)

	checker(err)

	return metrics
}

