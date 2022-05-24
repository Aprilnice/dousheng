package doushengjwt

import (
	"bou.ke/monkey"
	"dousheng/pkg/snowflaker"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// TestGenerateToken 测试生成token的功能
func TestGenerateToken(t *testing.T) {
	// 函数打桩
	monkey.Patch(tokenExpiresAt, func() int64 {
		tokenExpireDuration := 24 * time.Hour * time.Duration(7)
		return time.Now().Add(tokenExpireDuration).Unix()
	})
	snowflaker.Init("2022-05-01", 1)
	userid := snowflaker.NextID()
	username := "douyin"
	token, err := GenerateToken(username, userid)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(userid)
	fmt.Println(token)
	assert.NotEqual(t, len(token), 0)

}

// TestParseToken 测试解析token的功能
func TestParseToken(t *testing.T) {
	// 函数打桩
	monkey.Patch(tokenExpiresAt, func() int64 {
		tokenExpireDuration := 24 * time.Hour * time.Duration(7)
		return time.Now().Add(tokenExpireDuration).Unix()
	})
	// 生成id 和 用户名
	snowflaker.Init("2022-05-01", 1)
	userid := snowflaker.NextID()
	username := "douyin"
	token, err := GenerateToken(username, userid)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(userid)
	fmt.Println(token)
	claims, err := ParseToken(token)
	assert.Equal(t, username, claims.UserName)
	assert.Equal(t, userid, claims.UserID)

}

// BenchmarkGenerateToken 性能测试
func BenchmarkGenerateToken(b *testing.B) {
	// 函数打桩
	monkey.Patch(tokenExpiresAt, func() int64 {
		tokenExpireDuration := 24 * time.Hour * time.Duration(7)
		return time.Now().Add(tokenExpireDuration).Unix()
	})
	// 生成id 和 用户名
	snowflaker.Init("2022-05-01", 1)

	username := "douyin"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userid := snowflaker.NextID()
		GenerateToken(username, userid)
	}
}
