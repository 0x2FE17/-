package main

import (
	"fmt"
	"math"
	"strconv"
)

func dfs(i int) {
	println(i)
	dfs(i + 1)
}

//func main() {
//	println("Hello World")
//	dfs(1)
//}

func removeStones(stones [][]int) int {
	uf := NewUnionFind()
	for _, stone := range stones {
		uf.Union("r"+strconv.Itoa(stone[0]), "c"+strconv.Itoa(stone[1]))
	}
	return len(stones) - uf.count
}

type UnionFind struct {
	parent map[string]string
	count  int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[string]string),
		count:  0,
	}
}
func (uf *UnionFind) Union(x, y string) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.parent[rootX] = rootY
	uf.count--
}
func (uf *UnionFind) Find(x string) string {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.count++
		return x
	}
	if uf.parent[x] == x {
		return x
	}
	uf.parent[x] = uf.Find(uf.parent[x])
	return uf.parent[x]
}
func kuohaoqujianpipei() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return
	}
	n := len(input)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}
		if memo[l][r] != -1 {
			return memo[l][r]
		}
		if input[l] == ']' || input[l] == ')' {
			ans := 1 + dfs(l+1, r)
			memo[l][r] = ans
			return ans
		}
		ans := math.MaxInt32
		if input[l] == '[' {
			for i := l + 1; i <= r; i++ {
				if input[i] == ']' {
					p := dfs(l+1, i-1) + dfs(i+1, r)
					if p < ans {
						ans = p
					}
				}
			}
			notTake := dfs(l+1, r) + 1
			if notTake < ans {
				ans = notTake
			}
		} else {
			for i := l + 1; i <= r; i++ {
				if input[i] == ')' {
					p := dfs(l+1, i-1) + dfs(i+1, r)
					if p < ans {
						ans = p
					}
				}
			}
			notTake := dfs(l+1, r) + 1
			if notTake < ans {
				ans = notTake
			}
		}
		memo[l][r] = ans
		return ans
	}
	ans := dfs(0, n-1)
	fmt.Println(ans)
}
