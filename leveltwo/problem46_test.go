package leveltwo

import "testing"

func TestGoldbachCon(t *testing.T) {
	e := 5777
	v := FindGoldbachConjectureFail()

	if e != v {
		t.Errorf("Expceted composite %d but got %d", e, v)
	}

}
