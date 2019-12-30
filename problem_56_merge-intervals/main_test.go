package problem_56_merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type test struct {
		name      string
		intervals [][]int
		expect    [][]int
	}

	tests := make([]test, 0)
	tests = append(tests, test{
		name:      "normal",
		intervals: [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
		expect:    [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
	})
	tests = append(tests, test{
		name:      "empty intervals",
		intervals: [][]int{},
		expect:    [][]int{},
	})
	tests = append(tests, test{
		name:      "single interval",
		intervals: [][]int{{1, 6}},
		expect:    [][]int{{1, 6}},
	})
	tests = append(tests, test{
		name:      "two intervals",
		intervals: [][]int{{1, 6}, {2, 8}},
		expect:    [][]int{{1, 8}},
	})
	tests = append(tests, test{
		name:      "two identical intervals",
		intervals: [][]int{{1, 6}, {1, 6}},
		expect:    [][]int{{1, 6}},
	})
	tests = append(tests, test{
		name:      "a interval within another",
		intervals: [][]int{{1, 4}, {2, 3}},
		expect:    [][]int{{1, 4}},
	})
	tests = append(tests, test{
		name:      "one of them is zero length",
		intervals: [][]int{{1, 4}, {5, 5}},
		expect:    [][]int{{1, 5}},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Merge(test.intervals)
			if !reflect.DeepEqual(got, test.expect) {
				t.Fatalf("got:%#v, expect:%#v", got, test.expect)
			}
		})
	}
}
