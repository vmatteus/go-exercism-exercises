package cards

import (
	"fmt"
	"testing"
)

func TestFavoriteCards(t *testing.T) {
	got := FavoriteCards()
	fmt.Println(got)
}

func TestGetItem(t *testing.T) {
	got := GetItem([]int{1, 2, 4, 1}, 10)
	fmt.Println(got)
}

func TestSetItem(t *testing.T) {
	index := 1
	newCard := 6
	got := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(got)
}

func TestPrependItems(t *testing.T) {
	slice := []int{3, 2, 6, 4, 8}
	got := PrependItems(slice, 5, 1)
	fmt.Println(got)
}

func TestRemoveItem(t *testing.T) {
	got := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(got)
}
