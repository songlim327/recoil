.PHONY: init

init:
	go run init/init.go

win-release:
	fyne package -os windows
	7z a -tzip "recoil-1.0.0_windows_amd64.zip" Recoil.exe

release: win-release