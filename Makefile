jwtgen: main.go
	go build -o $@ $<


clean: 
	rm jwtgen

.PHONY: clean
