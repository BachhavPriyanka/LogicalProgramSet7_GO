package main
import (
	"encoding/json"
	"fmt"
)

func main(){
	jsonAppleRecipie := `{
	"Thinly Sliced Peeled Apples" : "6 Cups",
	"sugar" : "3/4 cup",
	"flour" : "2 tablespoons",
	"cinamon" : "1/4 teaspoon",
	"nutmeg" : "1/8 tablesppon",
	"lemon juice" : "1 tablespoon"}`

	store := map[string]string{}

	json.Unmarshal([]byte(jsonAppleRecipie), &store)
	fmt.Println(store)
}



