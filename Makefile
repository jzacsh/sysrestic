OUT   :=  sysrestic
DEBV  :=  0.1
PKG   :=  $(OUT)_$(DEBV)
SRC   :=  $(shell find ./ -type f -name '*.go')

$(OUT): $(SRC)
	go build ./...
	go build -o $@

$(PKG): $(OUT)
	sudo --remove-timestamp
	mkdir -p $(@)/usr/lib/$(OUT)/bin
	mkdir -p $(@)/usr/lib/$(OUT)/systemd
	mkdir -p $(@)/etc/
	mkdir -p $(@)/DEBIAN
	mkdir -p $(@)/usr/lib/systemd/system/
	mkdir -p $(@)/etc/systemd/system/$(OUT).service.d
	cp $(OUT) $(@)/usr/lib/$(OUT)/bin/
	cp debian/system.exclude $(@)/usr/lib/$(OUT)/default.exclude
	cp debian/system.exclude $(@)/etc/$(OUT).exclude
	cp debian/systemd.conf   $(@)/etc/systemd/system/$(OUT).service.d/systemd.conf
	cp debian/systemd.conf   $(@)/usr/lib/$(OUT)/systemd/example.conf
	cp debian/$(OUT).service $(@)/usr/lib/$(OUT)/systemd/$(OUT).service
	cp debian/$(OUT).timer   $(@)/usr/lib/$(OUT)/systemd/$(OUT).timer
	cd $(@)/usr/lib/systemd/system/ && ln -s /usr/lib/$(OUT)/systemd/$(OUT).service
	cd $(@)/usr/lib/systemd/system/ && ln -s /usr/lib/$(OUT)/systemd/$(OUT).timer
	echo /etc/$(OUT).exclude >> $(@)/DEBIAN/conffiles
	echo /etc/systemd/system/$(OUT).service.d/systemd.conf >> $(@)/DEBIAN/conffiles
	chmod 600 $(@)/etc/systemd/system/$(OUT).service.d/systemd.conf
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
