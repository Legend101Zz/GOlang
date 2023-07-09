package main

import (
	"fmt"
)


func add(x,y int)int{
	return x + y
}

func main(){
const name ="Mrigesh"

xs := []string{"hello","world"}

cxs := make([]string,2)
copy(cxs,xs)
xs[0]= "new"


mp := map[string]int {
"mrigesh" :1,
"nunu" : 2, 
}

for k,v := range mp {
	fmt.Println(k,v)
}
	// fmt.Printf(" hello %s\n , my fav num is ,%d \n",name, rand.Intn(100));
		fmt.Println(xs,"\n");
		fmt.Println(cxs,"\n");
		
	

}