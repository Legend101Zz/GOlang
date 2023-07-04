package main

import (
	"fmt"
	"math"
	"math/rand"
)


func add(x,y int)int{
	return x + y
}

func main(){
const name ="Mrigesh"

	fmt.Printf(" hello %s\n , my fav num is ,%d \n",name, rand.Intn(100));
	fmt.Println(" hello ",name , "my fav num is , \n", rand.Intn(100),"also",math.Pi, "adding", add(10,100))
}