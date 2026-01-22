# Douyu Golang Application Standard Makefile

SHELL:=/bin/bash
BASE_PATH:=$(shell dirname $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
BUILD_PATH:=$(BASE_PATH)/build

TITLE:=$(shell basename $(BASE_PATH))
VCS_INFO:=$(shell $(BUILD_PATH)/script/shell/vcs.sh)
BUILD_TIME:=$(shell date +%Y-%m-%d--%T)
APP_PKG:=$(shell $(BUILD_PATH)/script/shell/apppkg.sh)
MINERVA:=$(APP_PKG)/vendor/git.dz11.com/vega/minerva/application
MINERVA_NEW:=git.dz11.com/vega/minerva/application
ENABLE_GO_MOD:=$(shell $(BUILD_PATH)/script/shell/gomod.sh)
JUPITER:= github.com/douyu/jupiter/pkg
JUPITER_FLAGS:= "-X $(JUPITER).appName=$(APP_NAME) -X $(JUPITER).appID=$(APP_ID) -X $(JUPITER).buildTime=$(BUILD_TIME)" 

#if [ -d $(APP_PKG)/vendor/git.dz11.com/vega/minerva/application ] then
#	MINERVA:=$(APP_PKG)/vendor/git.dz11.com/vega/minerva/application
#else
#	MINERVA:=$(APP_PKG)/vendor/git.dz11.com/vega/application
#fi

LDFLAGS:="-X $(MINERVA).vcsInfo=$(VCS_INFO) -X $(MINERVA).buildTime=$(BUILD_TIME) -X $(MINERVA).name=$(APP_NAME) -X $(MINERVA).id=$(APP_ID) "$(JUPITER_FLAGS)  
LDFLAGS_GOMOD:="-X $(MINERVA_NEW).vcsInfo=$(VCS_INFO) -X $(MINERVA_NEW).buildTime=$(BUILD_TIME) -X $(MINERVA_NEW).name=$(APP_NAME) -X $(MINERVA_NEW).id=$(APP_ID) "$(JUPITER_FLAGS)  

all:print build zip
alltar:print build tar
#all:print fmt lint vet test build tar

print:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making print<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@echo SHELL:$(SHELL)
	@echo BASE_PATH:$(BASE_PATH)
	@echo BUILD_PATH:$(BUILD_PATH)
	@echo TITLE:$(TITLE)
	@echo VCS_INFO:$(VCS_INFO)
	@echo BUILD_TIME:$(BUILD_TIME)
	@echo APP_PKG:$(APP_PKG)
	@echo MINERVA:$(MINERVA)
	@echo BINS:$(BINS)
	@echo APP_NAME:$(APP_NAME)
	@echo LDFLAGS:$(LDFLAGS)
	@echo LDFLAGS_GOMOD:$(LDFLAGS_GOMOD)
	@echo ENABLE_GO_MOD:$(ENABLE_GO_MOD)
	@echo -e "\n"

fmt:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making fmt<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	go fmt $(TITLE)/app/...
	go fmt $(TITLE)/agent/...
	@echo -e "\n"

lint:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making lint<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@OUT=`$(GOPATH)/bin/golint $(BASE_PATH)/app/... | grep -v pb.go`;\
	if [ "$$OUT" != "" ]; then\
		echo $$OUT;\
		exit 1;\
	else\
		echo "lint success";\
	fi
	@echo -e "\n"

vet:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making vet<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@# 此处要加"2>&1"否则VETOUT变量获取不到vet输出信息
	@VETOUT=`go vet -v $(TITLE)/app/... 2>&1`;\
	if [ "$$VETOUT" != "" ]; then\
		echo $$VETOUT;\
		exit 1;\
	else\
		echo "vet success";\
	fi
	@echo -e "\n"

test:

build:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(BUILD_PATH)/script/shell/*.sh
	@if [ $(ENABLE_GO_MOD) == 0 ]; then \
		$(BUILD_PATH)/script/shell/build.sh $(LDFLAGS) $(BINS); \
	else \
		$(BUILD_PATH)/script/shell/buildgomod.sh $(LDFLAGS_GOMOD) $(BINS); \
	fi
	@echo -e "\n"

tar:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making tar<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	$(eval TAR=$(BASE_PATH)/$(TITLE)_$(VCS_INFO).tar.gz)
	@cd $(BASE_PATH) && tar zcf $(TAR) bin config build Gopkg.* resource >/dev/null
	@echo $(TAR)
	@echo -e "\n"

zip:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making zip<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@ZIP=`cd $(BASE_PATH) && zip -r $(PUB_ZIP) bin config lib template* Gopkg.* resource >/dev/null`
	@if [ "$(PUB_ZIP)" == "" ]; then\
		echo "no pub zip";\
	else\
		echo $$ZIP;\
	fi
	@echo $(PUB_ZIP)
	@echo -e "\n"

