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
