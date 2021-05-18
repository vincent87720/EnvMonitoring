WINDOWS := windows
LINUX := linux
PACKAGE := github.com/vincent87720/EnvMonitoring/cmd/EnvMonitoring
OPS := $(WINDOWS) $(LINUX)

build: rmbin release cpyaml

rmbin:
	rm -rf ./bin/*

cpyaml:
	for op in $(OPS) ; do \
		cp ./settings.yaml bin/$$op ; \
	done

##########BUILD##########
.PHONY: buildwindows
buildwindows:
	GOOS=$(WINDOWS) GOARCH=amd64 go build -o bin/$(WINDOWS)/EnvMonitoring.exe $(PACKAGE)

.PHONY: buildlinux
buildlinux:
	GOOS=$(LINUX) GOARCH=amd64 go build -o bin/$(LINUX)/EnvMonitoring $(PACKAGE)

.PHONY: release
release: buildwindows buildlinux



##########RUN##########
run:
	go run $(PACKAGE)