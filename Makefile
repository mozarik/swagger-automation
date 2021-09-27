

# Installing using docker pull https://goswagger.io/install.html
check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger-generate:
	swagger generate spec -o ./swagger.yaml --scan-models

# To View fancy webpage BATCHEST WOW UI UX
swagger-serve:
	swagger serve -F=swagger swagger.yaml


swagger-install:
	./install-goswagger.sh

