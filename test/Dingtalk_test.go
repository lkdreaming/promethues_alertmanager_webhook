package test

import (
	"github.com/blinkbean/dingtalk"
	"testing"
)

func TestPostBot(t *testing.T) {
	// 使用方法: https://github.com/blinkbean/dingtalk
	dingToken := "6100b926d9487209463f3f745fb67ba78484c38c47c73b4c25f0103624bfd289"
	secret := "SEC88eb7edb022d3140ff1dc10369efca030491c9516e032807302d262f3f4ec56f"
	cli := dingtalk.InitDingTalkWithSecret(dingToken, secret)
	err := cli.SendTextMessage("刘兆祎 刘兆熙 大帅哥和小美女")
	if err != nil {
		t.Error(err)
	}
	text := "### 刘兆祎 刘兆熙 大帅哥和小美女\n" +
		"> 哈哈哈 \n\n" +
		"```java\n\n" +
		"public class OperationBoot {\n\n    public static void main(String[] args) {\n\n        SpringApplication.run(OperationBoot.class, args);\n\n    }\n\n}\n\n" +
		"```\n\n" +
		"**我爱你们**"
	err = cli.SendMarkDownMessage("我的小可爱们", text)

	if err != nil {
		t.Error(err)
	}
}
