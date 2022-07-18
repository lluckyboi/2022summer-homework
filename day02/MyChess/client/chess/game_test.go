package chess

import (
	_ "image/png"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

func TestNewGame(t *testing.T) {
	NewGame(-2)
}

func TestGame_Update(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
		side           int
	}
	type args struct {
		screen *ebiten.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "case1",
			fields: fields{
				sqSelected:     20,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         nil,
				audios:         nil,
				audioContext:   nil,
				singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
				side:           1,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
				side:           tt.fields.side,
			}
			if err := g.Update(tt.args.screen); (err != nil) != tt.wantErr {
				t.Errorf("Game.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGame_Layout(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
		side           int
	}
	type args struct {
		outsideWidth  int
		outsideHeight int
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantScreenWidth  int
		wantScreenHeight int
	}{
		{name: "case1", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         nil,
			audios:         nil,
			audioContext:   nil,
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           1,
		}, args: args{
			outsideWidth:  520,
			outsideHeight: 576,
		},
			wantScreenWidth:  520,
			wantScreenHeight: 576,
		},

		{name: "case1", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         nil,
			audios:         nil,
			audioContext:   nil,
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           1,
		}, args: args{
			outsideWidth:  555,
			outsideHeight: 555,
		},
			wantScreenWidth:  520,
			wantScreenHeight: 576,
		},

		//固定宽高 520 576
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
				side:           tt.fields.side,
			}
			gotScreenWidth, gotScreenHeight := g.Layout(tt.args.outsideWidth, tt.args.outsideHeight)
			if gotScreenWidth != tt.wantScreenWidth {
				t.Errorf("Game.Layout() gotScreenWidth = %v, want %v", gotScreenWidth, tt.wantScreenWidth)
			}
			if gotScreenHeight != tt.wantScreenHeight {
				t.Errorf("Game.Layout() gotScreenHeight = %v, want %v", gotScreenHeight, tt.wantScreenHeight)
			}
		})
	}
}

func TestGame_loadResource(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
		side           int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "case1", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         make(map[int]*ebiten.Image),
			audios:         make(map[int]*audio.Player),
			audioContext:   &audio.Context{},
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           1,
		}, want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
				side:           tt.fields.side,
			}
			if got := g.loadResource(); got != tt.want {
				t.Errorf("Game.loadResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_clickSquare(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
		side           int
	}
	tests := []struct {
		name   string
		fields fields
		sq     int
		want   bool
	}{
		{name: "case1", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         make(map[int]*ebiten.Image),
			audios:         make(map[int]*audio.Player),
			audioContext:   &audio.Context{},
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           1,
		}, sq: 51, want: true},

		{name: "case2", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         make(map[int]*ebiten.Image),
			audios:         make(map[int]*audio.Player),
			audioContext:   &audio.Context{},
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           2,
		}, sq: 51, want: true},

		{name: "case3", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      true,
			showValue:      "win",
			images:         make(map[int]*ebiten.Image),
			audios:         make(map[int]*audio.Player),
			audioContext:   &audio.Context{},
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           2,
		}, sq: 51, want: true},

		{name: "case3", fields: fields{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      true,
			showValue:      "win",
			images:         make(map[int]*ebiten.Image),
			audios:         make(map[int]*audio.Player),
			audioContext:   &audio.Context{},
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           2,
		}, sq: 51, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
				side:           tt.fields.side,
			}
			if tt.name == "case3" {
				g.sqSelected = 67
			}

			g.clickSquare(&ebiten.Image{}, tt.sq)
		})
	}
}
