package users

import "fmt"

type User struct {
	Id       int
	Name     string
	LastName string
	Age      int
	friends  []User
}

func (u User) SayHello() {
	fmt.Println("Hi, my name is", u.Name)
}

func (u *User) AddFriends(friend User) {
	u.friends = append(u.friends, friend)
}

func (u User) ListFriends() {
	for i, f := range u.friends {
		fmt.Printf("%d. %s\n", i+1, f.Name+" "+f.LastName)
	}
}

func (u *User) RemoveFriend(id int) {
	index := u.findFriend(id)
	if index < 0 {
		return
	}
	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

func (u User) findFriend(id int) int {
	for i, f := range u.friends {
		if f.Id == id {
			return i
		}
	}
	return -1
}
