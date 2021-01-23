
swagger_instatll:
	which ${GOPATH}/bin/swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger_generate:
	${GOPATH}/bin/swagger generate spec -o ./swagger.yaml --scan-models

swagger:
	make swagger_instatll && make swagger_generate

.PHONY: swagger
