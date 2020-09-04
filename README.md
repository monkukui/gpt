# gpt（go-procon-tools）

`gpt` combine multiple files into oneA

## Install
```
TODO かく
```

## Usage
```
$ go vet -vettool=`which gpt` -main <main_file_path> -lib <library_dir_path> -gen <gen_file_path>
```

## Example

```Go
// original code
package main

import (
	"bufio"
	"fmt"
	"os"

	monkukui "a/lib"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	uf := monkukui.NewUnionFind(n)

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
```

```Go
// library code
package lib

type UnionFind struct {
	par []int
}

/* コンストラクタ */
func NewUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

/* xの所属するグループを返す */
func (u UnionFind) Find(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.Find(u.par[x])
	return u.par[x]
}

/* xの所属するグループ と yの所属するグループ を合体する */
func (u UnionFind) Union(x, y int) {
	xr := u.Find(x)
	yr := u.Find(y)
	if xr == yr {
		return
	}
	if u.Size(yr) < u.Size(xr) {
		yr, xr = swap(yr, xr)
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

/* xとyが同じグループに所属するかどうかを返す */
func (u UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

/* xの所属するグループの木の大きさを返す */
func (u UnionFind) Size(x int) int {
	return -u.par[u.Find(x)]
}

// not exported but used
func swap(a int, b int) (int, int) {
	return b, a
}

// exported but not used
func UnUsedFunction() string {
	return "I am unused exported function"
}

// not exported and not used
func unUsedFunction() string {
	return "I am unused unexported function"
}

// token.VAR
var hoge = 10

// token.CONST
const huga = 100

var (
	v1 = 1
	v2 = 2
	v3 = 3
)
```
```Go
// generated code
package main

import (
	"bufio"
	"fmt"
	"os"
)

func generated_monkukui_swap(a int, b int) (int, int) {
	return b, a
}
func (u generated_monkukui_UnionFind) Size(x int) int {
	return -u.par[u.Find(x)]
}
func (u generated_monkukui_UnionFind) Union(x, y int) {
	xr := u.Find(x)
	yr := u.Find(y)
	if xr == yr {
		return
	}
	if u.Size(yr) < u.Size(xr) {
		yr, xr = generated_monkukui_swap(yr, xr)
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}
func (u generated_monkukui_UnionFind) Find(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.Find(u.par[x])
	return u.par[x]
}
func generated_monkukui_NewUnionFind(N int) *generated_monkukui_UnionFind {
	u := new(generated_monkukui_UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

type generated_monkukui_UnionFind struct{ par []int }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscan(r, &n, &m)
	uf := generated_monkukui_NewUnionFind(n)
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
/* this code was generated by the code below
package main

import (
	"bufio"
	"fmt"
	"os"

	monkukui "a/lib"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)
	uf := monkukui.NewUnionFind(n)

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

*/
```
