all:
	go build . && ldid -e /bin/cp /tmp/ent.xml && ldid -S/tmp/ent.xml ./discordo

install:
	sudo cp ./discordo /usr/local/bin/
