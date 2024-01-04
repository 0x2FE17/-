package main

import (
	"fmt"
)

func main() {
	//testSlice()
	//first := getInput()
	//second := getInput()
	first := []string{"D", "K", "Q", "J", "10", "9", "9", "9"}
	second := []string{"X", "A", "A", "A", "A", "K", "Q", "J", "10"}

	haveSubmit := make([]string, 0)
	if canWin(first, second, haveSubmit) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getInput() []string {
	// 输入第一副牌,用逗号分隔.D,X,2,A,K,Q,J,10,9,8,7,6,5,4,3
	first := make([]string, 0)
	// 接收输入的牌
	var input string
	fmt.Scanln(&input)
	var temp string
	// 将输入的牌放入first中,按照逗号分隔
	for i := 0; i < len(input); i++ {
		if input[i] == ',' {
			first = append(first, temp)
			temp = ""
			continue
		}
		temp += string(input[i])
	}
	first = append(first, temp)
	return first
}

// 先手一副牌,后手一副牌,看谁赢
func canWin(first, second, haveSubmit []string) bool {
	//fmt.Println("出牌方:", first)
	//fmt.Println("另一方:", second)
	//fmt.Println("上一轮出牌:", haveSubmit)

	if len(first) == 0 {
		return true
	}
	if len(second) == 0 {
		return false
	}
	canSubmit := make([][]string, 0)
	getCandidate(first, make([]string, 0), haveSubmit, 0, &canSubmit)
	if len(canSubmit) == 0 {
		return !canWin(second, first, make([]string, 0))
	}
	for i := 0; i < len(canSubmit); i++ {
		// 递归调用canWin
		if !canWin(second, remove(first, canSubmit[i]), canSubmit[i]) {
			fmt.Println("---------------")
			fmt.Println("出牌方:", first)
			fmt.Println("另一方:", second)
			fmt.Println("上一轮出牌:", haveSubmit)
			fmt.Println("出牌:", canSubmit[i])
			return true
		}
	}
	return false
}

// 从一副牌中去掉另一副牌
func remove(first, second []string) []string {
	m := make(map[string]int)
	for i := 0; i < len(second); i++ {
		m[second[i]]++
	}
	res := make([]string, 0)
	for i := 0; i < len(first); i++ {
		if m[first[i]] == 0 {
			res = append(res, first[i])
		} else {
			m[first[i]]--
		}
	}
	return res
}

func getCandidate(x, candidate, haveSubmit []string, i int, res *[][]string) {
	if i == len(x) {
		if len(candidate)+len(haveSubmit) == 0 {
			return
		}

		// 先验证candidate是否合法,candidate合法则加入res.candidate就是扑克牌的一种排列
		if !isValid(candidate) {
			return
		}
		if !isSameType(candidate, haveSubmit) {
			return
		}
		if !isBigger(candidate, haveSubmit) {
			return
		}
		willAppend := make([]string, len(candidate))
		copy(willAppend, candidate)
		*res = append(*res, willAppend)
		return
	}
	getCandidate(x, candidate, haveSubmit, i+1, res)
	candidate = append(candidate, x[i])
	getCandidate(x, candidate, haveSubmit, i+1, res)
	// 回溯
	//candidate = candidate[:len(candidate)-1]
}

// isBigger 检查candidate是否比haveSubmit大
func isBigger(candidate, haveSubmit []string) bool {
	if len(haveSubmit) == 0 {
		return true
	}
	// 检查是否是同一种类型
	if !isSameType(candidate, haveSubmit) {
		return false
	}
	// 检查candidate是否比haveSubmit大
	// 先找到candidate和haveSubmit的类型
	candidateType := whichValidator(candidate)
	// 检查candidate是否比haveSubmit大
	return comparators[candidateType](candidate, haveSubmit)
}

// DX2AKQJ109876543
var puke = []string{"D", "X", "2", "A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3"}
var validators []func([]string) bool
var comparators []func([]string, []string) bool

func init() {
	validators = append(validators, isValidDanZhi)
	comparators = append(comparators, isBiggerDanZhi)

	validators = append(validators, isSanDaiYi)
	comparators = append(comparators, isBiggerSanDaiYi)

	validators = append(validators, isSanZhang)
	comparators = append(comparators, isBiggerSanZhang)

	validators = append(validators, isLianDui)
	comparators = append(comparators, isBiggerLianDui)

	validators = append(validators, isShunzi)
	comparators = append(comparators, isBiggerShunzi)

	validators = append(validators, isDuiZi)
	comparators = append(comparators, isBiggerDuiZi)

	validators = append(validators, isZhaDan)
	comparators = append(comparators, isBiggerZhaDan)
}
func isBiggerZhaDan(candidate, haveSubmit []string) bool {
	return getIndex(candidate[0]) < getIndex(haveSubmit[0])
}
func isBiggerDuiZi(candidate, haveSubmit []string) bool {
	return len(candidate) == len(haveSubmit) && getIndex(candidate[0]) < getIndex(haveSubmit[0])
}
func isBiggerShunzi(candidate, haveSubmit []string) bool {
	return len(candidate) == len(haveSubmit) && getIndex(candidate[0]) < getIndex(haveSubmit[0])
}
func isBiggerLianDui(candidate, haveSubmit []string) bool {
	return len(candidate) == len(haveSubmit) && getIndex(candidate[0]) < getIndex(haveSubmit[0])
}
func isBiggerSanZhang(candidate, haveSubmit []string) bool {
	return getIndex(candidate[0]) < getIndex(haveSubmit[0])
}

// isBiggerSanDaiYi 检查candidate是否比haveSubmit大
func isBiggerSanDaiYi(candidate, haveSubmit []string) bool {
	// 拿到candidate的三张牌
	a := ""
	if candidate[0] == candidate[1] && candidate[1] == candidate[2] {
		a = candidate[0]
	} else {
		a = candidate[1]
	}
	b := ""
	if haveSubmit[0] == haveSubmit[1] && haveSubmit[1] == haveSubmit[2] {
		b = haveSubmit[0]
	} else {
		b = haveSubmit[1]
	}
	return getIndex(a) < getIndex(b)
}

// isBiggerDanZhi 检查candidate是否比haveSubmit大
func isBiggerDanZhi(candidate, haveSubmit []string) bool {
	return getIndex(candidate[0]) < getIndex(haveSubmit[0])
}

// 验证扑克牌是否合法
func isValid(candidate []string) bool {
	for i := 0; i < len(validators); i++ {
		if validators[i](candidate) {
			return true
		}
	}
	return false
}

// 检查用哪个验证器可以验证
func whichValidator(candidate []string) int {
	for i := 0; i < len(validators); i++ {
		if validators[i](candidate) {
			return i
		}
	}
	return -1
}

// 检查两幅出牌是否是同一种类型
func isSameType(first, second []string) bool {
	if len(first) == 0 || len(second) == 0 {
		return true
	}
	// 先检查是否是同一种类型
	firstType := whichValidator(first)
	secondType := whichValidator(second)
	if firstType == -1 || secondType == -1 {
		return false
	}
	return firstType == secondType
}

// 炸弹
func isZhaDan(candidate []string) bool {
	if len(candidate) != 4 {
		return false
	}
	// 不能有D,X
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == "D" || candidate[i] == "X" {
			return false
		}
	}
	// 检查三张牌是否相同
	if candidate[0] == candidate[1] && candidate[1] == candidate[2] && candidate[2] == candidate[3] {
		return true
	}
	return false
}

// 单只
func isValidDanZhi(candidate []string) bool {
	return len(candidate) == 1
}

// 对子
func isDuiZi(candidate []string) bool {
	return len(candidate) == 2 && candidate[0] == candidate[1]
}

// 三带一
func isSanDaiYi(candidate []string) bool {
	if len(candidate) != 4 {
		return false
	}
	// 不能是4张相同的牌
	if candidate[0] == candidate[1] && candidate[1] == candidate[2] && candidate[2] == candidate[3] {
		return false
	}
	// 检查三张牌是否相同
	if candidate[0] == candidate[1] && candidate[1] == candidate[2] {
		return true
	}
	if candidate[1] == candidate[2] && candidate[2] == candidate[3] {
		return true
	}
	return false
}

// 三张
func isSanZhang(candidate []string) bool {
	if len(candidate) != 3 {
		return false
	}

	// 不能有D,X
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == "D" || candidate[i] == "X" {
			return false
		}
	}
	return candidate[0] == candidate[1] && candidate[1] == candidate[2]
}

// 连对
func isLianDui(candidate []string) bool {
	if len(candidate) < 6 {
		return false
	}
	// 偶数
	if len(candidate)%2 != 0 {
		return false

	}
	// 不能有2和D,X
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == "2" || candidate[i] == "D" || candidate[i] == "X" {
			return false
		}
	}
	// 检查每个对子是否相同
	for i := 0; i < len(candidate); i += 2 {
		if candidate[i] != candidate[i+1] {
			return false
		}
	}
	// 将每个对子的第一个牌放入一个数组中
	first := make([]string, 0)
	for i := 0; i < len(candidate); i += 2 {
		first = append(first, candidate[i])
	}

	// 查看所有牌是否连续,默认已经从大到小排序
	for i := 0; i < len(first)-1; i++ {
		// 看这个牌和下一个牌是否连续
		preIndex := getIndex(first[i])
		nextIndex := getIndex(first[i+1])
		if preIndex-nextIndex != -1 {
			return false
		}

	}
	return true

}

// 顺子
func isShunzi(candidate []string) bool {
	if len(candidate) < 5 {
		return false
	}
	// 不能有2
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == "2" {
			return false
		}
	}
	// 不能有相同的牌
	for i := 0; i < len(candidate); i++ {
		for j := i + 1; j < len(candidate); j++ {
			if candidate[i] == candidate[j] {
				return false
			}
		}
	}
	// 不能有D,X
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == "D" || candidate[i] == "X" {
			return false
		}
	}
	// 查看所有牌是否连续,默认已经从大到小排序
	for i := 0; i < len(candidate)-1; i++ {
		// 看这个牌和下一个牌是否连续
		preIndex := getIndex(candidate[i])
		nextIndex := getIndex(candidate[i+1])
		if preIndex-nextIndex != -1 {
			return false
		}
	}
	return true
}
func getIndex(x string) int {
	for i := 0; i < len(puke); i++ {
		if puke[i] == x {
			return i
		}
	}
	return -1
}
