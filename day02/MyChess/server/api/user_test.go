package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_login(t *testing.T) {
	type args struct {
		c *gin.Context
	}

	tests := []struct {
		name string
		args args
	}{
		//造context
		//{name: "case1",args: args{c: &gin.Context{{0xc0001d7ea0,-1 ,200} 0xc000138400, 0xc000138100, [] ,[0x7cb6e0 0x7cc520 0x808800 0x80bd60], 3, "/login ",0xc000514340, 0xc00000c048 ,0xc00000c078, {{0, 0} ,0 ,0, 0 ,0}, map[],  [] map[], map[], 0,},},},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//
		})
	}
}
