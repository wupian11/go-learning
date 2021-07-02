//demo_6.go

package main

func checkAddress(a int) {
	println(a)  //赋值后的值
	println(&a) //赋值后的内存地址
}

func checkValue(b *int) {
	println(b)  //赋值后的内存地址
	println(*b) //赋值后的值
}

func main() {

	i := 1
	println(i)  //i的值
	println(&i) //i的指针

	checkAddress(i)
	checkValue(&i)
}
