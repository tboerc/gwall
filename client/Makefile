DIST=.\dist

prepare:
	IF EXIST "$(DIST)" RMDIR /S /Q $(DIST)
	MKDIR $(DIST)

build: prepare
	go build -o $(DIST)\gwall.exe

start: build
	$(DIST)\gwall.exe
