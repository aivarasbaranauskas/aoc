mains = $(wildcard cmd/*/*/main.go) $(wildcard cmd/*/*/*/main.go)
inputs = $(patsubst %/main.go,%/input.txt,$(mains))
bins = $(patsubst cmd/%/main.go,bin/%,$(mains))

%/input.txt:
	touch $@

bin/%: cmd/%/main.go
	go build -o $@ $^

init-inputs: $(inputs)

all: $(bins)