OS=linux darwin
ARCHS=arm amd64
APPS=mntp
BUILD_VERSION=$(shell git describe --tags)
PACKAGE_NAME=github.com/hysios/mntp
# VERSION_OPT=-X $(PACKAGE_NAME)/apps//cmd.Version=$(BUILD_VERSION)

build:
	@-for os in $(OS) ; do \
		for arch in $(ARCHS); do \
			for app in $(APPS); do \
				echo build versoin $(BUILD_VERSION); \
				GOOS=$$os GOARCH=$$arch go build -ldflags="-s -X $(PACKAGE_NAME)/$$app.main.Version=$(BUILD_VERSION)" -o bin/$$app-$$os-$$arch ./example; \
			done; \
		done; \
	done

