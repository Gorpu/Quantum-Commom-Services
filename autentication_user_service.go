package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func AutenticationUser(user, encripted_pass string) string {
	key := "HELIO"
	keyLen := len(key)
	keyPos := -1
	offset := 0
	dest := ""
	offsetHex := encripted_pass[:2]
	offset64, _ := strconv.ParseInt(offsetHex, 16, 0)
	offset = int(offset64)

	srcPos := 2
	for srcPos < len(encripted_pass) {
		// pega cada byte (2 chars = 1 byte)
		srcAscHex := encripted_pass[srcPos : srcPos+2]
		srcAsc64, _ := strconv.ParseInt(srcAscHex, 16, 0)
		srcAsc := int(srcAsc64)

		if keyPos < keyLen-1 {
			keyPos++
		} else {
			keyPos = 0
		}

		tmpSrcAsc := srcAsc ^ int(key[keyPos])

		if tmpSrcAsc <= offset {
			tmpSrcAsc = 255 + tmpSrcAsc - offset
		} else {
			tmpSrcAsc = tmpSrcAsc - offset
		}

		dest += string(rune(tmpSrcAsc))

		offset = srcAsc
		srcPos += 2
	}

	return dest
}
func PasswordDay() string {
	// Obtém a data atual no formato YYYYMMDD
	dataStr := time.Now().Format("20060102")

	// Gera o hash MD5 da string da data
	hash := md5.Sum([]byte(dataStr))
	hashHex := hex.EncodeToString(hash[:])

	// Pega os primeiros 5 caracteres do hash
	parteHash := hashHex[:5]

	// Converte de hexadecimal para inteiro
	code, err := strconv.ParseInt(parteHash, 16, 64)
	if err != nil {
		code = 0
	}

	// Se for negativo (não deveria em Go), transforma em positivo
	if code < 0 {
		code = -code
	}

	// Retorna como string
	return fmt.Sprintf("%d", code)
}
