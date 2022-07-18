/**
 *@Author:sario
 *Date:2022/7/16
 *@Desc:
 */
package chess

import (
	"testing"
)

func Test_jiangSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqSrc: 56,
			sqDst: 73,
		}, want: false},

		{name: "case2", args: args{
			sqSrc: 17,
			sqDst: 1,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jiangSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("jiangSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqSrc: 57,
			sqDst: 1,
		}, want: false},

		{name: "case2", args: args{
			sqSrc: 16,
			sqDst: 1,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("shiSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xiangSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqSrc: 56,
			sqDst: 73,
		}, want: false},

		{name: "case2", args: args{
			sqSrc: 35,
			sqDst: 1,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xiangSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("xiangSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inBoard(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{sq: 55}, want: true},
		{name: "case2", args: args{sq: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inBoard(tt.args.sq); got != tt.want {
				t.Errorf("inBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inFort(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{sq: 55}, want: true},
		{name: "case2", args: args{sq: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inFort(tt.args.sq); got != tt.want {
				t.Errorf("inFort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorSquare(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{sq: 23}, want: 23},
		{name: "case2", args: args{sq: 2}, want: 12},
		{name: "case3", args: args{sq: 32}, want: 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorSquare(tt.args.sq); got != tt.want {
				t.Errorf("mirrorSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareForward(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sq: 1,
			sd: 1,
		}, want: 17},

		{name: "case2", args: args{
			sq: 18,
			sd: 1,
		}, want: 34},

		{name: "case3", args: args{
			sq: 200,
			sd: 1,
		}, want: 216},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareForward(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("squareForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xiangPin(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sqDst: 18,
			sqSrc: 1,
		}, want: 9},

		{name: "case2", args: args{
			sqDst: 18,
			sqSrc: 10,
		}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xiangPin(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("xiangPin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maPin(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sqDst: 18,
			sqSrc: 1,
		}, want: 1},

		{name: "case2", args: args{
			sqDst: 18,
			sqSrc: 10,
		}, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maPin(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("maPin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noRiver(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sq: 18,
			sd: 1,
		}, want: true},

		{name: "case2", args: args{
			sq: 180,
			sd: 1,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noRiver(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("noRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasRiver(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sq: 18,
			sd: 1,
		}, want: false},

		{name: "case2", args: args{
			sq: 180,
			sd: 1,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRiver(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("hasRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameRiver(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqSrc: 18,
			sqDst: 1,
		}, want: true},

		{name: "case2", args: args{
			sqSrc: 180,
			sqDst: 1,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameRiver(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameX(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqDst: 18,
			sqSrc: 1,
		}, want: false},

		{name: "case2", args: args{
			sqDst: 19,
			sqSrc: 18,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameX(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameY(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{
			sqDst: 0,
			sqSrc: 16,
		}, want: true},

		{name: "case2", args: args{
			sqDst: 19,
			sqSrc: 180,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameY(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorMove(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			mv: 1,
		}, want: 3597},

		{name: "case2", args: args{
			mv: 10,
		}, want: 3588},

		{name: "case3", args: args{
			mv: 18,
		}, want: 3612},

		{name: "case4", args: args{
			mv: 44,
		}, want: 3618},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorMove(tt.args.mv); got != tt.want {
				t.Errorf("mirrorMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getY(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sq: 1,
		}, want: 0},

		{name: "case2", args: args{
			sq: 10,
		}, want: 0},

		{name: "case3", args: args{
			sq: 19,
		}, want: 1},

		{name: "case4", args: args{
			sq: 189,
		}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getY(tt.args.sq); got != tt.want {
				t.Errorf("getY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getX(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sq: 1,
		}, want: 1},

		{name: "case2", args: args{
			sq: 10,
		}, want: 10},

		{name: "case3", args: args{
			sq: 180,
		}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getX(tt.args.sq); got != tt.want {
				t.Errorf("getX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareXY(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			x: 1,
			y: 1,
		}, want: 17},

		{name: "case2", args: args{
			x: 10,
			y: 1,
		}, want: 26},

		{name: "case3", args: args{
			x: 10,
			y: 10,
		}, want: 170},

		{name: "case4", args: args{
			x: 108,
			y: 1,
		}, want: 124},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareXY(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("squareXY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareFlip(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sq: 1,
		}, want: 253},

		{name: "case2", args: args{
			sq: 10,
		}, want: 244},

		{name: "case3", args: args{
			sq: 100,
		}, want: 154},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareFlip(tt.args.sq); got != tt.want {
				t.Errorf("squareFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xFlip(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			x: 1,
		}, want: 13},

		{name: "case2", args: args{
			x: 10,
		}, want: 4},

		{name: "case3", args: args{
			x: 100,
		}, want: -86},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xFlip(tt.args.x); got != tt.want {
				t.Errorf("xFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yFlip(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			y: 1,
		}, want: 14},

		{name: "case2", args: args{
			y: 10,
		}, want: 5},

		{name: "case3", args: args{
			y: 100,
		}, want: -85},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yFlip(tt.args.y); got != tt.want {
				t.Errorf("yFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sideTag(t *testing.T) {
	type args struct {
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sd: 1,
		}, want: 16},

		{name: "case2", args: args{
			sd: 2,
		}, want: 24},

		{name: "case3", args: args{
			sd: 3,
		}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sideTag(tt.args.sd); got != tt.want {
				t.Errorf("sideTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oppSideTag(t *testing.T) {
	type args struct {
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sd: 1,
		}, want: 8},

		{name: "case2", args: args{
			sd: 2,
		}, want: 0},

		{name: "case3", args: args{
			sd: 3,
		}, want: -8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oppSideTag(tt.args.sd); got != tt.want {
				t.Errorf("oppSideTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_src(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			mv: 1,
		}, want: 1},

		{name: "case2", args: args{
			mv: 2,
		}, want: 2},

		{name: "case3", args: args{
			mv: 3,
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := src(tt.args.mv); got != tt.want {
				t.Errorf("src() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dst(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			mv: 1,
		}, want: 0},

		{name: "case2", args: args{
			mv: 2,
		}, want: 0},

		{name: "case3", args: args{
			mv: 3,
		}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dst(tt.args.mv); got != tt.want {
				t.Errorf("dst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{
			sqSrc: 1,
			sqDst: 1,
		}, want: 257},

		{name: "case2", args: args{
			sqSrc: 10,
			sqDst: 1,
		}, want: 266},

		{name: "case3", args: args{
			sqSrc: 100,
			sqDst: 1,
		}, want: 356},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_legalMove(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	type args struct {
		mv int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "case1", fields: fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args: args{mv: 1}, want: false},

		{name: "case2", fields: fields{
			SdPlayer:    2,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args: args{mv: 32}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.legalMove(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.legalMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
