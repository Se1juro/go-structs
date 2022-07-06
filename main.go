package main

import (
	"fmt"

	"github.com/go-structs/users"
)

func main() {
	daniel := users.User{Id: 2, Name: "Daniel", LastName: "Morales", Age: 22}
	martha := users.User{Id: 1, Name: "Martha", LastName: "Gomez", Age: 20}
	john := users.User{Id: 3, Name: "John", LastName: "Doe", Age: 20}
	nicole := users.User{Id: 4, Name: "Nicole", LastName: "Smith", Age: 21}

	martha.SayHello()
	martha.AddFriends(daniel)
	martha.AddFriends(john)
	martha.AddFriends(nicole)
	fmt.Printf("%+v\n", martha)
	martha.ListFriends()
	martha.RemoveFriend(3)
	fmt.Println("===================")
	martha.ListFriends()
}
