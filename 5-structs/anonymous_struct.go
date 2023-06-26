package main

/* An anonymous struct is defined without a name and therefore cannot be referenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type: */ 

func main(){
	myCar := struct {
		Make string
		Model string
	  } {
		Make: "tesla",
		Model: "model 3",
	  }
	myCar.Model = "model 2" 
}