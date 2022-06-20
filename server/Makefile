DIST=.\dist

prepare:
	IF EXIST "$(DIST)" RMDIR /S /Q $(DIST)
	MKDIR $(DIST)

build: prepare
	go build -o $(DIST)\server.exe

start: build
	$(DIST)\server.exe
