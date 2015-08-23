ifndef GOPATH
	GOPATH := ${PWD}/gopath
	export GOPATH
endif

pcd-api: *.go ${GOPATH}/src/github.com/spf13/viper
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o pcd-api

${GOPATH}/src/github.com/spf13/viper:
	echo "${GOPATH}"
	go get -v github.com/spf13/viper
