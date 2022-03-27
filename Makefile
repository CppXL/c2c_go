
COMPILER=/usr/local/src/go118/bin/go
args:=""
bs:
	$(COMPILER) build -o ./bin/server c2c/cmd/server

bc:
	$(COMPILER) build -o ./bin/client c2c/cmd/client

bsc:
	$(COMPILER) build -o ./bin/control c2c/cmd/control


rs:bs
	./bin/server $(args)

rc:bc
	./bin/client $(args)

rsc:bc bs bsc
	./bin/control $(args)

ra:bs bc
	./bin/server > server.log &>/dev/null
	./bin/client



.PHONY:clean
clean:

	-rm ./bin/server
	-rm ./bin/client