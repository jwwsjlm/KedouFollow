package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

// 随机生成姓名
func GenerateRandomName() string {
	// 姓氏列表
	lastNames := []string{"张", "王", "李", "赵", "陈", "刘", "杨", "黄", "吴", "周"}

	// 名字列表
	firstNames := []string{"伟", "芳", "娜", "秀英", "敏", "静", "丽", "强", "磊", "军"}

	// 随机生成姓氏和名字
	lastName := lastNames[rand.Intn(len(lastNames))]
	firstName := firstNames[rand.Intn(len(firstNames))]

	// 返回拼接后的姓名
	return lastName + firstName
}
func GenerateRandomString(length int) string {
	targetString := "bw3EL+3NtZBrtnyF7SpepA=="

	// 解码目标字符串为字节切片
	decodedBytes, err := base64.StdEncoding.DecodeString(targetString)
	if err != nil {
		fmt.Println("解码失败:", err)
		return ""
	}

	// 生成相同长度的随机字节切片
	randomBytes := make([]byte, len(decodedBytes))
	_, err = rand.Read(randomBytes)
	if err != nil {
		fmt.Println("生成随机字节失败:", err)
		return ""
	}

	// 将随机字节切片编码为 Base64 字符串
	randomString := base64.StdEncoding.EncodeToString(randomBytes)
	//fmt.Println("随机生成的字符串:", randomString)
	return randomString
}
