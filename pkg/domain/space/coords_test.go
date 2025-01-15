package space

import (
	"testing"
)

func TestWorldToTileDim(t *testing.T) {
	type args struct {
		w WorldDim
	}
	type want struct {
		result TileDim
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0.0",
			args: args{
				w: 0.0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "world_per_tile-0.1",
			args: args{
				w: WorldPerTile - 0.1,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "world_per_tile",
			args: args{
				w: WorldPerTile,
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WorldToTileDim(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestTileToWorldDim(t *testing.T) {
	type args struct {
		t TileDim
	}
	type want struct {
		result WorldDim
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				t: 0,
			},
			want: want{
				result: 0.0,
			},
		},
		{
			name: "1",
			args: args{
				t: 1,
			},
			want: want{
				result: WorldPerTile,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TileToWorldDim(tt.args.t)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestTileToWorldCtrDim(t *testing.T) {
	type args struct {
		t TileDim
	}
	type want struct {
		result WorldDim
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				t: 0,
			},
			want: want{
				result: WorldPerTile * 0.5,
			},
		},
		{
			name: "1",
			args: args{
				t: 1,
			},
			want: want{
				result: WorldPerTile * 1.5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TileToWorldCtrDim(tt.args.t)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
