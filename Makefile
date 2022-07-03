run:
	@go run cmd/main.go
	
clean:
	@docker system prune

dbuild:
	@docker image build -f Dockerfile -t spellserviceimage .
	@docker container run -p 8080:8080 --detach --name forum spellserviceimage:1.13.0