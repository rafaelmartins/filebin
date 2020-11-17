package magic

import (
	"errors"
	"strings"
	"sync"
	"unsafe"
)

// #cgo pkg-config: libmagic
// #include <magic.h>
import "C"

var (
	cookie C.magic_t
	m      sync.Mutex
)

func getError() error {
	if rv := C.magic_error(cookie); rv != nil {
		return errors.New("magic: " + C.GoString(rv))
	}

	return errors.New("magic: an error occurred") // *shrug*
}

func Init() error {
	m.Lock()
	defer m.Unlock()

	if cookie != nil {
		return errors.New("magic: already initialized")
	}

	cookie = C.magic_open(C.MAGIC_MIME_TYPE)
	if cookie == nil {
		return errors.New("magic: failed to open")
	}

	if rv := C.magic_load(cookie, nil); rv != 0 {
		err := getError()
		C.magic_close(cookie)
		return err
	}

	return nil
}

func Close() {
	m.Lock()
	defer m.Unlock()

	if cookie == nil {
		return
	}

	C.magic_close(cookie)
}

func Detect(data []byte) (string, error) {
	m.Lock()
	defer m.Unlock()

	if cookie == nil {
		return "", errors.New("magic: not initialized")
	}

	if rv := C.magic_buffer(cookie, unsafe.Pointer(&data[0]), C.size_t(len(data))); rv != nil {
		return strings.TrimSpace(C.GoString(rv)), nil
	}

	return "", getError()
}
