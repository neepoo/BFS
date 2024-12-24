package main

import "fmt"

type Dot struct {
	x, y, step int
}

const N = 310

var (
	dx = []int{-2, -2, -1, -1, 1, 1, 2, 2}
	dy = []int{1, -1, 2, -2, 2, -2, 1, -1}
)

func BFS(sx, sy, ex, ey, l int) int {
	if sx == ex && sy == ey {
		return 0
	}

	var (
		queue = make([]Dot, 0, 1000)
		vis   [N][N]bool
	)
	queue = append(queue, Dot{sx, sy, 0})
	vis[sx][sy] = true
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		// 遍历八个方向
		for i := range 8 {
			ax := head.x + dx[i]
			ay := head.y + dy[i]
			if ax < l && ay < l && ax >= 0 && ay >= 0 && !vis[ax][ay] {
				// 在棋盘内
				if ax == ex && ay == ey {
					return head.step + 1
				}
				queue = append(queue, Dot{ax, ay, head.step + 1})
				vis[ax][ay] = true
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(BFS(0, 0, 7, 0, 8))
	fmt.Println(BFS(0, 0, 30, 50, 100))
	fmt.Println(BFS(1, 1, 1, 1, 10))
}
