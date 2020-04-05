package model

import (
	"context"
	"dk/pkg/config"

	"firebase.google.com/go/db"
)

// Firebase ...
type Firebase interface {
	FindAll(reference string) ([]Profile, error)
	Set(data map[string]interface{}) error
	// Update(data map[string]interface{}) error
}

// BootModel ...
type BootModel struct {
	Conf config.Config
	App  *db.Client
}

// NewFirebase ...
func NewFirebase(cfg config.Config) Firebase {
	return &BootModel{Conf: cfg}
}

// FindAll ...
func (model BootModel) FindAll(reference string) ([]Profile, error) {
	var (
		ctx = context.Background()
		raw = []Profile{}
	)

	err := model.App.NewRef(model.Conf.GetString("firebase_reference")).Get(ctx, &raw)
	if err != nil {
		return raw, err
	}

	return raw, nil
}

// Set ...
func (model BootModel) Set(data map[string]interface{}) error {
	return model.App.NewRef(model.Conf.GetString("firebase_reference")).Set(context.Background(), data)
}
