build:
	go build -o main src/*.go

package-maclinux:
	go build src/*.go -o main
	tar -czvf dvisualize.tar.gz main scripts/

package-windows:
	GOOS=windows GOARCH=amd64 go -o main.exe src
	zip -r dvisualize.zip main.exe scripts/

run:
	./main

clean:
	rm -f ./main
