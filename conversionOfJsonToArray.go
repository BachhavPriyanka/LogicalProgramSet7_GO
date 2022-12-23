package main

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	Name            string
	MobNumber       string
	AlternateMobNum string
	Address         string
	City            string
}

func main() {
	jsonInput := `[{"Name": "Harry","Mob": "9089898989","AlternateMobNum": "7978787878","Address": "Bunglow no.34","City": "Bangalore"},
	{"Name": "Raj","Mob": "9009009009","AlternateMobNum": "-","Address": "Chennai Near XYZ ","City": "Hitech city"},
	{"Name": "Jenny","Mob": "8282828282","AlternateMobNum": "022543200","Address": "Thaane","City": "Mumbai"},
	{"Name": "Diksha","Mob": "8282828282","AlternateMobNum": "022543200","Address": "Lonawala","City": "Pune"},
	{"Name": "Malika","Mob": "999977777","AlternateMobNum": "998837342","Address": "ojhar","City": "Nasik"},
	{"Name": "Gaurav","Mob": "888776587","AlternateMobNum": "022543200","Address": "nizamuddin","City": "Delhi"}]`
	var friend []Friend

	err := json.Unmarshal([]byte(jsonInput), &friend)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("data %+v\n", friend)

}
