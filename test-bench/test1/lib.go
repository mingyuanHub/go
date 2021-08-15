package test1

func Add(x int, y int) (z int) {
	z = x + y
	return
}


type ForTest struct {
	num int ;
}

func (this * ForTest) Loops() {
	for i:=0 ; i  < 10000 ; i++ {
		this.num++
	}
}