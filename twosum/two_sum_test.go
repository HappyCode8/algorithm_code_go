package twosum

//每个单元测试都必须导入testing包，一般测试文件的命名是被测文件_test.go
import (
	"fmt"
	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

// 测试函数的命名类似必须以Test开头，入参必须是 *testing.T，当运行 go test 命令时，go test 会遍历所有的 *_test.go 中符合上述命名规则的函数，
// 然后生成一个临时的 main 包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
func Test_Two_sum(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	//遍历qs数组
	for _, q := range qs {
		//遍历每一个q，包含了数组与结果的约束
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n")
}

func Test_twoSum(t *testing.T) {
	PatchConvey("test", t, func() {
		PatchConvey("testdata case1", func() {
			nums := []int{2, 7, 11, 15}
			target := 9
			got := twoSum(nums, target)
			//断言方法
			So(got, ShouldResemble, []int{0, 1})
		})
	})
}
