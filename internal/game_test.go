package internal

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

func TestGameRandTile(t *testing.T) {
	type field struct {
		board         [][]int
		randomPosFunc randomPositoner
		randomFunc    randomGenerator
	}
	tests := []struct {
		name  string
		input field
		want  [][]int
	}{
		{
			name: "Case1",
			input: field{
				board: [][]int{
					{0, 0, 0, 0},
					{0, 2, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 4},
				},
				randomPosFunc: func(TotalSize int) int {
					return 1
				},
				randomFunc: func() float64 {
					return 0.2
				},
			},
			want: [][]int{
				{0, 2, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 4},
			},
		},
		{
			name: "Case2",
			input: field{
				board: [][]int{
					{2, 4, 8, 16},
					{32, 64, 128, 256},
					{512, 1024, 2048, 0},
					{2, 4, 8, 16},
				},
				randomPosFunc: func(TotalSize int) int {
					return 0
				},
				randomFunc: func() float64 {
					return 0.09
				},
			},
			want: [][]int{
				{2, 4, 8, 16},
				{32, 64, 128, 256},
				{512, 1024, 2048, 4},
				{2, 4, 8, 16},
			},
		},
		{
			name: "Case3",
			input: field{
				board: [][]int{
					{2, 4, 2, 4},
					{4, 2, 4, 2},
					{2, 4, 2, 4},
					{4, 2, 4, 0},
				},
				randomPosFunc: func(TotalSize int) int {
					return 0
				},
				randomFunc: func() float64 {
					return 0.5
				},
			},
			want: [][]int{
				{2, 4, 2, 4},
				{4, 2, 4, 2},
				{2, 4, 2, 4},
				{4, 2, 4, 2},
			},
		},
		{
			name: "Case4",
			input: field{
				board: [][]int{
					{8, 16, 32, 2},
					{4, 2, 8, 4},
					{16, 32, 4, 8},
					{2, 4, 16, 0},
				},
				randomPosFunc: func(TotalSize int) int {
					return 0
				},
				randomFunc: func() float64 {
					return 0.5
				},
			},
			want: [][]int{
				{8, 16, 32, 2},
				{4, 2, 8, 4},
				{16, 32, 4, 8},
				{2, 4, 16, 2},
			},
		},
		{
			name: "Case5 - full of tile no change",
			input: field{
				board: [][]int{
					{2, 4, 8, 16},
					{32, 64, 128, 256},
					{512, 1024, 2048, 2},
					{4, 8, 16, 32},
				},
				randomPosFunc: func(TotalSize int) int {
					return 0
				},
				randomFunc: func() float64 {
					return 0.5
				},
			},
			want: [][]int{
				{2, 4, 8, 16},
				{32, 64, 128, 256},
				{512, 1024, 2048, 2},
				{4, 8, 16, 32},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, tt.input.randomPosFunc, tt.input.randomFunc)
			// 模擬隨機產生 tile
			game.addRandomTile()
			assert.Equal(t, tt.want, game.board)
		})
	}

}

func TestGameMoveLeft(t *testing.T) {
	type field struct {
		board [][]int
	}
	tests := []struct {
		name  string
		input field
		want  [][]int
	}{
		{
			name: "case1: 單行多次合併",
			input: field{
				board: [][]int{
					{4, 4, 4, 4},
					{2, 2, 0, 0},
					{2, 0, 2, 0},
					{8, 0, 0, 8},
				},
			},
			want: [][]int{
				{8, 8, 0, 0},
				{4, 0, 0, 0},
				{4, 0, 0, 0},
				{16, 0, 0, 0},
			},
		},
		{
			name: "case2: 新生成的數字不參與當回合合併",
			input: field{
				board: [][]int{
					{2, 2, 4, 8},
					{0, 0, 0, 0},
					{4, 4, 8, 8},
					{2, 2, 2, 2},
				},
			},
			want: [][]int{
				{4, 4, 8, 0},
				{0, 0, 0, 0},
				{8, 16, 0, 0},
				{4, 4, 0, 0},
			},
		},
		{
			name: "case3: 無合併，只有移動",
			input: field{
				board: [][]int{
					{2, 4, 8, 16},
					{0, 2, 0, 4},
					{8, 0, 4, 0},
					{2, 4, 2, 4},
				},
			},
			want: [][]int{
				{2, 4, 8, 16},
				{2, 4, 0, 0},
				{8, 4, 0, 0},
				{2, 4, 2, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			// 模擬左移
			game.moveLeft()
			assert.Equal(t, tt.want, game.board)
		})
	}
}
