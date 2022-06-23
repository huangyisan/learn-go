package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  27,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	//
	iterator := userCollection.createIterator()

	// hasNext返回bool 可以被for去迭代
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("%v", user)
	}
}
