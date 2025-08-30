package main

// NOT NEEDED IF JUST USING: $ go test
import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

// NOT NEEDED IF JUST USING: $ go test
// NEEDED FOR: $ go run hello.go
func main(){
	fmt.Println(Hello("name"))
}