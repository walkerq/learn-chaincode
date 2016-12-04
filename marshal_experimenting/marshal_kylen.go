package main

import "encoding/json"
import "fmt"

func main() {

	type SmartestContract struct {
		creator  string
		approver string
		data     string
		state    int
	}

	scA := SmartestContract{"Walker", "Bob", "aaa", 0}

	scB := SmartestContract{"Alice", "Ed", "bbb", 0}

	scC := SmartestContract{"Bob", "Alice", "ccc", 0}

	slcA := []SmartestContract{scA, scB, scC}

	slcAsBytes, _ := json.Marshal(slcA)
	fmt.Println("abc!")
	fmt.Println(scA)
	fmt.Println(string(slcMarshalled))
	err = stub.PutState("_debug1", slcAsBytes)
}
