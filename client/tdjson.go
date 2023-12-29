//go:build darwin
// +build darwin

package client

/*
#cgo darwin CFLAGS: -I/usr/local/include
#cgo darwin LDFLAGS: -L/usr/local/lib -L/opt/homebrew/Cellar/openssl@3/3.2.0_1/lib/ -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
*/
import "C"
