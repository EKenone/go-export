build:
	rm -rf target
	mkdir target
	cp api/http/conf.yaml target/http.yaml
	cp api/rpc/conf.yaml target/rpc.yaml
	go build -o target/http api/http/http.go
	go build -o target/rpc api/rpc/rpc.go

run:
	nohup target/http -conf=target/http.yaml -env=release 2>&1 > target/http.log &
	nohup target/rpc -conf=target/rpc.yaml 2>&1 > target/rpc.log &

stop:
	pkill -f target/http
	pkill -f target/rpc
