package main

import "fmt"

func main() {
	user1 := &User{
		name: "ali",
		age:  21,
	}

	user2 := &User{
		name: "amir",
		age:  22,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
