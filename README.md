# gopher



# mac上编译linux和windows二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 

# linux上编译mac和windows二进制
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

# windows上编译mac和linux二进制
SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build main.go
SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go

GOOS - Target Operating System	GOARCH - Target Platform

```
android	arm
darwin	386
darwin	amd64
darwin	arm
darwin	arm64
dragonfly	amd64
freebsd	386
freebsd	amd64
freebsd	arm
linux	386
linux	amd64
linux	arm
linux	arm64
linux	ppc64
linux	ppc64le
linux	mips
linux	mipsle
linux	mips64
linux	mips64le
netbsd	386
netbsd	amd64
netbsd	arm
openbsd	386
openbsd	amd64
openbsd	arm
plan9	386
plan9	amd64
solaris	amd64
windows	386
windows	amd64
```
Golang交叉编译各个平台的二进制文件
https://studygolang.com/articles/14376

Golang交叉编译中的那些坑
https://blog.csdn.net/Three_dog/article/details/94640507?utm_medium=distribute.pc_feed_404.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_feed_404.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecas


