.PHONY: clean build package

clean: 
	rm -rf ./hello package.zip bootstrap.zip
	
build:
	go get github.com/aws/aws-lambda-go/lambda
	GOOS=linux GOARCH=amd64 go build -o bootstrap .

package:
	zip -r bootstrap.zip bootstrap
