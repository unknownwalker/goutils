package utils

import (
	"errors"
)

func BlockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("源缓冲区 索引超出范围")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("目标缓冲区 索引超出范围")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}

func BlockReplace(src []byte) (bool, error, []byte) {
	srcLen := len(src)
	tmp := make([]byte, srcLen)

	for i := 0; i < srcLen; i++ {
		if (i + 3) < srcLen {
			if src[i] == 0x00 && src[i+1] == 0x00 && src[i+2] == 0x01 && src[i+3] == 0xe0 && srcLen >= 9 {
				a := int(src[i+8])

				tmp = make([]byte, srcLen-8-a-1)
				copy(tmp, src[0:i])
				BlockCopy(src, i+8+a+1, tmp, i, srcLen-i-8-a-1)
				src = tmp
				srcLen = len(tmp)
				i = i - 1
			}
		}
	}
	return true, nil, tmp
}

func BABlockReplace(src []byte) (bool, error, []byte) {
	srcLen := len(src)
	tmp := make([]byte, srcLen)

	for i := 0; i < srcLen; i++ {
		if (i + 13) < srcLen {
			if src[i] == 0x00 && src[i+1] == 0x00 && src[i+2] == 0x01 && src[i+3] == 0xba && srcLen >= 9 {
				Packstuff := int(src[i+13])
				a := int(Packstuff & 0x7)
				tmp = make([]byte, srcLen-13-a-1)
				copy(tmp, src[0:i])
				BlockCopy(src, i+13+a+1, tmp, i, srcLen-i-13-a-1)
				src = tmp
				srcLen = len(tmp)
				i = i - 1
			}
		}
	}
	return true, nil, tmp
}
