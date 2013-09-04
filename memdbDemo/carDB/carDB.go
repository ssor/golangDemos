package carDB

import (
	"fmt"
)

func New_db() DbCar {
	return make(DbCar)
}
func (db *DbCar) Add_car_to_db(car_id string) *car {
	list := make([]Point, point_list_cap)
	db_temp := *db
	db_temp[car_id] = &car{id: car_id, point_list: list}
	return db_temp[car_id]
}
func (db_p *DbCar) getCar(car_id string) *car {
	db := *db_p
	car := db[car_id]
	if car == nil {
		return db.Add_car_to_db(car_id)
	} else {
		return car
	}
}
func (db DbCar) Print_car_list() {
	for k, _ := range db {
		fmt.Println("car list -> ", k)
	}
}
func (db_p *DbCar) Print_point_list(car_id string) {
	db := *db_p
	car := db[car_id]
	fmt.Println("point list => ", car.point_list)
}
func (db *DbCar) Add_point(p Point) {
	car := db.getCar(p.Id)
	car.Add_point(p)
}
func (db *DbCar) Get_latest_point(car_id string) (Point, bool) {
	car := db.getCar(car_id)
	//fmt.Println("Get_latest_point => ", car)
	return car.Get_latest_point()
}

//------------------------------------------------------------------------------

func (car *car) Add_point(p Point) {

	for i := 1; i < point_list_cap; i++ {
		car.point_list[i-1] = car.point_list[i]
	}
	car.point_list[point_list_cap-1] = p

	//length := len(car.point_list)
	//if length >= point_list_cap {
	//	for i := 1; i < length; i++ {
	//		car.point_list[i-1] = car.point_list[i]
	//	}
	//	car.point_list[length-1] = p
	//} else {
	//	car.point_list[length] = p
	//}
}
func (car car) Get_latest_point() (Point, bool) {

	list := car.point_list
	p := list[point_list_cap-1]
	nilP := Point{}
	if p == nilP {
		return Point{}, false
	} else {
		return p, true
	}
}
