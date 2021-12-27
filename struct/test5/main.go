package main


type car1 struct {
	name string
}

type car2 struct {
	car1
	year int
}

func (c *car1) GetName() string {
	return "222"
}

func (c *car2) GetName() string {
	return "333"
}

func main() {

}