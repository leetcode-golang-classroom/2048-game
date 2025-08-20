package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
