default:
	curl -L -o include/hle-emulation.h https://raw.githubusercontent.com/andikleen/tsx-tools/master/include/hle-emulation.h
	mkdir -p build
	gcc -c -I include/ -o build/libgohle.so src/hle_locker.c

clean:
	rm -rf build
