package models

const CoordinatesTableName = "coordinates"

type Coordinates struct {
	ID        string  `db:"id"`
	Location  string  `db:"location"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`
}

func (c Coordinates) TableName() string {
	return CoordinatesTableName
}
