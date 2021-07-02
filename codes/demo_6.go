//demo_6.go

package main

/*
 * 在Go语言中声明变量时，将在计算机内存中给它分配一个位置，以便能够存储、修改和获取变量的值。
 * 要获取变量在计算机内存中的地址，可在变量名前加上&字符。
 */
func checkValue(a int) {
	println(a)  //赋值后的值
	println(&a) //赋值后的内存地址
}

/*
 * 指针是Go语言中的一种类型，指向变量所在的内存单元。要声明指针，可在变量名前加上*字符。
 */
func checkPointer(b *int) {
	println(b)  //赋值后的内存地址
	println(*b) //赋值后的值
}

func main() {

	i := 1
	println(i)  //i的值
	println(&i) //i的内存地址

	checkValue(i)
	checkPointer(&i)
}
