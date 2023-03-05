package main

import (
	"fmt"
	"sort"
)

type pair struct {
	Name  string
	Count int
}

func MostAppearItem(items []string) []pair {
	itemMap := make(map[string]int)
	for _, item := range items {
		itemMap[item]++
	}

	var itemList []pair
	for name, count := range itemMap {
		itemList = append(itemList, pair{name, count})
	}
	sort.Slice(itemList, func(i, j int) bool {
		return itemList[i].Count < itemList[j].Count
	})
	for _, item := range itemList {
		fmt.Println(item.Name, "=>", item.Count)
	}
	return itemList
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))

}
