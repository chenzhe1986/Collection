//go语言生成二维码

package main

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	qrcode.WriteFile("QQ:419511590", qrcode.Medium, 256, "./myQQ.png")
	log.Println("Build qrCode done...")
}
