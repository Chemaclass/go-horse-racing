build:
	go build -o out/horse-racing

run: build
	./out/horse-racing
