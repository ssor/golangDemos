package main

import (
	"./carDB"
	"fmt"
)

func main() {
	db := carDB.New_db()
	db.Add_car_to_db("car1")
	db.Add_car_to_db("car2")

	db.Add_point(carDB.Point{"car1", "2013-06-01 12:12:12", "lng1", "lat1"})
	db.Add_point(carDB.Point{"car1", "2013-06-02 12:12:12", "lng2", "lat2"})
	db.Add_point(carDB.Point{"car1", "2013-05-02 12:12:12", "lng3", "lat3"})

	db.Print_point_list("car1")

	p0, ok := db.Get_latest_point("car1")
	if ok == true {
		fmt.Println(p0)
	} else {
		fmt.Println("no latest point")
	}
}
