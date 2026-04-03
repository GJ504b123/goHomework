/*方案 1：循环实现*/
package main

import "fmt"



// 空数组检查
func isEmpty(arr []int) bool {
	return len(arr) == 0
}

// 用户输入数组
func inputArray() []int {
	var n int
	fmt.Print("请输入数组长度：")
	fmt.Scan(&n)
	if n <= 0{
		return []int{}
	}
	arr := make([]int, n)
	fmt.Print("请输入数组元素：")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

// 1.1 作用：求数组最大值
func getMaxLoop(arr []int) int { //1.2 输入：int型的数组；输出：int型的数组里的最大值
	max := arr[0]
	for i := 1; i < len(arr); i++ {  
		if arr[i] > max {//1功能：遍历找到最大值 ；如果删除会：找不到最大值，默认返回数组第一个值
			max = arr[i]
		}
	}
	return max // 1.3 有1个返回值，即：int型的数组里的最大值
}

// 2.1 作用：判断数组是否递增
func isIncreaseLoop(arr []int) bool { //2.2 输入：int型的数组；输出：bool型数组是否递增
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] { //2功能：判断数组是否非递增 ；如果删除会：默认数组递增
			return false// 2.3 有1个返回值，即：bool型数组是否递增
		}
	}
	return true
}

// 3.1 作用：求数组元素所有值之和
func getSumLoop(arr []int) int { //3.2 输入：int型的数组；输出：int型的数组里的所有值的和
	sum := 0
	for _, v := range arr { //3功能：将数组中所有数字遍历累加 ；如果删除会：无法得到元素之和
		sum += v
	}
	return sum // 3.3 有1个返回值，即：int型的数组里的所有值的和
}

// 4.1 作用：整个程序的入口
func main() {
	arr := inputArray()//4.2 输入：int型的数组；输出对应的结果
	if isEmpty(arr){
		fmt.Println("数组输入格式有误")  
		return } //错误处理：空数组判断
	fmt.Println("数组：", arr)  // 4.3 无返回值
	fmt.Println("最大值（循环）：", getMaxLoop(arr))
	fmt.Println("是否递增（循环）：", isIncreaseLoop(arr))
	fmt.Println("元素和（循环）：", getSumLoop(arr))
}



/*方案 2：递归实现*/
// package main

// import "fmt"

// // 1.1 作用：求数组最大值
// func getMaxRecursive(arr []int, index int, max int) int { //1.2 输入：整型数组，整型下标，整型最大值 ，输出：整型最大值
// 	if index == len(arr) {
// 		return max //1.3 返回值：有，整型最大值
// 	}
// 	if arr[index] > max {
// 		max = arr[index]
// 	}
// 	return getMaxRecursive(arr, index+1, max)//1.3 返回值：有，整型最大值
// 	//1功能：遍历完数组里的每一项找到最有返回值，bool型是否递增大值；如果删除会：除非数组里只有一个元素，不然报错：没有返回值
// }

// // 2.1 作用：判断递增
// func isIncreaseRecursive(arr []int, index int) bool {// 2.2 输入：整型数组，整型下标 ；输出：bool值是否递增
// 	if index == len(arr)-1 {
// 		return true // 2.3：有返回值，bool型是否递增
// 	}
// 	if arr[index] > arr[index+1] {
// 		return false// 2.3：有返回值，bool型是否递增
// 	}
// 	return isIncreaseRecursive(arr, index+1) // 2.3：有返回值，bool型是否递增
// 	//2功能：没有遍历完继续递归遍历直至到最后一项或者因为不是递增而退出；如果删除会：除非数组里元素小于3且元素递增 ||数组有三个元素，且不是递增，不然报错：没有返回值
// }



// // 3.1 作用：求和
// func getSumRecursive(arr []int, index int) int { // 3.2：输入：整型数组，整型下标；输出：整型最大值
// 	if index == len(arr) {
// 		return 0 //3.3：有返回值，整型最大值
// 	}
// 	return arr[index] + getSumRecursive(arr, index+1)//3.3：有返回值，整型最大值
// 	//3功能：遍历完所有数组元素并累加；如果删除会：除非数组里元素只有一个元素且是0，不然报错：没有返回值
	
// }

// // 4.1 作用：整个程序的入口
// func main() {
// 	arr := []int{1, 3, 5, 7, 9}//4.2 输入：int型的数组；输出对应的结果
// 	fmt.Println("数组：", arr) // 4.3 无返回值
// 	fmt.Println("最大值（递归）：", getMaxRecursive(arr, 0, arr[0]))
// 	fmt.Println("是否递增（递归）：", isIncreaseRecursive(arr, 0))
// 	fmt.Println("元素和（递归）：", getSumRecursive(arr, 0))
// }