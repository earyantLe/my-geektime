package m3u8

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// aes128Decrypt decrypts data using AES-128-CBC
func aes128Decrypt(data, key, iv []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, fmt.Errorf("invalid key length: %d, expected 16", len(key))
	}

	// If IV is empty or not 16 bytes, use zero IV
	if len(iv) != 16 {
		iv = make([]byte, 16)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("create cipher failed: %w", err)
	}

	// CBC mode always works in whole blocks
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("data is not a multiple of block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	return data, nil
}

// decryptAliyunVOD decrypts data using Aliyun VOD encryption (AES-128-ECB)
// Aliyun VOD uses AES-128 in ECB mode for TS segment encryption
func decryptAliyunVOD(data []byte, key string) []byte {
	if len(key) != 16 {
		// If key length is invalid, return original data
		return data
	}

	keyBytes := []byte(key)

	// Create AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		// If cipher creation fails, return original data
		return data
	}

	// ECB mode works on fixed-size blocks (16 bytes for AES)
	blockSize := block.BlockSize()
	decrypted := make([]byte, len(data))

	// Process each block independently (ECB mode)
	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		if be > len(data) {
			break
		}
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

// ensureTSSyncByte ensures TS data starts with sync byte 0x47
// Some TS files do not start with SyncByte 0x47, they cannot be played after merging
// Need to remove the bytes before the SyncByte 0x47 (71)
func ensureTSSyncByte(data []byte) []byte {
	syncByte := uint8(0x47) // 71

	for i := 0; i < len(data); i++ {
		if data[i] == syncByte {
			return data[i:]
		}
	}

	// If no sync byte found, return original data
	return data
}
