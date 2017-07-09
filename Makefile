OUT   :=  sysrestic
DEBV  :=  0.1
DEBT  :=  $(OUT)_$(DEBV)
SRC   :=  $(shell find ./ -type f -name '*.go')

all: clean test $(OUT)

$(OUT): $(SRC)
	go build ./...
	go build -o $@

$(DEBT): $(OUT)
	sudo --remove-timestamp
	mkdir -p $(@)/usr/lib/$(OUT)/bin
	mkdir -p $(@)/etc
	mkdir -p $(@)/DEBIAN
	cp $(OUT) $(@)/usr/lib/$(OUT)/bin/
	cp debian/system.exclude $(@)/usr/lib/$(OUT)/default.exclude
	cp debian/system.exclude $(@)/etc/$(OUT).exclude
	printf 'etc/%s.exclude\n' $(OUT) >> $(@)/DEBIAN/conffiles
	cp debian/systemd.conf $(@)/etc/$(OUT).conf
	chmod 600 $(@)/etc/$(OUT).conf
	printf 'etc/%s.conf\n' $(OUT) >> $(@)/DEBIAN/conffiles
	cp debian/control $(@)/DEBIAN/
	sed --in-place "s/VERSION_HERE/$(DEBV)/g" $(@)/DEBIAN/control
	sudo chown -R root:root $(DEBT)

$(DEBT).deb: $(DEBT)
	@echo "not yet implemented" >&2
	@test -d /dev/null

test:
	go test ./...

clean:
	$(RM) -rf $(OUT) $(OUT)_*

.PHONY: test all clean
