package model

type UserModel struct {
	Name string
	Age  int
}

func NewUserModel(name string, age int) *UserModel {
	return &UserModel{Name: name, Age: age}
}

type UserModelPrivate struct {
	Name string
	age  int
}

func NewUserModePrivate(name string, age int) *UserModelPrivate {
	return &UserModelPrivate{Name: name, age: age}
}

type UserDto struct {
	Name string
	Age  int
}
