package space

import (
	"testing"
)

func TestDirection_TurnCW(t *testing.T) {
	type args struct {
		d Direction
	}
	type want struct {
		result Direction
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "north",
			args: args{
				d: North,
			},
			want: want{
				result: East,
			},
		},
		{
			name: "east",
			args: args{
				d: East,
			},
			want: want{
				result: South,
			},
		},
		{
			name: "south",
			args: args{
				d: South,
			},
			want: want{
				result: West,
			},
		},
		{
			name: "west",
			args: args{
				d: West,
			},
			want: want{
				result: North,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.d.TurnCW()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestDirection_TurnCCW(t *testing.T) {
	type args struct {
		d Direction
	}
	type want struct {
		result Direction
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "north",
			args: args{
				d: North,
			},
			want: want{
				result: West,
			},
		},
		{
			name: "east",
			args: args{
				d: East,
			},
			want: want{
				result: North,
			},
		},
		{
			name: "south",
			args: args{
				d: South,
			},
			want: want{
				result: East,
			},
		},
		{
			name: "west",
			args: args{
				d: West,
			},
			want: want{
				result: South,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.d.TurnCCW()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestDirection_Invert(t *testing.T) {
	type args struct {
		d Direction
	}
	type want struct {
		result Direction
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "north",
			args: args{
				d: North,
			},
			want: want{
				result: South,
			},
		},
		{
			name: "east",
			args: args{
				d: East,
			},
			want: want{
				result: West,
			},
		},
		{
			name: "south",
			args: args{
				d: South,
			},
			want: want{
				result: North,
			},
		},
		{
			name: "west",
			args: args{
				d: West,
			},
			want: want{
				result: East,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.d.Invert()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestAddDirection(t *testing.T) {
	type args struct {
		t TileVec
		d Direction
	}
	type want struct {
		result TileVec
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(10,20);north",
			args: args{
				t: TileVec{10, 20},
				d: North,
			},
			want: want{
				result: TileVec{10, 19},
			},
		},
		{
			name: "(10,20);east",
			args: args{
				t: TileVec{10, 20},
				d: East,
			},
			want: want{
				result: TileVec{11, 20},
			},
		},
		{
			name: "(10,20);south",
			args: args{
				t: TileVec{10, 20},
				d: South,
			},
			want: want{
				result: TileVec{10, 21},
			},
		},
		{
			name: "(10,20);west",
			args: args{
				t: TileVec{10, 20},
				d: West,
			},
			want: want{
				result: TileVec{9, 20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddDirection(tt.args.t, tt.args.d)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestGetDirection(t *testing.T) {
	type args struct {
		from TileVec
		to   TileVec
	}
	type want struct {
		result Direction
		ok     bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,1);(1,0)",
			args: args{
				from: TileVec{1, 1},
				to:   TileVec{1, 0},
			},
			want: want{
				result: North,
				ok:     true,
			},
		},
		{
			name: "(1,1);(2,1)",
			args: args{
				from: TileVec{1, 1},
				to:   TileVec{2, 1},
			},
			want: want{
				result: East,
				ok:     true,
			},
		},
		{
			name: "(1,1);(1,2)",
			args: args{
				from: TileVec{1, 1},
				to:   TileVec{1, 2},
			},
			want: want{
				result: South,
				ok:     true,
			},
		},
		{
			name: "(1,1);(0,1)",
			args: args{
				from: TileVec{1, 1},
				to:   TileVec{0, 1},
			},
			want: want{
				result: West,
				ok:     true,
			},
		},
		{
			name: "(1,1);(2,2)",
			args: args{
				from: TileVec{1, 1},
				to:   TileVec{2, 2},
			},
			want: want{
				result: 0,
				ok:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := GetDirection(tt.args.from, tt.args.to)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
			if ok != tt.want.ok {
				t.Fatalf("ok: %t expected, got %t", tt.want.ok, ok)
			}
		})
	}
}
