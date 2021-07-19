Running your executable go program without building (it compiles and runs):
	go run <program>.go

Running your executable go program after building
	cd <program dir>; go build; ./<program>
	# If you want the binary to have a particular name
	cd <program dir>; go build -o <name>; ./<name>

Building the executable for WindowsOS from linux
	cd <program dir>; GOOS=windows GOARCH=amd64 go build -o <appname>.exe
Building the executable for MacOS from linux
	cd <program dir>; GOOS=darwin GOARCH=amd64 go build -o <appname>.exe

Formatting go code
	gofmt -w <file>.go
