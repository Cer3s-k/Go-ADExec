LDFLAGS = -ldflags "-s -w -extldflags '-static'"

.PHONY: all
all: clean win linux
	
.PHONY: win
win: 
	GOOS=windows GOARCH=amd64 go build -tags release $(LDFLAGS) -o bin/Go-ADExec-win.exe

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -tags release $(LDFLAGS) -o bin/Go-ADExec-lin.exe

clean:
	del /q bin\*

