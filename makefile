TARGETDIR=.\deploy
proj=github.com\billgraziano\go-windows-svc
sha1ver := $(shell git rev-parse HEAD)
test := $(shell date /t)


all: vet test  buildEXE

vet:
	go vet -all -shadow .\cmd\gosvc
	go vet -all -shadow .\app

test: 
	go.exe test -timeout 30s $(proj)\app

# The sha1 stuff isn't working as of now
buildEXE:
	go build -o "$(TARGETDIR)\gosvc.exe" -a -ldflags "-X main.sha1ver=$(sha1ver)" .\cmd\gosvc  
