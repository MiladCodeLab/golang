package main

import (
	"fmt"
	"net/netip"
	"sync"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := sync.Map{}
	user.Store(1, User{
		Name: "Milad",
		Age:  28,
	})

	user.Store(2, User{
		Name: "Reza",
		Age:  45,
	})

	user.Store(3, User{
		Name: "Omid",
		Age:  30,
	})

	//user.Clear() // Delete all

	// if the key doesn't exist it will create it but it the key exist it does nothing.
	//val, ok := user.LoadOrStore(2, User{
	//	Name: "Mohsen",
	//	Age:  21,
	//})
	//if ok {
	//	fmt.Println(val)
	//}

	// change the key with new value
	//prev, ok := user.Swap(1, User{
	//	Name: "Mohsen",
	//	Age:  21,
	//})
	//fmt.Printf("prev: %v, ok: %v\n", prev, ok)

	// delete the key
	//user.Delete(2)

	// if the old value is equal to the stored one. it will do the swap and ok is ture. other than that it will be untouched
	//ok := user.CompareAndSwap(1, User{
	//	Name: "Milad",
	//	Age:  27,
	//}, User{
	//	Name: "Mohsen",
	//	Age:  21,
	//})
	//fmt.Printf("ok: %v\n", ok)

	// it will delete if the value is equal to the stored one and ok will be true
	//ok := user.CompareAndDelete(1, User{
	//	Name: "Milad",
	//	Age:  28,
	//})
	//fmt.Printf("ok: %v\n", ok)

	// it will remove if the key exist int the map and loaded will be true.
	val, loaded := user.LoadAndDelete(2)
	if loaded {
		fmt.Println(val)
	}

	ip := netip.Addr{}

	user.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
