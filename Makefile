# deploy function
deploy:
	go mod vendor
	gcloud alpha functions deploy ping \
		--entry-point F \
		--memory 128MB \
		--runtime go111 \
		--trigger-http
test:
	go test . -v