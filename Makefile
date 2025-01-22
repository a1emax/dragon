.PHONY: install-ebitenmobile
install-ebitenmobile:
	go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@v2.8.6

.PHONY: build-windows
build-windows:
	mkdir -p .local
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
		-o .local/dragon-windows.exe \
		dragon/app/desktop

.PHONY: build-android
build-android:
	mkdir -p .local
	ebitenmobile bind \
		-target android \
		-androidapi 24 \
		-javapkg com.github.a1emax.dragon.go \
		-o .local/dragon-android-intern.aar \
		dragon/app/android_intern
	cp .local/dragon-android-intern.aar app/android/intern/default.aar
