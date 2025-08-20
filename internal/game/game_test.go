package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameInit(t *testing.T) {
	type field struct {
		board [][]int
	}
	tests := []struct {
		name  string
		input field
		want  [][]int
	}{
		{
			name: "Empty Input",
			input: field{
				board: nil,
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Case1",
			input: field{
				board: [][]int{
					{2, 4, 0, 0},
					{8, 0, 0, 0},
					{16, 0, 4, 0},
					{0, 0, 0, 0},
				},
			},
			want: [][]int{
				{2, 4, 0, 0},
				{8, 0, 0, 0},
				{16, 0, 4, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Case2",
			input: field{
				board: [][]int{
					{128, 64, 32, 16},
					{8, 0, 0, 2},
					{4, 0, 0, 4},
					{2, 4, 8, 16},
				},
			},
			want: [][]int{
				{128, 64, 32, 16},
				{8, 0, 0, 2},
				{4, 0, 0, 4},
				{2, 4, 8, 16},
			},
		},
		{
			name: "Case3",
			input: field{
				board: [][]int{
					{0, 0, 0, 0},
					{0, 2048, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 2048, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "Case4",
			input: field{
				board: [][]int{
					{16, 8, 4, 2},
					{32, 16, 8, 4},
					{64, 32, 16, 8},
					{128, 64, 32, 16},
				},
			},
			want: [][]int{
				{16, 8, 4, 2},
				{32, 16, 8, 4},
				{64, 32, 16, 8},
				{128, 64, 32, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			assert.Equal(t, tt.want, game.board)
		})
	}
}
