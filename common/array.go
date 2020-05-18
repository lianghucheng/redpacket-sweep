package common

import (
	"math/rand"
)

func Shuffle(a []int) []int {
	b := append([]int{}, a...)
	rand.Shuffle(len(a), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	return b
}

func Index(a []int, sep int) int {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == sep {
			return i
		}
	}
	return -1
}

// 判断 sep 是否在数组中
func InArray(a []int, sep int) bool {
	for _, v := range a {
		if sep == v {
			return true
		}
	}
	return false
}

// 从数组中移除最开始出现的 sep
func RemoveOnce(a []int, sep int) []int {
	i := Index(a, sep)
	switch i {
	case -1:
		return a
	case 0:
		return a[1:]
	case len(a) - 1:
		return a[:i]
	default:
		b := append([]int{}, a[:i]...)
		b = append(b, a[i+1:]...)
		return b
	}
}

// 从数组中移除所有的 sep
func RemoveAll(a []int, sep int) []int {
	b := make([]int, 0)
	for _, v := range a {
		if v != sep {
			b = append(b, v)
		}
	}
	return b
}

func Remove(a []int, sub []int) []int {
	b := append([]int{}, a...)
	for _, v := range sub {
		b = RemoveOnce(b, v)
	}
	return b
}

func ReplaceAll(a []int, old, new int) []int {
	if old == new {
		return a
	}
	if InArray(a, old) {
		b := make([]int, 0)
		for _, v := range a {
			if old == v {
				b = append(b, new)
			} else {
				b = append(b, v)
			}
		}
		return b
	}
	return a
}

func Deduplicate(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	}
	m := make(map[int]bool)

	b := []int{a[0]}
	m[a[0]] = true
	for i := 1; i < n; i++ {
		if !m[a[i]] {
			b = append(b, a[i])
			m[a[i]] = true
		}
	}
	return b
}

// 比较两个数组的元素是否相等
func Equal(x, y []int) bool {
	if len(x) == len(y) {
		return Contain(x, y)
	}
	return false
}

func Contain(x, y []int) bool {
	if len(x) < len(y) {
		return false
	}
	temp := Deduplicate(y)
	for _, v := range temp {
		if Count(x, v) < Count(y, v) {
			return false
		}
	}
	return true
}

func Count(a []int, sep int) int {
	count := 0
	for _, v := range a {
		if sep == v {
			count++
		}
	}
	return count
}
