package main

import (
        "fmt"
        "net/http"
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

func sayHello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello golang!")
}

func main() {
        http.HandleFunc("/", sayHello)
        fmt.Println("start 127.0.0.1:8080...")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                fmt.Printf("http server failed, err:%v\n", err)
                return
        }
}
