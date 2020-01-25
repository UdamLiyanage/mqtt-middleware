package main

import (
	"context"
	"testing"
)

func TestDatabaseConnect(t *testing.T) {
	res, err := database.CollectionExists(context.TODO(), "devices")
	if err != nil {
		t.Fail()
	}
	if !res {
		t.Fail()
	}
}
