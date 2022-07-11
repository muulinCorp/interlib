
gen-code:
	protoc --go_out=. --go-grpc_out=. $(SER)/proto/*.proto