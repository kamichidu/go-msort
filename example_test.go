package msort_test

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kamichidu/go-msort"
)

type User struct {
	Order int64

	Email string
}

func Example_struct() {
	users := []User{
		{4, "dev04@example.com"},
		{1, "dev01@example.com"},
		{3, "dev13@example.com"},
		{3, "dev03@example.com"},
		{0, "dev00@example.com"},
		{2, "dev02@example.com"},
		{0, "dev10@example.com"},
	}
	sorter := msort.NewBuilderFrom(users).
		ByFunc(func(i, j int) bool {
			return users[i].Order < users[j].Order
		}).
		ByFunc(func(i, j int) bool {
			return strings.Compare(users[i].Email, users[j].Email) < 0
		}).
		ToSorter()
	sort.Stable(sorter)
	fmt.Println("**ascending**")
	for i := range users {
		fmt.Printf("%d - %s\n", users[i].Order, users[i].Email)
	}
	sort.Stable(sort.Reverse(sorter))
	fmt.Println("**descending**")
	for i := range users {
		fmt.Printf("%d - %s\n", users[i].Order, users[i].Email)
	}
	// Output:
	// **ascending**
	// 0 - dev00@example.com
	// 0 - dev10@example.com
	// 1 - dev01@example.com
	// 2 - dev02@example.com
	// 3 - dev03@example.com
	// 3 - dev13@example.com
	// 4 - dev04@example.com
	// **descending**
	// 4 - dev04@example.com
	// 3 - dev13@example.com
	// 3 - dev03@example.com
	// 2 - dev02@example.com
	// 1 - dev01@example.com
	// 0 - dev10@example.com
	// 0 - dev00@example.com
}
