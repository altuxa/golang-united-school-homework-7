package coverage

import (
	"errors"
	"os"
	"strconv"
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

var ErrLenght = errors.New("Rows need to be the same length")
var _, ErrAtoi = strconv.Atoi("a")

func TestNew(t *testing.T) {
	tData := map[string]struct{
		str string
		Expected *Matrix
		Err error
	}{
		"succes": {str: "12 2 3 \n 5 6 9 \n10 23 4", Expected: &Matrix{rows:3,cols:3,data: []int{12,2,3,5,6,9,10,23,4}}, Err: nil},
		"errorLenght": {str: "1 2 \n 3", Expected: nil, Err: ErrLenght},
		"errorAtoi": {str: "a\n a",Expected: nil, Err: ErrAtoi},
	}
	for tName, tCase := range tData {
		v := tCase
		t.Run(tName, func(t *testing.T) {
			got,err := New(v.str)
			check := false
			if err != nil && err.Error() != v.Err.Error(){
				t.Errorf("[%s] expected error: %v, got error: %v", tName, v.Err, err)
			}
			if v.Expected != nil {
				for ind, i := range got.data {
					if v.Expected.data[ind] != i {
						check = true
					}
				}
				if got.cols != v.Expected.cols || got.rows != v.Expected.rows || check {
					t.Errorf("[%s] expected: %v, got: %v",tName, v.Expected, got)
				}
			}
		})
	}
}
