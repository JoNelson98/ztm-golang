//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

var items = []Item{
	{name: "item1", tag: true},
	{name: "item2", tag: true},
	{name: "item3", tag: true},
	{name: "item4", tag: false},
}

func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(stuff []Item) {
	fmt.Println("checking out...")
	for i := range stuff {
		deactivate(&stuff[i].tag)
	}
}

func main() {
	fmt.Println("initial items", items)
	checkout(items)
	activate(&items[0].tag)
	fmt.Println("After items", items)
}
