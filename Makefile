all: webclient webserver clean

webclient:
	go install github.houston.softwaregrp.net/CSB/webserver/cmd/webclient

webserver:
	go install github.houston.softwaregrp.net/CSB/webserver/cmd/webserver

clean:
	rm -rf ./buildtmp
