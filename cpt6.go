package main
import "fmt"
// 7 12 13 17 20
// 4 6 8 9 16
// docker
/* 第7题：声明一个 4 行 2 列 的字符串数组 array */
// func main(){
// 	var array [4][2] string //二维数组声明
// 	array[0][0] = "A1"
// 	array[0][1] = "A2"
// 	array[1][0] = "B1"
// 	array[1][1] = "B2"
// 	array[2][0] = "C1"
// 	array[2][1] = "C2"
// 	array[3][0] = "D1"
// 	array[3][1] = "D2"
	
// 	fmt.Println("4行2列数组为：")
// 	for i:= 0; i < 4;i ++{
// 		fmt.Println(array[i])
// 	}

// }
/* 第12题 */
// 在 Go 语言中，获取切片的长度和容量，不需要调用「函数」，而是直接使用内置的关键字 / 内建函数：
// 长度：.len()
// 容量：.cap()


// func main() {
// 	// 声明一个切片
// 	s := make([]int, 3, 5) // 长度3，容量5
// 	fmt.Println("切片s：", s) //[0,0,0]

// 	length := len(s) //3
// 	capacity := cap(s) // 5

// 	fmt.Println("切片长度 len(s) =", length)
// 	fmt.Println("切片容量 cap(s) =", capacity)
// }

/* 第13题 */
// Go 中追加切片需要用 append 函数，并且用 ... 解包第二个切片，把元素逐个追加进去。
// 切片1 = append(切片1, 切片2...)
// func main(){
// 	x:= []int{6,5,3,4}
// 	y := []int{3,1,2}
// 	x = append(x,y...)
// 	for i:=0;i<len(x);i ++{
// 		fmt.Printf("%d\n",x[i])
// 	}
// }

/* 第17题 */
// func main() {
// 	// 定义 map numsMap
// 	numsMap := map[string]int{
// 		"one":   1,
// 		"two":   10,
// 		"three": 3,
// 	}
// 	fmt.Println("删除前 numsMap：", numsMap)

// 	// 核心代码：使用 delete 内建函数删除键为 "two" 的键值对
// 	delete(numsMap, "two")

// 	fmt.Println("删除后 numsMap：", numsMap)
// }
/* 第20题 */


// func main() {
// 	// 1. 初始化原列表（切片）lt
// 	lt := []interface{}{"初始元素"} // 用 interface{} 支持混合类型（int + string）
// 	fmt.Println("原列表：", lt)

// 	// 2. 头部添加元素 5：将 5 与原切片拼接，生成新切片
// 	lt = append([]interface{}{5}, lt...)
// 	fmt.Println("头部加5后：", lt)

// 	// 3. 尾部依次添加 b、c、a
// 	lt = append(lt, "b", "c", "a")
// 	fmt.Println("尾部加元素后：", lt)

// 	// 4. 遍历列表，输出元素
// 	fmt.Println("\n遍历列表：")
// 	for i, v := range lt {
// 		fmt.Printf("索引%d：%v\n", i, v)
// 	}
// }

// 题目4: 3
// 题目6: (0,c)(1,b)(2,a) ❌❌格式不对
// 题目6: oc1b2a
// 题目8: [3,1,2,0,0]
// 题目9: 
/*声明二维数组 array := [3][2]int{...}，第一维长度为 3（共 3 行）
len(array) 统计的是第一维的长度（行数），与第二维无关
因此输出为3*/
// 题目16:7