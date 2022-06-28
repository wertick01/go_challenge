package db

import (
	"fmt"
	"log"
	"net/http"
	metrics "./metric_model"
	"testing"
	"time"
	"context"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var expect = metrics.Metrics{
	Name: "pasha",
	Id: 99,
	Value: 15, 
}

func TestAddSuccess(t *testing.T) {
	var cont = context.Background()
	db, mock, err := sqlmock.New()
	check(err, t)

	defer db.Close()

	newstorage := NewStorage(db)
	rows := sqlmock.NewRows([]string{"id", "name", "value"})
	//expect := metrics.Metrics{
	//	Name: "pasha",
	//	Id: 99,
	//	Value: 15, 
	//}

	rows.AddRow(expect.Name, expect.Value)

	mock.ExpectQuery("INSERT INTO metrics(id, name, value) VALUES ($1, $2, $3)").WithArgs(expect.Id, expect.Name, expect.Value).WillReturnResult(rows)
	err1 := newstorage.Add(cont, *expect)
	if err1 != nil {
		t.Errorf("{err1}")
		return
	}

	err2 := mock.ExpectationsWereMet()
	if err2 != nil {
		t.Errorf("{err2}")
		return
	}

	//prerequisite := item.Id != expect.Id && item.Name != expect.Name && item.Value != expect.Value
	//if prerequisite {
	//	t.Errorf("something is wrong between expect: %v and item: %v", expect, item)
	//}

}

func TestAddError(t *testing.T) {
	var cont = context.Background()
	db, mock, err := sqlmock.New()
	check(err, t)

	defer db.Close()

	newstorage := NewStorage(db)
	rows := sqlmock.NewRows([]string{"id", "name", "value"})

	rows.AddRow(expect.Name, expect.Value)

	mock.ExpectQuery("INSERT INTO metrics(id, name, value) VALUES ($1, $2, $3)").WithArgs(expect.Id, expect.Name, 5).WillReturnResult(rows)
	err1 := newstorage.Add(cont, *expect)
	if err1 != nil {
		t.Errorf("{err1}")
		return
	}

	err2 := mock.ExpectationsWereMet()
	if err2 != nil {
		t.Errorf("{err2}")
		return
	}

	//prerequisite := item.Id != expect.Id && item.Name != expect.Name && item.Value != expect.Value
	//if prerequisite {
	//	t.Errorf("something is wrong between expect: %v and item: %v", expect, item)
	//}

}

func check(err error, t *testing.T) {
	if err != nil {
		t.Errorf("{err}")
	}
}