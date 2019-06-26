package main

import "github.com/skip2/go-qrcode"

func main() {
	_ = qrcode.WriteFile("https://apply.scry.info/apply", qrcode.Highest, 512, "apply.png")
	_ = qrcode.WriteFile("https://apply.scry.info/sign", qrcode.Highest, 512, "sign.png")
}
