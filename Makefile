pcd-api: *.go
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o pcd-api
