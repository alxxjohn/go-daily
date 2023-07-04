// package main

// import "testing"

// func TestPrependItems(t *testing.T) {
// 	type args struct {
// 		slice []int
// 		value []int
// 	}
// 	tt := struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		name: "Prepend one item",
// 		args: args{
// 			slice: []int{5, 2, 10, 6, 8, 7, 0, 9},
// 			value: []int{1},
// 		},
// 		want: []int{1, 5, 2, 10, 6, 8, 7, 0, 9},
// 	}

// 	got := PrependItems(tt.args.slice, tt.args.value...)
// 	if !slicesEqual(got, tt.want) {
// 		t.Errorf("PrependItems(slice:%v, value:%v) = %v, want %v",
// 			tt.args.slice, tt.args.value, got, tt.want)
// 	}

// }