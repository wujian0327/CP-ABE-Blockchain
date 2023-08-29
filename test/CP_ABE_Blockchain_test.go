package test

import (
	"CP-ABE-Blockchain"
	"testing"
)

func TestDeviceCreateDTObject(t *testing.T) {
	main.CreateDTObject()
}

func TestDeviceEncryptDataUpload(t *testing.T) {
	main.EncryptDataUpload()
}

func TestUserDecryptData(t *testing.T) {
	main.DecryptData("urn:uuid:0804d572-cce8-422a-bb7c-4412fcd56f06", "device_data/data_enc.txt", "device_data/data_dec.txt")
}
