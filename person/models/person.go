package person

import "fmt"

type User struct {
	Name string
	Id   int
}

func NewUser(name string, id int) *User {
	return &User{
		Name: name,
		Id:   id,
	}
}

func InitUser(name string, id int) User {
	return User{
		Name: name,
		Id:   id,
	}
}

func (u *User) SetID() {
	u.Id = 1000

}

func (u User) Print() {
	fmt.Println(u.Name, u.Id)
}
