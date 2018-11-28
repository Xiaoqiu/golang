package main

import (
	"fmt"
	"os"
	"sort"
)
/**

 */
func main(){
	// #### sort 排序
	//- sort方法对于内置类型有效
	//- in-place sorting, 所以返回同一个slice而不是一个新的对象
	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("strings: ",strs)

	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints: ",ints)

	//- 判断是否已经排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:",s)

	// # sort by function 重写sort接口自定义查询

	//#### panic 异常捕捉
	panic("a problem")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}
}