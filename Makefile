# deploy function
deploy:
	gcloud alpha functions deploy ping \
		--entry-point F \
		--memory 128MB \
		--runtime go111 \
		--trigger-http
test:
	go test . -v