package db

import "github.com/SKilliu/grpc-service/db/models"

type CoordinatesQ interface {
	Insert(user models.Coordinates) error
}

type CoordinatesWrapper struct {
	parent *DB
}

func (d *DB) CoordinatesQ() CoordinatesQ {
	return &CoordinatesWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (u *CoordinatesWrapper) Insert(c models.Coordinates) error {
	return u.parent.db.Model(&c).Insert()
}
