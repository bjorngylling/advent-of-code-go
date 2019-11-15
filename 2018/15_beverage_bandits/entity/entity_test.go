package entity

import "testing"

func TestEntity_Attack(t *testing.T) {
	a := New(GOBLIN, Pos(0, 0))
	b := New(ELF, Pos(0, 1))
	e := a.Attack(b)
	if e != nil {
		t.Error(e)
	}
	if a.HP != 200 {
		t.Errorf("Expected entity a to have 200 hp but was %d", a.HP)
	}
	if b.HP != 197 {
		t.Errorf("Expected entity b to have 197 hp but was %d", b.HP)
	}

	b.Pos = Pos(0, 2)
	e = a.Attack(b)
	if e == nil {
		t.Errorf("Expected error when attacking target out of reach, [%+v] attacking [%+v]", a, b)
	}
	if a.HP != 200 {
		t.Errorf("Expected entity a to have 200 hp but was %d", a.HP)
	}
	if b.HP != 197 {
		t.Errorf("Expected entity b to have 197 hp but was %d", b.HP)
	}
}
