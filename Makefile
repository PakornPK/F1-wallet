TARGET_PATH = bin
GOARCH = GOARCH=amd64
GOMODULE = github.com/PakornPK/F1-wallet/cmd/f1-wallet
GOPROJECT = F1-wallet

buildWindows:
	env GOOS=windows $(GOARCH) go build -o ./$(TARGET_PATH)/windows/${GOPROJECT}.exe $(GOMODULE)

buildMacOS:
	env GOOS=darwin $(GOARCH) go build  -o ./$(TARGET_PATH)/macos/${GOPROJECT} $(GOMODULE)

buildLinux:
	env GOOS=linux $(GOARCH) go build -o ./$(TARGET_PATH)/linux/${GOPROJECT} $(GOMODULE)

buildARM:
	env GOOS=linux GOARCH=arm64 go build -o ./$(TARGET_PATH)/arm/${GOPROJECT} $(GOMODULE)

build: buildWindows buildMacOS buildLinux buildARM

clean:
	rm -rf bin

all: clean build