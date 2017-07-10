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
	mkdir -p $(@)/usr/bin/
	mkdir -p $(@)/usr/lib/$(OUT)/bin
	mkdir -p $(@)/etc/$(OUT)
	mkdir -p $(@)/DEBIAN
	cp $(OUT) $(@)/usr/lib/$(OUT)/bin/
	cp $(OUT) $(@)/usr/bin/$(OUT)
	cp debian/system.exclude $(@)/usr/lib/$(OUT)/default.exclude
	cp debian/system.exclude $(@)/etc/$(OUT)/system.exclude
	cp debian/systemd.conf $(@)/etc/$(OUT)/dummy-systemd.conf
	echo /etc/$(OUT)/default.exclude >> $(@)/DEBIAN/conffiles
	echo /etc/$(OUT)/dummy-systemd.conf >> $(@)/DEBIAN/conffiles
	chmod 600 $(@)/etc/$(OUT)/dummy-systemd.conf
	cp debian/control $(@)/DEBIAN/
	sed --in-place "s/VERSION_HERE/$(DEBV)/g" $(@)/DEBIAN/control
	sudo chown -R root:root $(DEBT)

$(DEBT).deb: $(DEBT)
	dpkg-deb --build $(DEBT)
	@printf 'success; now consider `sudo` removing %s\n' "$(DEBT)"

test:
	go test ./...

clean:
	$(RM) -rf $(OUT) $(OUT)_*

.PHONY: test all clean
