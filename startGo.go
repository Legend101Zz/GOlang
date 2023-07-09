package main

import (
	"fmt"
)


func sum (unlimited ...int)int{

	fmt.Println ("No. are :",unlimited)
n := 0 
	for _,v := range unlimited {
n += v
	}

	return n
}

type person struct {

	name string 
	age int
	home string}

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
"mrigesh" :0,
"nunu" : 2, 
}

for k,v := range mp {
	fmt.Println(k,v)
}


v,ok := mp["mrigesh"]
if(ok){
	fmt.Println(v,ok)
}else{
fmt.Println("key not vfound")
}


p1:= person {
	name : "Mrigesh",
	age : 20 ,
	home : "Kullu",
}
// fmt.Printf(" hello %s\n , my fav num is ,%d \n",name, rand.Intn(100));
		// fmt.Println(xs,"\n");
		// fmt.Println(cxs,"\n");gi

		fmt.Println(p1)
		
		x:= sum(1,2,3,3,1,3,33,3,6,7,8,100,111,2,2,1,1,1,1,3,3,34,5,6,8,9)
	
fmt.Println("sum is :-", x)
}