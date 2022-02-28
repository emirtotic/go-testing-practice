package main

import "testing"

func TestInts(t *testing.T) {

	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"One to Five", []int{1, 2, 3, 4, 5}, 15},
		{"No numbers", nil, 0},
		{"From -1 to 1", []int{1, -1}, 0},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.sum {
				t.Fatalf("Test %v result: Sum of %v should be %v!", tc.name, tc.numbers, tc.sum)
			}
		})
	}

}

//Run one single test :    go test -run TestInts -v