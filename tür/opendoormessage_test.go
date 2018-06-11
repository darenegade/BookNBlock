package tür

import (
	"testing"
)

func DecryptTest(t *testing.T) {
	msg := OpenDoorMessageHyperledger{
		DoorID:  "10",
		Payload: []byte("hkzkZMzbwPf34qipra19wm/p0tI1owu2Tt9YrtOEYniAiziVJEtFLL98VJh7xzeu36cpSWcJLoSLqzwZFi0S8BiX1yuNcb6oAR8pGzoft3dFTKxkMLYJPmhsF9DiK7RglRujw6yUQvnElntwC9wsSv4fJUievLja82xryjrSr/YVNInoR2cM53d8BqnwtOBeKi4oezpu1bwGXW4seNdX7zC7wvDuOocIyZ8E4c4s85srdGad2tp/dsF8oWKzOZN3gwxrwfIUtvivDRokdEu0dQqyCJG2aww6kHZF8Y5FuJWrzp3O/ZZLoYUcwrcto+CPwNqv2WAb0MK2/e/RTzgb4VxsSyJAoVQDoGj9EQpP+nn7/KPy9Nej3kyfJWf1uQuKhp2VVgb0+axfhIF3xcQIYC4oqHDvCvrQxfu3yRSrVMsXD0qkUQauaec9MCntXuXNiM6llq4TpzrVArY/Tc7UO12zz89hJDwcOENbygXFkdd2SKzFJ63kIReucev8+o1lYCWvcds7eedHc606W8CCXPQjHXLqLmXOlhSv0LwjR7NWKdEhuwB7zprfOXMSJDMo84IZkzqcO0xCAt6LEqoXV1RsOjaKEWFmc1hyMPB1nLPnfHSneJV9Wdfzg3mNEwYT5W77T8pJcOcf7+7wRRHay9IXPnhvo7NKmd415sPCsic54n6qumXXZHjOQUiPtD0RfO15oK9Ksg9IEowLxGFmWvrAvMkDrqqMAV1pWETtlkjSPyoPws1GMtf7KWnWqfVg"),
	}

	renterPubkey := "-----BEGIN PUBLIC KEY-----\nMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMrNw+Z0L3AIJYgZIdB+4SzWS3tzkxPt\nACoamUO8BkSOiLe45BC6vbKDau67N0ZzyQQnGAE0UV0K/3zasmtl98kCAwEAAQ==\n-----END PUBLIC KEY-----"
	timestamp := 1528629676970

	key := `-----BEGIN PRIVATE KEY-----
	MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEAi3G84XBu32FqDfn9
	unJDRI37dAVrMmVmtoRxT+ebf9M+eY0tClhqReHg4uDxTfGhHn5OXnqGNcAVsdzY
	wJIV2QIDAQABAkAMNgKPT1Q2NYafALRKXnUrjK9nYo4XlK+g7goqMCL613ijAHlN
	QYSHiYZafbzZYRKK+9dD77IGQjv42wXAyF+xAiEAy6ixO/8pfsVATga57z2+HliR
	0+9jP52JKeMOaM0RBMUCIQCvSC2ahJyeSQDyzx9/t1jwjHPk68T6OrThlrxMuGTm
	BQIgaraFx843y/lHbJsRqk5L5FK8drSk6Jx/VrdmwXtSQ5ECIDuNMEFBMNzuPK5C
	BJeluUfw0CdEmyXQ8Ed8qPj/5PfpAiB8Wz+H6ETCHUHw6M1pzAfRBIJ12RD1HqBr
	U/2JlROdoQ==
	-----END PRIVATE KEY-----
	-----BEGIN PUBLIC KEY-----
	MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAItxvOFwbt9hag35/bpyQ0SN+3QFazJl
	ZraEcU/nm3/TPnmNLQpYakXh4OLg8U3xoR5+Tl56hjXAFbHc2MCSFdkCAwEAAQ==
	-----END PUBLIC KEY-----`

	msg.Decrypt(key)

	if msg.DoorID != "10" {
		t.Error("DoorID hat sich geändert, ist nicht 10")
	}

	if msg.RenterPubkey != renterPubkey {
		t.Error("Renterpubkey stimmt nicht")
	}

	if msg.Timestamp != timestamp {
		t.Error("Timestamp ist nicht 1528629676970")
	}
}
