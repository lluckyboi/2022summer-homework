/**
 *@Author:sario
 *Date:2022-7-22
 *@Desc:
 */
package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendMail(t *testing.T) {
	type args struct {
		mails []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{name: "case1", args: args{mails: []string{"1598273095@qq.com"}}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.mails); (err != nil) && err != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authHandler(t *testing.T) {
	t.Skip()
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{c: &gin.Context{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authHandler(tt.args.c)
		})
	}
}

func TestJWTAuthMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want func(c *gin.Context)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JWTAuthMiddleware(); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("JWTAuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
