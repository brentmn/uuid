package uuid

import (
	"testing"
)

func TestScan(t *testing.T) {
	token := new(UUID)
	test := "2a44643dc6a04b1ab23968fcf3d4f2e1"
	err := token.Scan(test)
	if nil != err {
		t.Logf("expected nil err, received %s", err)
		t.Fail()
	}
	if token.String() != test {
		t.Logf("expected %s, received %s", test, token.String())
		t.Fail()
	}
	b := []byte("2a44643dc6a04b1ab23968fcf3d4f2e1")
	err = token.Scan(b)
	if nil != err {
		t.Logf("expected nil err, received %s", err)
		t.Fail()
	}
	if token.String() != test {
		t.Logf("expected %s, received %s", test, token.String())
		t.Fail()
	}
	test = "not valid"
	err = token.Scan(test)
	if nil == err {
		t.Log("expected err, received nil")
		t.Fail()
	}
	itest := 45
	err = token.Scan(itest)
	if nil == err {
		t.Log("expected err, received nil")
		t.Fail()
	}
	ierr := "can only scan uuid into string or []byte, int was provided"
	if ierr != err.Error() {
		t.Logf("expected '%s', received '%s'", ierr, err)
		t.Fail()
	}
}
