package directdebit_test

import (
	"testing"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func TestGenerateRSASignature(t *testing.T) {
	timestamp := "2023-03-28T22:13:27+07:00"
	clientID := "fceda4da-ae95-4d14-8cb6-eef392139a1b"
	privkey := `-----BEGIN PRIVATE KEY-----
MIIG/AIBADANBgkqhkiG9w0BAQEFAASCBuYwggbiAgEAAoIBgQCrb37xqR+7N0aJ
JY7juytGVC8mlQYFI3hcXZJVhnMlLsL0y2F/OBxduMXNlM30dBA9dYr4ggKdSYMu
toOMsGCgYavYemmIW0YlpTz2up+O50khPhdIHsA1R0B1v9QCDkYEYpE9RL9bDTP+
zJb/uq4QA6f/EQeuSwNUaV1s/6uBmrtk5vXSJi/AsZ5trAknVb3Onv3GXIbhaIAB
Y/e/f+19Z9y48UOJL0kSqAGhSM7Pd6AakuJ+aV2hCwhB3zbsI55bPOORRbqosBNU
wkW+xid837u+vc2SCm5FMvVrvkIwMicqhDKsQlvHpG1+VSv5S92HUwGpC5U1+XYO
Y7yHaX/KLcq1lYAttjrwRWs5EVHao8hYjjtqrmQicRBSfQHDmY2nklvylHveTzKp
j83oQxZYRyy+lqD9IjSHBRVsDcaMDfDMqzrUCWqavMxQVqnSDLgHt/CW8osuJ0UM
2o01hZB/srEE0BXCW5HoZlaW7udd8Ti7ZAs36+PP/LMs+rus0bcCAwEAAQKCAYAR
ByaPP1KxCEj/x1S9hvJB7ouuY9/ws7i5R+wIha27ND1WDjt1ZO/gWUGAbXbVgI+6
Ywn2LAexcsNOaP+BAmXemET24BXKXvKFO7fl89x0V8G6RQ4P8kn6IMUkzPR0bdGD
jvzJHqJ5G0MeXFjlNrgiTBKsMZdXNwkyIbMPaAezfFh/qbch8/wLQjkvwIY6O3h6
ZO1k/fzBt9z7BmBty3md2qqgTgp8vk8eRMTArdgo4ENtUEih8LpFjDB6Rn8Qjmr7
GkPKeAENCfLoZbIkzFguROIdnBHDWZghBxxo0MwNStOsi/DX+kJfrTgz4bzS+UEG
6Ylf64z7E68mmVfkC7M3VzvFgRqPK1oG+iu/4U7etz+oErKX4cqJnlKJgItXK46D
pmsHcJowcsgziugaity4DOj+gFie5POtN1ux9pI7Do0ZAD6HHpei53l8RpQGgBtl
+iAlPx2eQpW/uV1bapreWbYbNajT3jxkf7LtYsIsGsr17ti99pNMh6HMO4tSjckC
gcEA1WxDMo1KlydmEz7smlUFCH3rWUPdm8F8mCRzo0Q7F4QFoXe0rDLZZlt6puFn
Bez8B8bEtDmB5bYjEZhdPyhbl+HQ/4XlUo88GJUfWwpIiwWecxalFwJE5K3bLP1C
Yzgx8BarGsr4/miY1JD/Kjw8wVEgPGG6oZeBdS4Z9uqlUuiLMsIo8oSXPdkfI9Gg
Sxn2GcBGLsM0IPcdfEgXJw64VF0kGAu//Vkpv4xc5LSqKNeR28zNEUF2xyTLgYIM
SHepAoHBAM2i5KcDBVpE+iqKvXdTB85XJCnZXqG8uVXHwonC7P/UiQeZWL2xJlFV
ibxguhaLsbCG+gXzVJdwSSoKX+p0pk8G+Lf5JYhQLtQQEZUc2mSOC2e5bxhK/Lcw
yVuSFIftg0+fc6fNiYTXBX839r1wLWs+cshGMvAO0POimhw6TjVcDjbaAaTpN77Z
UW7644i2eRenPKFbNvww16jVHhIVdY7oSCAKAGivuyFvSKRghZMh/JP456OJXHQ6
lWIS019aXwKBwEX6vrnnrEqNx6GN42TjdcgICdB2OUbmFaWJZkVljP6z8mi0aJCC
B9jRLBFmHTLLNwSRv1Pc+2PH6g3N6N1ZrVbK2429aKk+gBULaIGgiJLVH9Ra230E
6HQXMaO50zfXaEByHl6lqSk6QMqKVLCTmdRFdo11+g0cMX2rxSW6YMUjrOjS0zxa
D4FfHR/Qj3+wnoppClow9XnNrWRf+v96iyRWegxMZgJ7Zv4A10DCoHzN2my45ZC/
52N7BCON8dsdKQKBwDPZyAfothfN3rqNYzrMP+KijGbU/YyQtrbPeNkdwn67i5XT
79Fc8sl9ZQ6P4TxAGxzk2/RWJ9VLpdco6IiIw0qX+m0BMJqPhU9JgfV0YgkK3Ata
cY3RkqlqbstdKTohBIQ2M4ZzSCKrySIL7XZU687n3y9qq/tl8QAN1wgZF5FS1e60
x8daWwkPaP4v2uGlCSGStLIG+vVaJ3bVzhBHQu422cDiZLoA3ZGPquRvxh6Uakix
cU8GGr7f6rzg/FVFxwKBwEkrok6ccZf3MVLnlNmOCAb7nP+ehquOVv75fRdPV+q5
B+PKt975LyM4bbmKw0AHz/jKoFTCPyDFgmA43i75VDjJ33ADALcLoPy333unC5I1
3SVZtXzV5nLeynEq9ZsvvcEvABokiqoe1Vo3evNrfRxb4meIg5spLEGsBRRoOaPO
6UtnlQQpNKvFyLx3mCwtYkcp/5KJs0WSPJakpoQRlsXjEzzmJesBxVn80ble2qDu
5U0gfFZB2kJ9FaBV+zZotA==
-----END PRIVATE KEY-----`

	expectedOutput := "6ff5530e0b1a90e4af1ba9fe284e9b3546ac40cfd222c77c9b46022761146430c588b05e7388fc90900083df2d36e558e93e79f54f5e34085b4214b21abc5d1548fffc8813b263a8848ccc0a16e5cb3083c7eea8a0f276154d9596aade158d0e8c207a49f214550a48c01e604d402a91f33235f5d1621c23ef3182d635054dd9f88da976f887f3e44761fd53039aad7c7781ecb2a2d6b6a064299b16c301e09d762b5fd8e1e15ed9433cea8f02a4979b466250c1f8b1172461d1a6735d3ac056512552c030451774ababf5b7c07c6e2241b417bd17cdbdb14e7655cc4f002e8c120c48307a83c63bcb8dd03c3b38d66591a9b872797c2b1de0e930b10e269317b0d6be6d55e919169cd264538c0955d48649bdd54808207088c2ee42a1550bf8c9f5ff3b0270605f2272305437d5bb1304c9ba93ea2fc9ee02d73086411d5e7586c3ba460143af995a551084ac66d3a9e826e0307a84db6766fe47cfb948010ae559790f2d4512ab899803c234f660df34bb475dddeb3b27145fdb80caa2e829"

	signature, err := directdebit.GenerateRSASignature(timestamp, privkey, clientID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if signature != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, signature)
	}
}

func TestGenerateHmacSignature(t *testing.T) {
	httpMethod := "POST"
	path := "/api/v2/registration-account-binding"
	accessToken := "ed8f7f42a6d741fbb891bae654e81678"
	jsonString := `{"partnerReferenceNo":"qic5taezfgta7yhn74u7cppr47wan027","authCode":"8ecx3bnfk9bu2o4y0z8ur8pfhH12eltF","merchantId":"AYOPOP"}`
	timestamp := "2023-03-28T22:13:27+07:00"
	clientSecret := "fceda4da-ae95-4d14-8cb6-eef392139a2c"

	expectedOutput := "2e955289baf7f297dda75b830c00f15b81a71710c2d0a0bbdf5884ae15bf47bb46e627a0b25cff1c1da16c42682ec69945950a1b120b6be490a53b7d613cf7e4"

	signature := directdebit.GenerateHmacSignature(httpMethod, path, accessToken, jsonString, timestamp, clientSecret)

	if signature != expectedOutput {
		t.Errorf("Expected %v, got %v", expectedOutput, signature)
	}
}

func TestGenerateRSASignature_InvalidPrivateKey(t *testing.T) {
	// Using an invalid private key
	privkey := `-----BEGIN RSA PRIVATE KEY-----
invalid
-----END RSA PRIVATE KEY-----`

	_, err := directdebit.GenerateRSASignature("2023-03-28T22:13:27+07:00", privkey, "some_client_id")
	if err == nil {
		t.Fatal("Expected an error from invalid private key, got nil")
	}
}

func TestGenerateRSASignature_InvalidPKCS8Key(t *testing.T) {
	timestamp := "2023-03-28T22:13:27+07:00"
	clientID := "fceda4da-ae95-4d14-8cb6-eef392139a1b"
	// This is a valid PEM block but not a valid RSA private key.
	privkey := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJ9wKHG/JDCr+6eZ6as5hHspDhZEUkveTF9GnCcDqRVmS/E4rhYw
-----END RSA PRIVATE KEY-----`

	_, err := directdebit.GenerateRSASignature(timestamp, privkey, clientID)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
	// Check if the error message contains what you expect, to ensure it's this error that's thrown.
}
