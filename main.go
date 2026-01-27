package main

import "fmt"

func main(){
	done := Done{}
	done.add("Buy Milk")
	done.add("Buy Bread")
	fmt.Printf("%v \n",done)

	done.Delete(0)
	fmt.Printf("%v \n",done)

}