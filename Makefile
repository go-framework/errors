.PHONY: protoc gorm-support var clean

# external variable
HIDE ?= @
PROTOC_INCLUDES ?= $(GOSRC)
PROTOC_OUT_FLAG ?= $(PROTOC_GO_OUT_FLAG)
PROTOC_OUT ?= .
PROTOC_OUT_M ?=
PROTOC_OUT_PLUGINS ?=
PROTOC_GOGO_OUT_ENABLE ?= false
PROTOC_GRPC_OUT_ENABLE ?= false
PROTOC_REDIS_OUT_ENABLE ?= false
PROTOC_GORM_OUT_ENABLE ?= false
PROTOC_MICRO_OUT_ENABLE ?= false

# current floder
FOLDER := .


# Golang GOSRC
GOSRC := $(shell echo $(GOPATH) | tr ':' ' ')
GOSRC := $(GOSRC:%=%/src)
GOSRC := $(shell echo $(GOSRC) | tr ' ' ':')


ifeq ($(PROTOC_GRPC_OUT_ENABLE),true)
PROTOC_OUT_PLUGINS = plugins=grpc,
endif


# protobuf
## protoc
PROTOC_GOPATH := $(subst :, ,$(GOPATH))
PROTOC_FILES := $(wildcard $(PROTOC_OUT)/$(FOLDER)/*.pb.go) $(wildcard $(PROTOC_OUT)/$(FOLDER)/*/*.pb.go)
PROTOC_INCLUDE_GOPATH := $(PROTOC_GOPATH:%=-I%/src)
PROTOC_INCLUDE_FLAG := -I. $(PROTOC_INCLUDE_GOPATH) $(PROTOC_INCLUDES)
## protoc build flag
### go out
PROTOC_GO_OUT_FLAG := --go_out=$(PROTOC_OUT_PLUGINS)$(PROTOC_OUT_M):$(PROTOC_OUT)/$(FOLDER)
### gogo out
PROTOC_GOGO_OUT_FLAG := --gogo_out=$(PROTOC_OUT_PLUGINS)$(PROTOC_OUT_M):$(PROTOC_OUT)/$(FOLDER)
## redis out
PROTOC_REDIS_OUT_FLAG := --redis_out=$(PROTOC_OUT_M):$(PROTOC_OUT)/$(FOLDER)
## gorm out
PROTOC_GORM_OUT_FLAG := --gorm_out=inject-path=$(PROTOC_OUT)/$(FOLDER),$(PROTOC_OUT_M):$(PROTOC_OUT)/$(FOLDER)
## micro out
PROTOC_MICRO_OUT_FLAG := --micro_out=$(PROTOC_OUT_M):$(PROTOC_OUT)/$(FOLDER)


ifeq ($(PROTOC_GOGO_OUT_ENABLE),true)
PROTOC_OUT_FLAG = $(PROTOC_GOGO_OUT_FLAG)
endif

ifeq ($(PROTOC_REDIS_OUT_ENABLE),true)
PROTOC_OUT_FLAG += $(PROTOC_REDIS_OUT_FLAG)
endif

ifeq ($(PROTOC_GORM_OUT_ENABLE),true)
PROTOC_OUT_FLAG += $(PROTOC_GORM_OUT_FLAG)
endif

ifeq ($(PROTOC_MICRO_OUT_ENABLE),true)
PROTOC_OUT_FLAG += $(PROTOC_MICRO_OUT_FLAG)
endif


# protoc
protoc:
	@mkdir -p $(PROTOC_OUT)/$(FOLDER)
	$(HIDE)protoc *.proto -I.:$(PROTOC_INCLUDES) $(PROTOC_OUT_FLAG)
	@echo build $(FOLDER) protobuf files succeed.



# echo variable
var:
	@echo golang src: $(GOSRC)
	@echo protoc files: $(PROTOC_FILES)
	@echo protoc out: $(PROTOC_OUT)/$(FOLDER)
	@echo PROTOC_INCLUDES: $(PROTOC_INCLUDES)
	@echo PROTOC_OUT: $(PROTOC_OUT)
	@echo PROTOC_OUT_M: $(PROTOC_OUT_M)


# clean
clean:
	@rm -rf $(PROTOC_FILES)
