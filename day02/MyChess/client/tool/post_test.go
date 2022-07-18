/**
 *@Author:sario
 *Date:2022/7/16
 *@Desc:
 */
package tool

import (
	"reflect"
	"testing"
)

var accode = "150278"

func TestGetMailAc(t *testing.T) {
	//
	t.Skip()
	type args struct {
		mail string
	}
	tests := []struct {
		name string
		args args
		want MailResp
	}{
		{name: "case1", args: args{mail: "1598273095@qq.com"},
			want: MailResp{
				Code: 200,
				Info: "Send Email Success",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMailAc(tt.args.mail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMailAc() = %v, want %v", got, tt.want)
			}
		})
	}
}

//测试时需要替换accode 有效期五分钟
func TestLogin(t *testing.T) {
	type args struct {
		mail   string
		name   string
		accode string
	}
	tests := []struct {
		name string
		args args
		want LoginResp
	}{
		{name: "case1", args: args{
			mail:   "1598273095@qq.com",
			name:   "user",
			accode: accode,
		}, want: LoginResp{
			Code:     200,
			UserName: "user",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(tt.args.mail, tt.args.name, tt.args.accode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

//同上
func TestGetToken(t *testing.T) {
	type args struct {
		mail   string
		name   string
		accode string
	}
	tests := []struct {
		name string
		args args
		want GetTokenResp
	}{
		{name: "case1", args: args{
			mail:   "1598273095@qq.com",
			name:   "user",
			accode: accode,
		}, want: GetTokenResp{
			Code:  2000,
			Info:  "succes",
			Token: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetToken(tt.args.mail, tt.args.name, tt.args.accode)
			if got.Code != 2000 || got.Info != "success" || got.Token == "" {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWinCount(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWinCount(tt.args.token); got != tt.want {
				t.Errorf("GetWinCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
