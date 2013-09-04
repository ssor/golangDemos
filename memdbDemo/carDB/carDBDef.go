package carDB

const (
	point_list_cap = 3
)

type DbCar map[string]*car

type Point struct {
	Id, Time, Lng, Lat string
}
type car struct {
	id         string
	point_list []Point
}
