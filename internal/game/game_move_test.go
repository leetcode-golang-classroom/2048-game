package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGameMove(t *testing.T) {

	type field struct {
		board [][]int
	}
	type result struct {
		left  [][]int
		right [][]int
		up    [][]int
		down  [][]int
	}
	tests := []struct {
		name  string
		input field
		want  result
	}{
		{
			name: "case1",
			input: field{
				board: [][]int{
					{2, 2, 0, 4},
					{0, 4, 4, 8},
					{2, 0, 2, 0},
					{8, 8, 8, 8},
				},
			},
			want: result{
				left: [][]int{
					{4, 4, 0, 0},
					{8, 8, 0, 0},
					{4, 0, 0, 0},
					{16, 16, 0, 0},
				},
				right: [][]int{
					{0, 0, 4, 4},
					{0, 0, 8, 8},
					{0, 0, 0, 4},
					{0, 0, 16, 16},
				},
				up: [][]int{
					{4, 2, 4, 4},
					{8, 4, 2, 16},
					{0, 8, 8, 0},
					{0, 0, 0, 0},
				},
				down: [][]int{
					{0, 0, 0, 0},
					{0, 2, 4, 0},
					{4, 4, 2, 4},
					{8, 8, 8, 16},
				},
			},
		},
		{
			name: "case2",
			input: field{
				board: [][]int{
					{4, 4, 8, 16},
					{4, 0, 4, 0},
					{2, 2, 2, 2},
					{0, 0, 0, 0},
				},
			},
			want: result{
				left: [][]int{
					{8, 8, 16, 0},
					{8, 0, 0, 0},
					{4, 4, 0, 0},
					{0, 0, 0, 0},
				},
				right: [][]int{
					{0, 8, 8, 16},
					{0, 0, 0, 8},
					{0, 0, 4, 4},
					{0, 0, 0, 0},
				},
				up: [][]int{
					{8, 4, 8, 16},
					{2, 2, 4, 2},
					{0, 0, 2, 0},
					{0, 0, 0, 0},
				},
				down: [][]int{
					{0, 0, 0, 0},
					{0, 0, 8, 0},
					{8, 4, 4, 16},
					{2, 2, 2, 2},
				},
			},
		},
		{
			name: "case3 - 無法移動",
			input: field{
				board: [][]int{
					{2, 4, 8, 16},
					{4, 8, 16, 2},
					{8, 16, 2, 4},
					{16, 2, 4, 8},
				},
			},
			want: result{
				left: [][]int{
					{2, 4, 8, 16},
					{4, 8, 16, 2},
					{8, 16, 2, 4},
					{16, 2, 4, 8},
				},
				right: [][]int{
					{2, 4, 8, 16},
					{4, 8, 16, 2},
					{8, 16, 2, 4},
					{16, 2, 4, 8},
				},
				up: [][]int{
					{2, 4, 8, 16},
					{4, 8, 16, 2},
					{8, 16, 2, 4},
					{16, 2, 4, 8},
				},
				down: [][]int{
					{2, 4, 8, 16},
					{4, 8, 16, 2},
					{8, 16, 2, 4},
					{16, 2, 4, 8},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s left", tt.name), func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			// 模擬左移
			game.moveLeft()
			assert.Equal(t, tt.want.left, game.board)
		})
		t.Run(fmt.Sprintf("%s right", tt.name), func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			// 模擬右移
			game.moveRight()
			assert.Equal(t, tt.want.right, game.board)
		})
		t.Run(fmt.Sprintf("%s up", tt.name), func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			// 模擬上移
			game.moveUp()
			assert.Equal(t, tt.want.up, game.board)
		})
		t.Run(fmt.Sprintf("%s down", tt.name), func(t *testing.T) {
			game := NewGame()
			game.Init(tt.input.board, nil, nil)
			// 模擬下移
			game.moveDown()
			assert.Equal(t, tt.want.down, game.board)
		})
	}
}
