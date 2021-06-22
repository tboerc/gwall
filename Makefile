BIN_FOLDER=.\bin
LIB_FOLDER=.\lib
BIN_LIB_FOLDER=.\bin\lib\

prepare:
	IF EXIST "$(BIN_FOLDER)" RMDIR /S /Q $(BIN_FOLDER)
	MKDIR $(BIN_FOLDER)
	MKDIR $(BIN_LIB_FOLDER)

build: prepare
	COPY LICENSE $(BIN_FOLDER)
	XCOPY /S $(LIB_FOLDER) $(BIN_LIB_FOLDER)
	go build -o $(BIN_FOLDER)\gwall.exe

run:
	go run main.go