package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func NewPeopleMock() People {
	return People{
		Person{
			firstName: "Aboba",
			lastName: "Abobov",
			birthDay: time.Date(1998,01,29,10,10,10,0,time.UTC),
		},
		Person{
			firstName: "Bill",
			lastName: "Bank",
			birthDay:  time.Date(1980,10,15,0,0,0,0,time.UTC),
		},
		Person{
			firstName: "Bob",
			lastName: "Gin",
			birthDay: time.Date(1980,10,15,0,0,0,0,time.UTC),
		},
		Person{
			firstName: "Alex",
			lastName: "Grey",
			birthDay: time.Date(1980,10,15,0,0,0,0,time.UTC),
		},
	}
}

func TestLen(t *testing.T) {
	tData := map[string]struct{
		A People
		Expected int
	} {
		"succes": {A:NewPeopleMock(),Expected: 4},
	}
	for tName, tCase := range tData {
		v := tCase
		t.Run(tName,func(t *testing.T) {
			got := v.A.Len()
			if got != v.Expected {
				t.Errorf("[%s] expected: %d, got: %d",tName,v.Expected,got)
			}
		})
	}
}

func TestLess(t *testing.T) {
	tData := map[string]struct{
		A int
		B int
		people People
		Expected bool
	}{
		"succes case 1": {A: 1,B: 2,people: NewPeopleMock(),Expected: true},
		"succes case 2": {A: 2,B: 3, people: NewPeopleMock(),Expected: false},
		"succes case 3": {A: 0,B: 3,people: NewPeopleMock(),Expected: true},
	}
	for tName, tCase := range tData {
		v := tCase
		t.Run(tName, func(t *testing.T) {
			got := v.people.Less(v.A,v.B)
			if got != v.Expected {
				t.Errorf("[%s] expected: %t got: %t",tName,v.Expected,got)
			}
		})
	}
}

// func expectedForSwap(p People) People {

// } 

func TestSwap(t *testing.T) {
	tData := map[string]struct{
		indexA int
		indexB int
		people People
		//Expected People
	}{
		"succes": {indexA: 0, indexB: 1,people: NewPeopleMock()},
	}
	for tName, tCase := range tData {
		v := tCase
		t.Run(tName, func(t *testing.T) {
			oldA := v.people[v.indexA]
			oldB := v.people[v.indexB]
			v.people.Swap(v.indexA,v.indexB)
			if v.people[v.indexA] != oldB || v.people[v.indexB] != oldA {
				t.Errorf("[%s] expected: %v got: %v", tName, oldA, v.people[v.indexB])
			} 
		})
	}
}