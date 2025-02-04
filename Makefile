build:
	go build -o main src/*.go

package-maclinux:
	go build -o main src/*.go 
	tar -czvf dvisualize.tar.gz main scripts/

package-windows:
	GOOS=windows GOARCH=amd64 go build -o main.exe src/*.go
	zip -r dvisualize.zip main.exe scripts/

run:
	./main

clean:
	rm -f ./main
