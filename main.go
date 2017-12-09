package main

import (
	"fmt"
)

type teststr struct {
	member []int
}

func deepCopy (hoge []teststr) []teststr {
	fuga := make([]teststr,len(hoge))
	
	for i := 0; i < len(hoge) ; i++ {
		fuga[i].member = make([]int,len(hoge[i].member))
		copy(fuga[i].member,hoge[i].member)
	}
	return fuga
}

func main() {
	hoge := make([]teststr, 5)
	for i := 0; i < 5; i++ {
		hoge[i].member = []int{1 * i, 2 * i, 3 * i, 4 * i, 5 * i}
	}
	fmt.Println(hoge)

	fuga := deepCopy(hoge)
	
	fmt.Println(fuga)

	fuga[0].member[4] = 123456
	hoge[1].member[4] = 654321

	fmt.Println(hoge, fuga)
}

