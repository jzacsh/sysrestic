OUT   :=  sysrestic
DEBV  :=  0.1
PKG   :=  $(OUT)_$(DEBV)
SRC   :=  $(shell find ./ -type f -name '*.go')

$(OUT): $(SRC)
	go build ./...
	go build -o $@

$(PKG): $(OUT)
	sudo --remove-timestamp
	mkdir -p $(@)/usr/bin/
	mkdir -p $(@)/usr/lib/$(OUT)/bin
	mkdir -p $(@)/etc/$(OUT)
	mkdir -p $(@)/DEBIAN
	mkdir -p $(@)/usr/lib/systemd/system
	cp $(OUT) $(@)/usr/lib/$(OUT)/bin/
	cp debian/system.exclude $(@)/usr/lib/$(OUT)/default.exclude
	cp debian/system.exclude $(@)/etc/$(OUT)/system.exclude
	cp debian/systemd.conf   $(@)/etc/$(OUT)/systemd.conf
	cp debian/sysrestic.service $(@)/usr/lib/systemd/system/sysrestic.service
	cp debian/sysrestic.timer   $(@)/usr/lib/systemd/system/sysrestic.timer
	echo /etc/$(OUT)/system.exclude >> $(@)/DEBIAN/conffiles
	echo /etc/$(OUT)/systemd.conf   >> $(@)/DEBIAN/conffiles
	chmod 600 $(@)/etc/$(OUT)/systemd.conf
	cp debian/control $(@)/DEBIAN/
	sed --in-place "s/VERSION_HERE/$(DEBV)/g" $(@)/DEBIAN/control
	sudo chown -R root:root $(PKG)

$(PKG).deb: $(PKG)
	dpkg-deb --build $(PKG)
	@printf 'success; now consider `sudo` removing %s\n' "$(PKG)"

all: clean test $(OUT)

deb: $(PKG).deb

test:
	go test ./...

clean:
	$(RM) -rf $(OUT) $(OUT)_*

.PHONY: test all clean deb
