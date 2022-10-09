package database_test

import (
	"foodDeliveryAppServer/database"
	"testing"
)

func TestConnectDb(t *testing.T) {

	db := database.ConnectDb()

	if db == nil {
		t.Fail()
	}
}
