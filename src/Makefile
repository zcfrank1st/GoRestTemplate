clean:
	rm -rf ../pkg ../bin

build-simple: clean
	go build -o ../bin/go_rest_t_simple ./commandline/simple_cmd.go

build-prompt: clean
	go build -o ../bin/go_rest_t_prompt ./commandline/prompt_cmd.go

build-all: clean
	go build -o ../bin/go_rest_t_simple ./commandline/simple_cmd.go
	go build -o ../bin/go_rest_t_prompt ./commandline/prompt_cmd.go