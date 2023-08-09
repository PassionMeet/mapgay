package password

import "testing"

func TestHashPassword(t *testing.T) {
	pwd := "chenm123456"
	hashRes, err := HashPassword(pwd)
	if err != nil {
		t.Fatalf("hash %s", err)
	}
	t.Logf("hash %d %s", len(hashRes), hashRes)
}

func TestCheckPasswordHash(t *testing.T) {
	pwd := "chenm123456"
	hashRes, err := HashPassword(pwd)
	if err != nil {
		t.Fatalf("hash %s", err)
	}
	t.Logf("hash %d %s", len(hashRes), hashRes)
	ok := CheckPasswordHash(pwd, "$2a$14$dHy6ns/u8u3Lj1Jx.Zd5JehpQej8nZ7caOGKD4wTTJiXu3yslgRB6")
	t.Log(ok)
}
