package go_rs_c_abi_ex

/*
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/prebuilt/x86_64-apple-darwin/librs_c_abi_ex.a
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/prebuilt/aarch64-apple-darwin/librs_c_abi_ex.a
#cgo linux,amd64  LDFLAGS: ${SRCDIR}/prebuilt/x86_64-unknown-linux-gnu/librs_c_abi_ex.a
#cgo linux,arm64  LDFLAGS: ${SRCDIR}/prebuilt/aarch64-unknown-linux-gnu/librs_c_abi_ex.a

#include <stdint.h>
#include <stdbool.h>

typedef struct {
    uint32_t id;
    int32_t value;
} Item;

bool select_one_ffi(
    const Item* items,
    size_t len,
    Item* out
);
*/
import "C"
import "unsafe"

type Item struct {
	ID    uint32
	Value int32
}

func SelectOne(items []Item) (Item, bool) {
	if len(items) == 0 {
		return Item{}, false
	}

	cItems := (*C.Item)(unsafe.Pointer(&items[0]))
	var out C.Item

	ok := C.select_one_ffi(
		cItems,
		C.size_t(len(items)),
		&out,
	)

	if !ok {
		return Item{}, false
	}

	return Item{
		ID:    uint32(out.id),
		Value: int32(out.value),
	}, true
}
