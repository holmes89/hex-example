TAG=$(git describe --abbrev=0 --tags)

.PHONY: deploy-lambda build-docker-ticket b uild-docker-user build-docker-all
deploy-lambda:
	GOOS=linux go build -o main cmd/ticketAPI/main.go
	zip main.zip cmd/ticketAPI/main
	aws s3 cp main.zip s3://hex-lambda/$TAG/main.zip
	cd terraform/prod/
	terraform apply -var "app_version=$TAG" -auto-approve
	cd ../../
	rm -rf main.zip

build-docker-user:
	docker build -f dockerfiles/Dockerfile.user -t holmes89/hex-user-api .

build-docker-ticket:
	docker build -f dockerfiles/Dockerfile.ticket -t holmes89/hex-ticket-api .

build-docker-all: build-docker-user build-docker-ticket
