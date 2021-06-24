.PHONY: build clean pack deploy ship

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o simple-go-app cmd/main.go

run: build
	./simple-go-app

clean:
	rm ./simple-go-app

pack:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build
	docker build -t rishiskeshav91/simple-go-app:latest .

upload:
	docker push rishiskeshav91/simple-go-app:latest

deploy:
	kubectl apply -f k8/deployment.yml
	kubectl apply -f k8/service.yml
	make rolling-restart
	# Having an issue with port-forward. Running this separatly works okay.
	#make port-forward

port-forward:
	kubectl port-forward service/simple-go-app 7000:7000

rolling-restart:
	kubectl rollout restart deployment/simple-go-app

ship: pack upload deploy clean