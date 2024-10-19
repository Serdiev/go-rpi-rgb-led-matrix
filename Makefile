

test:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go run examples/basic/main.go

test:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 MATRIX_TERMINAL_EMULATOR=1 go run examples/animation/main.go
	
emu:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go run examples/emu/main.go