package main

import (
	"fmt"
	//"time"
)
/*
// 定义一个Phone 接口
type Phone interface {
	call()
}

// 定义一个结构体 NokiaPhone
type NokiaPhone struct {
	
}

func (nokia NokiaPhone) call() {
	
	fmt.Println("Nokia , I am calling u ")

}
// 定义一个结构体 IPhone
type IPhone struct {

}

func (iPhone IPhone) call() {

	fmt.Println("IPhone , I am calling u ")
	
}
*/

/*
const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
)


type Books struct {
	title string
	author string
	subject string
	book_id int
}
*/

func main() {
	/*
	fmt.Println("Hello, World!")
	var bol bool = true
	fmt.Println(bol)
	var a, b, c string = "aaa", "bbb", "ccc"
	fmt.Println(a, b, c)
	var d int
	d = 1
	fmt.Println(d)
	e := "shengluo" //不应该是声明过的变量
	fmt.Println(e)
	f, g, h := 1, 2, 3
	fmt.Println(f, g, h)
	
	const LENGTH int = 10;
	const WIDTH int = 5;
	fmt.Println(LENGTH,WIDTH)
	
	const (
		Unknow = 0
		Femake = 1
		Male = 2
	)
	
	fmt.Println(Femake)
	
	
	fmt.Println(a, b, c) // import 的时候必须使用和 如果不是用那么 会报错 const可以定义在方法之外
	
		
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = "100"
		g
		h = iota
		i
	)
	
	fmt.Println(a, b, c, d, e, f, g, h, i)
	
	var a int = 1
	var ptr *int
	ptr = &a
	fmt.Println(ptr)

	// * 是指针变量 可以保存地址值 没有*的话保存的是一个整数
	// & 是地址值
	
	
	
	str1, str2 := swap("abc", "def")
	
	fmt.Println(str1, str2)
	
	var arr [10] int
	
	
	var a int = 10 // 定义变量
	var ptr *int   // 定义指针变量 不能初始化
	ptr = &a
	
	fmt.Println(a)
	fmt.Println(&a) // 地址
	fmt.Println(ptr) // 地址
	fmt.Println(*ptr) // 使用指针访问值
	
	
	var ptr *int
	fmt.Println(ptr)
	if ptr == nil {
		fmt.Println("空指针")
	}
	
	
	fmt.Println(Books{"Go 语言", "www", "Go 教程",1})
	fmt.Println(Books{title:"Go 语言", author:"aaa", subject:"go", book_id : 1})
	fmt.Println(Books{author:"aaa", title:"Go 语言", subject:"go", book_id : 1})
	
	
	s :=[] int {1, 2, 3}
	fmt.Println(s)
	
	
	var numbers = make([]int ,3 ,5)
	printSlice(numbers)
	
	
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// 原始切片
	fmt.Println("numbers == " , numbers)
	fmt.Println("numbers[1:4] == ", numbers[1:4]) // 注意:含左不含右
	fmt.Println("numbers[:3]", numbers[:3])
	fmt.Println("numbers[4:]", numbers[4:])
		
	printSlice(numbers)
	
	numbers1 := make([]int ,0 ,5)
	printSlice(numbers1)
	
	numbers2 := numbers[:2]
	printSlice(numbers2)
	
	numbers3 := numbers[2:5]
	printSlice(numbers3)
	
	
	var numbers []int
	
	numbers = append(numbers)
	printSlice(numbers)
	
	numbers = append(numbers, 0)
	printSlice(numbers)
	
	numbers = append(numbers, 1)
	printSlice(numbers)
	
	numbers = append(numbers, 0, 1, 2)
	printSlice(numbers)
	
	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	printSlice(numbers1)
	
	copy(numbers1, numbers) // 将numbers 拷贝到 numbers1
	printSlice(numbers1)
	
	
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	
	// range 也可以用在map的键值对上
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Printf("%d -> %s\n",k ,v)
	}
	
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	
	
	// 创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	
	countryCapitalMap ["France"] = "巴黎"
	countryCapitalMap ["Italy"] = "罗马"
	
	for country := range countryCapitalMap {
		fmt.Println(country, countryCapitalMap[country])
	}
	
	// 查看元素是否存在
	capital, ok := countryCapitalMap["Italy"]
	
	if(ok) {
		fmt.Println(capital)
	} else {
		fmt.Println("NO")
	}
	
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	
	// for country := range countryCapitalMap {
	// 	fmt.Println(country, countryCapitalMap[country])
	// }
	
	delete(countryCapitalMap, "France")
	
	for country := range countryCapitalMap {
		fmt.Println(country, countryCapitalMap[country])
	}	
	
	
	var i int = 3
	fmt.Println(Fac(uint64(i)))
	
	
	var sum int = 1
	var count int = 4
	var mean float32
	
	mean = float32(count) / float32(sum) 
	fmt.Printf("%f\n",mean)
	
	
	// 定义一个接口 里面有call()方法 
	var phone Phone 
	
	phone = new(NokiaPhone)//定义 一个Phone接口 里面赋值为Nokia和IPhone
	phone.call()// 调用call
	
	phone = new(IPhone)
	phone.call() // 调用call
	
	
	//============== GO 并发 ===============//
	go say("hello")
	say("world")
		
	
	fmt.Println("123")
	
	
	ch := make(chan int , 2)
	ch <- 1
	ch <- 2
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	*/
	
	c := make(chan int, 10)
	
	go fibonacci(cap(c), c)
	
	for i := range c {
		fmt.Println(i)
	}
	
	
	
	
	
	
	
	
	
	
}	

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)

}

