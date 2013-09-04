// demo project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type string_map map[string][]string
type Point struct {
	CarID, Time, Longitude, Latitude string

	Direction int
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//p1 := Point{"car1", "2013-06-01 12:12:12", "lng1", "lat1"}
	//b, _ := json.Marshal(p1)
	//fmt.Println(string(b))
	strjson := `{"Longitude":"418788021.23999995","Latitude":"143072387.64000002","Time":"2013-06-12 10:49:18","CarID":"J001"}`
	var p Point
	err2 := json.Unmarshal([]byte(strjson), &p)
	if err2 == nil {
		fmt.Println(p)
	} else {
		fmt.Println("json parse error => ", err2)
	}

	//fmt.Println("lis")
	//http.HandleFunc("/", dispatchRoutine)    //设置访问的路由
	//err := http.ListenAndServe(":8088", nil) //设置监听的端口
	//if err != nil {
	//	fmt.Println("ListenAndServe: ", err)
	//}

}

func dispatchRoutine(w http.ResponseWriter, r *http.Request) {
	go response(w, r)
}
func response(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "failed")
		return
	} else {
		fmt.Fprintf(w, string(body))
	}
}
