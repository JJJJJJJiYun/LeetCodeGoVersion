package o7

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestReconstructTreeWithPreAndInOrderTraversal(t *testing.T) {
	type args struct {
		pre []int
		in  []int
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			name: "case1",
			args: args{
				pre: []int{1, 2, 4, 7, 3, 5, 6, 8},
				in:  []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReconstructTreeWithPreAndInOrderTraversal(tt.args.pre, tt.args.in)
			utils.HierarchicalTraversalWithBreak(got)
		})
	}
}

func TestReconstructTreeWithInAndPostOrderTraversal(t *testing.T) {
	type args struct {
		in   []int
		post []int
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			name: "case1",
			args: args{
				in:   []int{9, 3, 15, 20, 7},
				post: []int{9, 15, 7, 20, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReconstructTreeWithInAndPostOrderTraversal(tt.args.in, tt.args.post)
			utils.HierarchicalTraversalWithBreak(got)
		})
	}
}
