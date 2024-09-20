package main

import "fmt"

func main() {
	type Animal struct {
		class  string
		age    int
		gender string
	}
	// order of fields doesn't matter
	teddy := Animal{
		age:    5,
		class:  "mammal",
		gender: "male",
	}

	teddy.age += 1

	fmt.Println(teddy)
	// order of fields matters when not using field names
	leo := Animal{"mammal", 3, "male"}

	fmt.Println(leo)

	var myCar car = car{
		make:   "Ford",
		model:  "Mustang",
		year:   2021,
		engine: engine{8, 450},
	}

	fmt.Println(myCar.getHp()) // 450
}

// deep dive into the struct type
type engine struct {
	cylinders int
	hp        uint16
}

type car struct {
	make  string
	model string
	year  uint16
	engine
}

type vehicle interface {
	getHp() uint16
}

func (c car) getHp() uint16 {
	return c.engine.hp
}
