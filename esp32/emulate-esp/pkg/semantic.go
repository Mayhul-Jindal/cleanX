package pkg

import "fmt"

func Semantic(){
	items := [][2]byte{{1,2},{3,4},{5,6}}
	sliceItems := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:])
		sliceItems = append(sliceItems, i)
	}

	fmt.Printf("%v\n", items);
	fmt.Printf("%v", sliceItems);
}