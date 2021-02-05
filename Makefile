build:
	rm -rf target
	mkdir target
	cp api/http/http.yaml.loc target/http.yaml
	cp api/rpc/rpc.yaml.loc target/rpc.yaml
	go build -o target/http api/http/http.go
	go build -o target/rpc api/rpc/rpc.go

run:
	nohup target/http -conf=target/http.yaml -env=release 2>&1 > target/http.log &
	nohup target/rpc -conf=target/rpc.yaml 2>&1 > target/rpc.log &

stop:
	pkill -f target/http
	pkill -f target/rpc
