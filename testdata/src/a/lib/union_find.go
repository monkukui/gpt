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
