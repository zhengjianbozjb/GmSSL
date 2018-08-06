/* +build cgo */
package gmssl

/*
#include <openssl/err.h>
#include <openssl/bio.h>


extern long _BIO_get_mem_data(BIO *b, char **pp);
*/
import "C"

import (
	"errors"
)

func GetErrors() error {
	bio := C.BIO_new(C.BIO_s_mem())
	if bio == nil {
		return errors.New("GetErrors function failure 1")
	}
	defer C.BIO_free(bio)
	C.ERR_print_errors(bio)
	var p *C.char
	len := C._BIO_get_mem_data(bio, &p)
	if len <= 0 {
		return errors.New("GetErrors function failure 2")
	}
	return errors.New(C.GoString(p))
}
