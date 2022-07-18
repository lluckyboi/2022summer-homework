/**
 *@Author:sario
 *Date:2022/7/16
 *@Desc:
 */
package tool

import "testing"

func TestLengthCheck(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{name: "case1", args: args{ss: "123"}, want: "", want1: true},
		{name: "case2", args: args{ss: "1234567891234567891234567"}, want: "超过长度限制！", want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LengthCheck(tt.args.ss)
			if got != tt.want {
				t.Errorf("LengthCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LengthCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPostLengthCheck(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{name: "case1", args: args{ss: "123"}, want: "", want1: true},
		{name: "case2", args: args{ss: "1234567891234567891234567"}, want: "超过长度限制！", want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PostLengthCheck(tt.args.ss)
			if got != tt.want {
				t.Errorf("PostLengthCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PostLengthCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScoreCheck(t *testing.T) {
	type args struct {
		score float32
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{name: "case1", args: args{score: 7.1}, want: "", want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ScoreCheck(tt.args.score)
			if got != tt.want {
				t.Errorf("ScoreCheck() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ScoreCheck() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
