server: check-env
	export "PORT=8080" && go build -o server.bin server.go && ./server.bin

check-env:
ifndef FLIPPY_SLACKTOKENS
	$(error FLIPPY_SLACKTOKENS needs to be set. Delimit multiple tokens with ':poop:')
endif
