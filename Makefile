.PHONY: proto

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/blindsig/proto/blindsig.proto
bench:
	cd ./pkg/blindsig; go test -run=XXX -bench=.
test:
	cd ./pkg/blindsig; go test -run=.