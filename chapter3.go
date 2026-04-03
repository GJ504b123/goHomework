// // package main

// // import "fmt"

// // func main() {
// // 	// 声明变量：var 变量名 类型
// // 	var name string
// // 	var age int
// // 	// 赋值
// // 	name = "张三"
// // 	age = 18
// // 	fmt.Println(name, age)
// // }
// package main

// import "fmt"

// // 冒泡排序函数
// // arr: 整型切片（Go推荐用切片替代数组，更灵活）
// func bubbleSort(arr []int) {
// 	// 获取数组长度
// 	n := len(arr)
// 	// 外层循环：控制排序轮数，共 n-1 轮
// 	for i := 0; i < n-1; i++ {
// 		// 内层循环：每轮比较相邻元素，末尾i个元素已排好序
// 		for j := 0; j < n-1-i; j++ {
// 			// 升序排序：前一个数 > 后一个数，交换位置
// 			if arr[j] > arr[j+1] {
// 				// Go 语法糖：直接交换，无需临时变量！
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// // 辅助函数：打印数组（测试用）
// func printArray(arr []int) {
// 	// 遍历切片，_ 是匿名变量（忽略索引）
// 	for _, value := range arr {
// 		fmt.Printf("%d ", value)
// 	}
// 	fmt.Println()
// }

// // 主函数：测试冒泡排序
// func main() {
// 	// 测试用无序数组（和C语言版本完全一致）
// 	arr := []int{64, 34, 25, 12, 22, 11, 90}

// 	fmt.Print("排序前的数组：")
// 	printArray(arr) // 打印原数组

// 	bubbleSort(arr) // 调用排序函数

// 	fmt.Print("排序后的数组：")
// 	printArray(arr) // 打印排序结果
// }


/* 题目10 */
// package main

// import "fmt"

// func main() {
//     var x float64
//     fmt.Print("请输入x的值：")
//     fmt.Scan(&x) // 读取用户输入的x

//     var y float64
//     if x < 0 {
//         y = 0
//     } else if x >= 0 && x < 5 {
//         y = x
//     } else {
//         y = 3*x - 5
//     }

//     fmt.Printf("对应的y值为：%.2f\n", y) // 保留2位小数输出
// }

/*题目11 */
// package main

// import "fmt"

// func main() {
// 	// var sum = 0 // 全局可用
//     sum := 0 // 创建变量（累加器）并初始化，只有函数内部可以用，但是更快更简洁，相当与js里的let sum = 0
//     for i := 1; i <= 100; i++ {
//         sum += i // 累加i到sum
//     }
//     fmt.Println("1+2+…+100 =", sum)
// }

/*题目12*/
// package main

// import "fmt"

// func main() {

//     nums := []int{1, 2, 3, 5} //切片
//     for idx, num := range nums { //range:把切片挨个拿出来
// 	//for 索引,当前值:= range 切片/数组{}
//         fmt.Print(num)
//         if idx != len(nums)-1 { // 最后一个数字后不加逗号
//             fmt.Print(",")
//         }
//     }
//     fmt.Println()


// }

/* 题目13 */
// package main
// import "fmt"
// func main (){
// 	max := 0
// 	for i :=0; i <= 200; i ++{
// 		if i % 13 == 0{
// 			max = i;
// 		}
		
// 	}
// 	fmt.Println(max)
// }// 195

/*题目14 */

// package main
// import "fmt"
// func main (){
// 	for i := 1; i <= 100; i ++{
// 		if i % 7 == 0 && i % 5 != 0{
// 			fmt.Println(i)
// 		}
// 	}
// }

/*题目16 */
//abc三位数
// 123 /10 = 3
// package main
// import "fmt"
// func main(){
// 	for i := 100; i < 1000; i++{
// 		c := i % 10
// 		b := i /10 % 10
// 		a := i / 100 % 10
// 		if i == c*c*c + a*a*a + b*b*b{
// 			fmt.Println(i)
// 		}
// 	}
// }

/* 题目17 */
// package main
// import "fmt"
// func main (){
// 	for i := 0; i <=30 ; i++{
// 		if 2 * i + 4 * (30 - i) == 90{
// 			fmt.Printf("鸡有：%d只,兔子有：%d只",i,30-i)
// 			//Printf :打印变量用它，打印纯文字用Println
// 		}
// 	}
// }

/*题目18 */
// package main
// import "fmt"
// func main (){
// 	fmt.Println("请输入两个正整数：")
// 	var x,y int //先声明类型才可以输入！！
// 	fmt.Scan(&x,&y)
// 	sum := x*y
// 	if x < y{
// 		tmp := x
// 		x = y
// 		y = tmp
// 	} // x是大数

	
// 	for y != 0{
// 		tmp := x % y
// 		x = y
// 		y = tmp
// 	}
// 		gcd := x
// 		lcm := sum /x
// 		fmt.Printf("最大公约数%d,最小公倍数%d",gcd,lcm)

// }

/* 题目19 */
// package main
// import (
// 	"fmt"
// 	// "math"
// )
// func main(){
// 	for i := 1; i <= 1000; i ++{
// 		tmp := 0
// 		for j := 1; j <= i/2; j ++{
// 			if i % j == 0{
// 			  tmp += j
// 			}
// 		}
// 		if tmp == i{
// 			fmt.Printf("%d 是完数\n",i)
// 		}
// 	}
// }

/**题目20 */

// package main
// import "fmt"
// func main (){
// 	for i:= 1;i<=5;i++{
// 		for j:=0;j < i;j++{
// 			fmt.Print("#") //Print不换行
// 		}
// 		fmt.Print("\n")
// 	}
// }