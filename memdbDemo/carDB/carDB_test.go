// testcarDB
package carDB

import (
	"testing"
)

var db DbCar

func TestNew_db(t *testing.T) {
	db = New_db()
	if db == nil {
		t.FailNow()
	}
}
func TestAdd_car_to_db(t *testing.T) {
	if db != nil {
		db.Add_point(Point{"car1", "2013-06-01 12:12:12", "lng1", "lat1"})
		db.Add_point(Point{"car1", "2013-05-02 12:12:12", "lng3", "lat3"})
		db.Add_point(Point{"car1", "2013-06-02 12:12:12", "lng2", "lat2"})

		p0, ok := db.Get_latest_point("car1")
		if ok == true {
			if p0.Time != "2013-06-02 12:12:12" {
				t.Fail()
			}
		} else {
			t.Fail()
		}

		_, ok1 := db.Get_latest_point("car2")
		if ok1 == true {
			t.Fail()
		}
	}
}
