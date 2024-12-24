package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var n int

func main() {
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for sc.Scan() {
		line := sc.Text()
		n, _ = strconv.Atoi(line)
		l := list.New()
		k := 2
		var i = 1
		for i <= n {
			l.PushBack(i)
			i++
		}
		for l.Len() > 3 {
			var cnt = 1
			for e := l.Front(); e != nil; {
				next := e.Next() // 先获取下一个元素
				if cnt%k == 0 {
					l.Remove(e)
				}
				cnt++
				e = next // 移动到下一个元素
			}

			//k = 5 - k
			k ^= 1 // (2 -> 3, 3->2)
		}
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d ", e.Value.(int))
		}
		fmt.Println()
	}
}
