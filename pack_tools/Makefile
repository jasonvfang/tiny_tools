#Makefile for pack tool

.PHONY: all clean 

all:
	go build all
	-go install all
	mv bin/main bin/ota_pack

clean:
	rm -rf bin
	rm -rf pkg 

