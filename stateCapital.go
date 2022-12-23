package main

import (
	"fmt"
	"encoding/json"
)
func main(){
	StateCapital := make(map[string]string,10)
	StateCapital = map[string]string{ 
		"Maharashtra" : "Mumbai",
		"Punjab" : "Amritsar",
		"Rajasthan" : "jaipur",
		"Madhya Pradesh" : "Bhopal",
		"Tamil nadu" : "Chennai",
		"Gujrat" : "Gandhinagar",
		"Telangana" : "Hyderbad",
		"Orissa" : "Bhbaneswar",
		"Bihar" : "Patna",
		"West Bengal" : "Kolkata",
		"Goa" : "Panji",
	}

	jsonMap, err := json.Marshal(StateCapital)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(jsonMap))

	jsonMap, err = json.MarshalIndent(StateCapital, " ", "\t")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(jsonMap))
	
}