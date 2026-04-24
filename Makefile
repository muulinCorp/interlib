# gen-code:
# 	protoc \
# 	-I=$(SER)/proto \
# 	-I=C:/Users/user/AppData/Local/Microsoft/WinGet/Packages/Google.Protobuf_Microsoft.Winget.Source_8wekyb3d8bbwe/include \
# 	--go_out=. --go-grpc_out=. $(SER)/proto/*.proto

gen-code:
	protoc --go_out=. --go-grpc_out=. $(SER)/proto/*.proto

remove-tag:
	git tag -d $(TAG)
	git push --delete origin $(TAG)

add-tag:
	git tag $(TAG)
	git push origin $(TAG)
