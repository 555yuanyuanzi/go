package main

import "fmt"

type Animal struct {
	Name string
}
type Pet struct {
	Name string
}
type Cat struct {
	Animal
	Pet
}

type UserInfo struct {
	Name string `json:"name"`
}

// 修改结构体的属性
func (this *UserInfo) SetName(name string) {
	this.Name = name
	fmt.Printf("this:%p\n", this)
}
func main() {
	/*
		var animal = Animal{Name: "猫猫"}
		var pet = Pet{Name: "蓝猫"}
		cat := Cat{
			Animal: animal, Pet: pet,
		}
		fmt.Println(cat.Animal.Name, cat.Pet.Name)
	*/
	user := UserInfo{
		Name: "小杰",
	}
	fmt.Printf("user:%p\n", &user)
	user.SetName("瑶瑶")
	fmt.Println(user.Name)
}
