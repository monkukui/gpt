package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	generated_lib_v1 = 1
	generated_lib_v2 = 2
	generated_lib_v3 = 3
)

const generated_lib_huga = 100

var generated_lib_hoge = 10

func generated_lib_unUsedFunction() string {
	return "I am unused unexported function"
}
func generated_lib_UnUsedFunction() string {
	return "I am unused exported function"
}
func generated_lib_swap(a int, b int) (int, int) {
	return b, a
}
func (u generated_lib_UnionFind) Size(x int) int {
	return -u.par[u.Find(x)]
}
func (u generated_lib_UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
func (u generated_lib_UnionFind) Union(x, y int) {
	xr := u.Find(x)
	yr := u.Find(y)
	if xr == yr {
		return
	}
	if u.Size(yr) < u.Size(xr) {
		yr, xr = generated_lib_swap(yr, xr)
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}
func (u generated_lib_UnionFind) Find(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.Find(u.par[x])
	return u.par[x]
}
func generated_lib_NewUnionFind(N int) *generated_lib_UnionFind {
	u := new(generated_lib_UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

type generated_lib_UnionFind struct{ par []int }

func generated_lib_ModPow() int64 {
	return int64(math.Max(1, 3))
}
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscan(r, &n, &m)
	uf := generated_lib_NewUnionFind(n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		a--
		b--
		uf.Union(a, b)
	}
	ans := 0
	for i := 0; i < n; i++ {
		if ans < uf.Size(i) {
			ans = uf.Size(i)
		}
	}
	fmt.Fprintln(w, ans)
}
