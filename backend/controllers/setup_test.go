package controllers

import (
	"os"
	"testing"

	"github.com/carlosarraes/qcmeback/models"
)

var app App

func TestMain(m *testing.M) {
	app.DB = &models.MockDB{}
	os.Exit(m.Run())
}
