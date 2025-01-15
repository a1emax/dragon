package space

import (
	"testing"
)

func TestClipViewDim(t *testing.T) {
	type args struct {
		area     TileDim
		focus    TileDim
		viewport WorldDim
	}
	type want struct {
		offset WorldDim
		from   TileDim
		to     TileDim
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		// The viewport is larger than the area.
		{
			name: "11;5;x12",
			args: args{
				area:     11,
				focus:    5,
				viewport: 12 * WorldPerTile,
			},
			want: want{
				offset: 0.5 * WorldPerTile,
				from:   0,
				to:     10,
			},
		},
		// The viewport is 5 tiles of 11.
		{
			name: "11;1;x5",
			args: args{
				area:     11,
				focus:    1,
				viewport: 5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   0,
				to:     4,
			},
		},
		{
			name: "11;4;x5",
			args: args{
				area:     11,
				focus:    4,
				viewport: 5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   2,
				to:     6,
			},
		},
		{
			name: "11;5;x5",
			args: args{
				area:     11,
				focus:    5,
				viewport: 5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   3,
				to:     7,
			},
		},
		{
			name: "11;6;x5",
			args: args{
				area:     11,
				focus:    6,
				viewport: 5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   4,
				to:     8,
			},
		},
		{
			name: "11;9;x5",
			args: args{
				area:     11,
				focus:    9,
				viewport: 5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   6,
				to:     10,
			},
		},
		// The viewport is 5.5 tiles of 11.
		{
			name: "11;1;x5.5",
			args: args{
				area:     11,
				focus:    1,
				viewport: 5.5 * WorldPerTile,
			},
			want: want{
				offset: 0.0,
				from:   0,
				to:     5,
			},
		},
		{
			name: "11;4;x5.5",
			args: args{
				area:     11,
				focus:    4,
				viewport: 5.5 * WorldPerTile,
			},
			want: want{
				offset: -0.75 * WorldPerTile,
				from:   1,
				to:     7,
			},
		},
		{
			name: "11;5;x5.5",
			args: args{
				area:     11,
				focus:    5,
				viewport: 5.5 * WorldPerTile,
			},
			want: want{
				offset: -0.75 * WorldPerTile,
				from:   2,
				to:     8,
			},
		},
		{
			name: "11;6;x5.5",
			args: args{
				area:     11,
				focus:    6,
				viewport: 5.5 * WorldPerTile,
			},
			want: want{
				offset: -0.75 * WorldPerTile,
				from:   3,
				to:     9,
			},
		},
		{
			name: "11;9;x5.5",
			args: args{
				area:     11,
				focus:    9,
				viewport: 5.5 * WorldPerTile,
			},
			want: want{
				offset: -0.5 * WorldPerTile,
				from:   5,
				to:     10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offset, from, to := ClipViewDim(tt.args.area, tt.args.focus, tt.args.viewport)
			if offset != tt.want.offset {
				t.Fatalf("offset: %g expected, got %g", tt.want.offset, offset)
			}
			if from != tt.want.from {
				t.Fatalf("from: %d expected, got %d", tt.want.from, from)
			}
			if to != tt.want.to {
				t.Fatalf("to: %d expected, got %d", tt.want.to, to)
			}
		})
	}
}
