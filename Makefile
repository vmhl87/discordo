all:
	go build .

test:
	./discordo -token `cat ~/misc/tkn`

install:
	sudo cp ./discordo /usr/local/bin/
