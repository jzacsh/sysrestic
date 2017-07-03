OUT  :=  sysrestic
SRC  :=  $(shell find ./ -type f -name '*.go')

all: clean test $(OUT)

$(OUT): $(SRC)
	go build -o $@

test:
	go test ./...

clean:
	$(RM) -f $(OUT)

.PHONY: test all clean
