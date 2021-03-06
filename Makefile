.PHONY: tools protoc install demo benchmark protoc-Windows protoc-Linux protoc-Darwin

protoc:
	protoc --gogoslick_out=. simple/*.proto
	# this has flags to let us use custom options (check them out better)
	protoc -I=. -I=vendor --gogo_out=. options/*.proto
	protoc -I=. -I=vendor --gogo_out=. oneof/*.proto
	protoc -I=. -I=vendor --gogo_out=. errors/*.proto

install:
	@go install ./cmd/pbdemo

demo: install
	@pbdemo

benchmark:
	cd benchmarks && go test -bench=.

### cross-platform check for installing protoc ###

ifeq ($(OS),Windows_NT)
    MYOS := Windows
else
    MYOS := $(shell uname -s)
endif

ifeq ($(MYOS),Windows)
	ZIP := protoc-3.4.0-win32.zip
endif
ifeq ($(MYOS),Darwin)  # Mac OS X
	ZIP := protoc-3.4.0-osx-x86_64.zip
endif
ifeq ($(MYOS),Linux)
	ZIP := protoc-3.4.0-linux-x86_64.zip
endif

##### scripts to install protoc and tooling on various platforms #####

protoc-Windows:
	@echo "Windows not yet supported"
	@false

protoc-Linux: /usr/local/bin/protoc

protoc-Darwin: /usr/local/bin/protoc

/usr/local/bin/protoc:
	@ curl -L https://github.com/google/protobuf/releases/download/v3.4.0/$(ZIP) > $(ZIP)
	@ unzip -q $(ZIP) -d protoc3
	@ rm $(ZIP)
	sudo mv protoc3/bin/protoc /usr/local/bin/
	@ sudo mv protoc3/include/* /usr/local/include/
	@ sudo chown `whoami` /usr/local/bin/protoc
	@ sudo chown -R `whoami` /usr/local/include/google
	@ rm -rf protoc3

tools: protoc-$(MYOS)
	@go get github.com/gogo/protobuf/proto
	@go get github.com/gogo/protobuf/gogoproto
	@go get github.com/gogo/protobuf/protoc-gen-gogofast
	@go get github.com/gogo/protobuf/protoc-gen-gogoslick
	# these are for custom extensions
	@go get github.com/gogo/protobuf/proto
	@go get github.com/gogo/protobuf/jsonpb
	@go get github.com/gogo/protobuf/protoc-gen-gogo
	@go get github.com/gogo/protobuf/gogoproto
	@# go get github.com/golang/protobuf/protoc-gen-go

