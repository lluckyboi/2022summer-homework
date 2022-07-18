package chess

import (
	"reflect"
	"testing"
)

func TestNewPositionStruct(t *testing.T) {
	tests := []struct {
		name string
		want *PositionStruct
	}{
		{"case1", &PositionStruct{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPositionStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPositionStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_startup(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "case1", fields: fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			p.startup()
			for sq := 0; sq < 256; sq++ {
				if p.UcpcSquares[sq] != cucpcStartup[sq] {
					t.Errorf("strat err")
				}
			}
		})
	}
}

func TestPositionStruct_changeSide(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"case1", fields{SdPlayer: 0, UcpcSquares: cucpcStartup, RoomId: "123456"}},
		{"case2", fields{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			sdp := p.SdPlayer
			p.changeSide()
			if sdp == p.SdPlayer {
				t.Errorf("expected:%v, got:%v", 1-sdp, p.SdPlayer)
			}
		})
	}
}

func TestPositionStruct_addPiece(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	type args struct {
		sq int
		pc int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 29,
			pc: 5,
		}},

		{"case2", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 0,
			pc: 1,
		}},

		{"case3", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 255,
			pc: 22,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			p.addPiece(tt.args.sq, tt.args.pc)
			if p.UcpcSquares[tt.args.sq] != tt.args.pc {
				t.Errorf("落子失败")
			}
		})
	}
}

func TestPositionStruct_delPiece(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	type args struct {
		sq int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 52,
		}},

		{"case2", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 53,
		}},

		{"case3", fields{
			SdPlayer:    0,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			sq: 54,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			p.delPiece(tt.args.sq)
			if p.UcpcSquares[tt.args.sq] != 0 {
				t.Errorf("删除棋子失败 %v", tt.args.sq)
			}
		})
	}
}

func TestPositionStruct_movePiece(t *testing.T) {
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
		want   int
	}{

		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mv: 1,
		}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.movePiece(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.movePiece() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_makeMove(t *testing.T) {
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
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mv: 1,
		}, true},

		{"case2", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mv: 256,
		}, true},

		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mv: 64,
		}, true},

		//不被将死都为true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.makeMove(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.makeMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_generateMoves(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	type args struct {
		mvs []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{

		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mvs: make([]int, MaxGenMoves),
		}, 44},
		//共44种走法
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.generateMoves(tt.args.mvs); got != tt.want {
				t.Errorf("PositionStruct.generateMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_checked(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, false},
		//判断是否将军
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.checked(); got != tt.want {
				t.Errorf("PositionStruct.checked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_undoMovePiece(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	type args struct {
		mv         int
		pcCaptured int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, args{
			mv:         1,
			pcCaptured: 20,
		}},
		//撤销一步
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			p.undoMovePiece(tt.args.mv, tt.args.pcCaptured)
		})
	}
}

func TestPositionStruct_isMate(t *testing.T) {
	type fields struct {
		SdPlayer    int
		UcpcSquares [256]int
		RoomId      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"case1", fields{
			SdPlayer:    1,
			UcpcSquares: cucpcStartup,
			RoomId:      "123456",
		}, false},
		//判断是否将死
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				SdPlayer:    tt.fields.SdPlayer,
				UcpcSquares: tt.fields.UcpcSquares,
				RoomId:      tt.fields.RoomId,
			}
			if got := p.isMate(); got != tt.want {
				t.Errorf("PositionStruct.isMate() = %v, want %v", got, tt.want)
			}
		})
	}
}
