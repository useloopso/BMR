build:
	go build -o bmr
run: build
	./bmr start
watch:
	reflex -s -r '\.go$$' make run
clean:
	rm bmr