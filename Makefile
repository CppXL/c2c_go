
COMPILER=/usr/local/src/go/bin/go

bs:
	$(COMPILER) build -o ./bin/server c2c/cmd/server

bc:
	$(COMPILER) build -o ./bin/client c2c/cmd/client


rs:bs
	./bin/server

rc:bc
	./bin/client

ra:bs bc
	./bin/server > server.log &>/dev/null
	./bin/client



.PHONY:clean
clean:

	-rm ./bin/server
	-rm ./bin/client