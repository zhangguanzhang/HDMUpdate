# HDMUpdate
cli update the HDM firmware
批量更新服务器的HDM固件,因此写了一个小cli工具
h3c 服务器批量更新服务器的HDM固件,需要自行搭建一个tftp
编译
```
go build .
```
命令帮助
```
main.exe
Error: required flag(s) "tftp" not set
Usage:
  HDMUpdate [flags]

Flags:
  -f, --filename string      the HDM bin filename from the tftp root path
  -h, --help                 help for HDMUpdate
  -p, --password string      HDM login password (default "Password@_")
      --processlimit uint8   image process limit (default 1)
  -t, --tftp string          HDM get bin file from the tftp ip
  -u, --user string          HDM login user (default "admin")
```
使用tftp `10.0.23.39`上的`HDM-1.11.29_signed.bin`去更新10.0.23.34的固件
```
main.exe -t 10.0.23.39 -f HDM-1.11.29_signed.bin 10.0.23.34
```
批量并发2个更新
```
main.exe -t 10.0.23.39 -f HDM-1.11.29_signed.bin --processlimit 2 10.0.23.34 10.0.23.35 10.0.23.36
```
