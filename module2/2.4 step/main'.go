package main

import "fmt"

// On, Ammo и Power
type testStruct struct {
	On          bool
	Ammo, Power int
}

func (ts *testStruct) Shoot() bool {
	if ts.On == false {
		return false
	} else {
		if ts.Ammo > 0 {
			ts.Ammo--
			return true
		} else {
			return false
		}
	}
}
func (ts *testStruct) RideBike() bool {
	if ts.On == false {
		return false
	} else {
		if ts.Power > 0 {
			ts.Power--
			return true
		} else {
			return false
		}
	}
}
func main() {
	testStruct := testStruct{true, 2, 1}
	fmt.Println(testStruct)
}
