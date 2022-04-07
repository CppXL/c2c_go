
COMPILER=/usr/local/src/go118/bin/go
args:=""
buildtcpagent:
	$(COMPILER) build -o ./bin/tcp/agent c2c/cmd/tcp/agent

buildtcpclient:
	$(COMPILER) build -o ./bin/tcp/client c2c/cmd/tcp/client

buildcontrol:
	$(COMPILER) build -o ./bin/control c2c/cmd/control


runtcpagent:buildtcpagent
	./bin/agent $(args)

runtcpclient:buildtcpclient
	./bin/client $(args)

runcontrol:buildtcpclient buildagent buildagentclient
	./bin/control $(args)

# ra:buildagent buildclient
# 	./bin/agent
# 	./bin/client



.PHONY:clean
clean:

	-rm ./bin/*
