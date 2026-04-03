/* 题目 7 */
// package main

// import "fmt"

// func main() {
// 	// 括号里：形参的数据类型，外面的是返回数据类型
// 	result := func(x int) int {
// 		return x * x
// 	}(10) // 括号里的10就是实参，直接调用

// 	fmt.Println("x的平方为：", result)

// }
/* 题目 12 */
// package main
// import "fmt"
// func main(){
// 	var n int;
// 	fmt.Scan(&n)
// 	res := func(x int)int{
// 		if x == 0  {
// 			return 1
// 		}
		
// 		sum := 1;
// 		for i := 1; i<= x;i ++{
// 			sum *= i
// 		}
// 		return sum
// 	}(n)
// 	fmt.Println("阶乘结果是：" ,res)
// }

/* 题目 15 */
// package main
// import "fmt"
// func main(){
// 	var n float64
// 	fmt.Scan(&n)
// 	res := func(x float64) float64{
// 		if x<0{
// 			return 0
// 		}else if x >= 0 && x < 5{
// 			return x
// 		}else if x >= 5&& x < 10{
		
// 			return 3*x - 5
// 		}else{
// 			return 0.5 * x - 2
// 		}
// 	}(n)
// 	fmt.Println("输出为：" ,res)
// }

/* 题目 16 */
// Go 语言中 ...int 是可变参数的语法糖，本质是切片
// package main

// import "fmt"

// // mul 函数：接收一个 int 切片，计算所有元素的乘积
// func mul(nums []int) int {
// 	product := 1
// 	// 遍历切片，累乘
// 	for _, num := range nums {
// 		product *= num
// 	}
// 	return product
// }

// func main() {
// 	// 调用示例：传入不同长度的切片，实现“可变参数”效果
// 	fmt.Println("mul([]int{2, 3}) =", mul([]int{2, 3}))       // 2*3=6
// 	fmt.Println("mul([]int{1, 2, 3, 4}) =", mul([]int{1, 2, 3, 4})) // 1*2*3*4=24
// 	fmt.Println("mul([]int{5}) =", mul([]int{5}))             // 5
// 	fmt.Println("mul([]int{}) =", mul([]int{}))               // 空切片返回1
// }


/* 题目 18 */
/*
标准导入用原名
别名导入起外号
点导入 不用名
下划线 只执行init*/

// 要只执行包的 init() 函数（不使用包内其他函数），需要使用 匿名导入（import _ "包路径"）。
// import (
// 	// 1. 常规导入
// 	"fmt"
// 	// 2. 别名导入，给包起外号
// 	importStr "strings"
// 	// 3. 点导入，直接用函数，不用写包名
// 	. "time"
// 	// 4. 匿名导入 匿名导入：只执行 init，不使用包
// 	_ "image/png"
// )//常见导包方式



/* 题目 19 */
// package main
// import "fmt"
// func visit(lt []int, f func(int)) {
//     for _, val := range lt {
//         f(val)
//     }
// }
// func main() {
//     //使用匿名函数打印切片中的元素
//     visit([]int{2, 1, 4, 3}, func(val int) {
//         fmt.Println(val)
//     })
// } // 2 1 4 3

/* 题目 20 */
/*defer 后面的函数参数，会在定义时 立刻执行
defer 本身，会在函数结束前才执行*/
// package main
// import "fmt"
// func tracing(s string) string {
// 	fmt.Println("Entering:", s)
// 	return s
// }

// func untracing(s string) {
// 	fmt.Println("Leaving:", s)
// }

// func f() {
// 	defer untracing(tracing("f"))//最后执行
// 	fmt.Println("in f()")
// }

// func g() {
// 	defer untracing(tracing("g"))
// 	fmt.Println("in g()")
// 	f()
// }

// func main() {
// 	g()
// }
/*
Entering: g
in g()
Entering: f
in f()
Leaving: f
Leaving: g
*/

/* 题目 21 */
/*一个函数，记住并访问它外面的变量，即使这个函数在它原来的作用域以外执行。*/
/*有一个内部函数
内部函数用了外部函数的变量
外部函数已经结束了，但内部函数还能继续用那个变量*/
// package main

// import "fmt"

// // evenGenerator 返回一个“生成偶数”的闭包函数
// func evenGenerator() func() int {
	
// 	num := 0 // 外部变量
// 	// 返回匿名函数，这个函数会记住 num 的值
// 	return func() int {
// 		num += 2  //内部函数用到了外部变量
// 		return num
// 	}
// }

// func main() {
// 	// 创建一个偶数生成器实例
// 	nextEven := evenGenerator()

// 	// 多次调用，依次获取下一个偶数
// 	fmt.Println(nextEven()) // 2
// 	fmt.Println(nextEven()) // 4
// 	fmt.Println(nextEven()) // 6
// 	fmt.Println(nextEven()) // 8

// 	// 再创建一个新的生成器，状态独立
// 	anotherEven := evenGenerator()
// 	fmt.Println(anotherEven()) // 2（新实例，从2开始）
// }