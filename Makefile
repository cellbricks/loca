.PHONY: proto

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/blindsig/proto/blindsig.proto

.PHONY: bench deep
bench:
	cd ./pkg/blindsig; go test -run=XXX -bench=.
deep:
	cd ./pkg/blindsig; go test -run=. -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out | tee bench.log

.PHONY: test
test:
	cd ./pkg/blindsig; go test -run=.
