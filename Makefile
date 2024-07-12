
gen-code:
	protoc --go_out=. --go-grpc_out=. $(SER)/proto/*.proto

remove-tag:
	git tag -d $(TAG)
	git push --delete origin $(TAG)

add-tag:
	git tag $(TAG)
	git push origin $(TAG)
