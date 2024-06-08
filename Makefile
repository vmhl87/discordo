all:
	go build . && ldid -e /usr/bin/cp > /tmp/ent.xml && sudo ldid -S/tmp/ent.xml discordo

test:
	./discordo -token `cat /var/mobile/misc/tkn`

install:
	sudo cp ./discordo /usr/local/bin/
