package example

import "fmt"

func ExampleHello() {
	greeting := Hello("Phil")
	fmt.Println(greeting)
	//Output:
	//Hello, Phil
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   false,
		"Janet": true,
		"Susan": true,
		"John":  false,
	}
	Page(checkIns)
	// Unordered Output:
	// Paging Alice; please see the front desk to check in.
	// Paging Eve; please see the front desk to check in.
	// Paging John; please see the front desk to check in.
}
