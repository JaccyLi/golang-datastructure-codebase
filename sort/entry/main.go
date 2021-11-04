package main

import (
	"fmt"
	"suosuoli-golangds/sort"
	"suosuoli-golangds/utils"
)

// =================algorithm main==================== //
func main() {

	arr := utils.GenSlice(10, 1000)
	fmt.Printf("un: %#v\nsorted:%#v\n", arr, sort.SelectSortAssend(arr))
	fmt.Printf("un: %#v\nsorted:%#v\n", arr, sort.SelectSortDescend(arr))

}
