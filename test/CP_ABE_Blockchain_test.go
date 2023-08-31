package test

import (
	"CP-ABE-Blockchain/core"
	"testing"
)

const AccessStructure = "((company_A AND repairman) OR (company_B AND consumer) OR admin)"

func TestDeviceCreateDTObject(t *testing.T) {
	core.CreateDTObject("td_sample.json", AccessStructure)
}

func TestDeviceEncryptDataUpload(t *testing.T) {
	core.EncryptDataUploadAndShareKey()
}

func TestUserDecryptData(t *testing.T) {
	core.DecryptData("urn:uuid:3deca264-4f90-4321-a5ea-f197e6a1c7cf", "device_data/data_enc.txt", "device_data/data_dec.txt")
}
