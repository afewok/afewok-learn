package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5694. 设计一个验证系统
 */

func Test_leetcode_5694(t *testing.T) {
	authenticationManager := Constructor5694(13)
	authenticationManager.Renew("ajvy", 1)
	fmt.Println(authenticationManager.CountUnexpiredTokens(3))
	fmt.Println(authenticationManager.CountUnexpiredTokens(4))
	authenticationManager.Generate("fuzxq", 5)
	authenticationManager.Generate("izmry", 7)
	authenticationManager.Renew("puv", 12)
	authenticationManager.Generate("ybiqb", 13)
	authenticationManager.Generate("gm", 14)
	fmt.Println(authenticationManager.CountUnexpiredTokens(15))
	fmt.Println(authenticationManager.CountUnexpiredTokens(18))
	fmt.Println(authenticationManager.CountUnexpiredTokens(19))
	authenticationManager.Renew("ybiqb", 21)
	fmt.Println(authenticationManager.CountUnexpiredTokens(23))
	fmt.Println(authenticationManager.CountUnexpiredTokens(25))
	fmt.Println(authenticationManager.CountUnexpiredTokens(26))
	authenticationManager.Generate("aqdm", 28)
	fmt.Println(authenticationManager.CountUnexpiredTokens(29))
	authenticationManager.Renew("puv", 30)
}

type AuthenticationManager struct {
	TimeToLive int
	mp         map[string]int
}

func Constructor5694(timeToLive int) AuthenticationManager {
	return AuthenticationManager{timeToLive, make(map[string]int)}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.mp[tokenId] = currentTime + this.TimeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if t, OK := this.mp[tokenId]; OK {
		if t > currentTime {
			this.mp[tokenId] = currentTime + this.TimeToLive
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	result := 0
	for _, v := range this.mp {
		if v > currentTime {
			result++
		}
	}
	return result
}
