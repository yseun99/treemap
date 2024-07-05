package example_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/igrmk/treemap/v2"
	"golang.org/x/exp/rand"
)

func TestExample(t *testing.T) {
	tm := treemap.New[int, string]()

	seed := time.Now().UnixNano()
	source := rand.NewSource(uint64(seed))
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		n := r.Intn(10)
		if tm.Contains(n) {
			fmt.Printf("contains: %d\n", n)
			continue
		}
		tm.Set(n, fmt.Sprintf("value_%d", n))
	}
	fmt.Println("-------")
	fmt.Printf("len: %d\n", tm.Len())
	for it := tm.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}

func TestIssue2(t *testing.T) {
	tm := treemap.New[int, string]()
	tm.Set(1, "World")
	tm.Set(0, "Hello")
	tm.Set(2, "Haha")
	tm.Set(3, "Wowo")

	fmt.Printf("len: %d\n", tm.Len())
	for it := tm.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	fmt.Println("-------")
	for it := tm.Reverse(); it.Valid(); it.Next() {
		if it.Value() == "Haha" {
			tm.Del(it.Key())
			fmt.Printf("Delete: %d %v\n", it.Key(), it.Value())
		}
	}
	fmt.Println("-------")
	fmt.Printf("len: %d\n", tm.Len())
	for it := tm.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	fmt.Println("------- greater than 0")
	for it := tm.UpperBound(0); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
	fmt.Println("------- greater than 3")
	for it := tm.UpperBound(4); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}
