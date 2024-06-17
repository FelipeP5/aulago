package main
import "fmt"

func main(){
	menu := map[string]float64{
		"pizza": 20.00,
		"suco" : 12.50,
		"x-tudo" : 22.90,
	}
	fmt.Println(menu["pizza"])

	for k, v := range menu{
		fmt.Println(k, "_", v)
	}
}