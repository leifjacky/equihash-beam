package equihash

/*
int verifySolution(const char* input, const char* nonce, const char* output);
#cgo LDFLAGS: -L. -Llib -lbeampow
*/
import "C"
import "unsafe"

func VerifySolution(input, nonce, output []byte) bool {
	valid := int(C.verifySolution(
		(*C.char)(unsafe.Pointer(&input[0])),
		(*C.char)(unsafe.Pointer(&nonce[0])),
		(*C.char)(unsafe.Pointer(&output[0])),
	))
	if valid == 1 {
		return true
	}
	return false
}
