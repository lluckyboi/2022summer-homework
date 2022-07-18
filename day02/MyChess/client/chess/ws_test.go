package chess

import (
	"testing"
)

func TestUpdateOtherBoard(t *testing.T) {
	type args struct {
		g *Game
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{g: &Game{
			sqSelected:     20,
			mvLast:         0,
			bFlipped:       false,
			bGameOver:      false,
			showValue:      "",
			images:         nil,
			audios:         nil,
			audioContext:   nil,
			singlePosition: &PositionStruct{SdPlayer: 1, UcpcSquares: cucpcStartup, RoomId: "123456"},
			side:           1},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateOtherBoard(tt.args.g)
			//_, message, err := Conn.ReadMessage()
			//if err != nil {
			//	log.Println("read:", err)
			//	wg.Done()
			//	return
			//}
			//sg:=&PositionStruct{}
			//json.Unmarshal(message, &sg)
			//if sg!=tt.args.g.singlePosition{
			//	t.Errorf("update other borad err %v",sg)
			//}
		})
	}
}
