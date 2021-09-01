package tekton_k8s

import (
	"fmt"
	"sort"
)

func Countstr(str string)  (result map[int]int) {

	map1 := make(map[int]int)
	for i := 0; i < len(str); i++ {
		if map1[int(str[i])] == 0 {
			map1[int(str[i])] = 1
		} else if map1[int(str[i])] > 0 {
			num := map1[int(str[i])]
			map1[int(str[i])] = num + 1
		}
	}
	//for key,count:= range map1 {
	//	fmt.Printf("%c : %d   \n",key,count)
	//}

	keys := make([]int, 0, len(map1))
	for key := range map1 {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	result = make(map[int]int)
	for _ ,key := range keys {
		fmt.Printf("%c : %d \n",key,map1[key])
		result[key] = map1[key]


	}
	return
}
