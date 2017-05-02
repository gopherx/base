package statemachine

import (
	"testing"
)

const (
	still = iota
	eat
	move
	shit
)



func TestStrict(t *testing.T) {
	sm := New("eat->move->shit", true,
		Transition(still, eat),
		Transition(eat, move),
		Transition(move, shit),
	)

	if err := sm.Check(still, still); err == nil {
		t.Fatal("self state change not allowed for strict statemachine", err)
	}
}

func TestCanChange(t *testing.T) {
	sm := New("eat->move->shit", false,
		Transition(still, eat),
		Transition(eat, move),
		Transition(move, shit),
	)

	tests := []struct {
		cur int32
		valid []int32
	} {
		{
			still,
			[]int32{eat},
		},
		{
			eat,
			[]int32{move},
		},
		{
			move,
			[]int32{shit},
		},
	}

	for _, c := range tests {
		checked := map[int32]bool{
			still: false, eat: false, move: false, shit: false,
		}

		if err := sm.Check(c.cur, c.cur); err != nil {
			t.Fatalf("(%d -> %d) failed; err:%+v", c.cur, c.cur, err)
		}
		checked[c.cur] = true

		for _, v := range c.valid {
			if err := sm.Check(c.cur, v); err != nil {
				t.Fatalf("(%d -> %d) failed; err:%+v", c.cur, v, err)
			}
			checked[v] = true
		}

		for s, marked := range checked {
			if marked {
				continue
			}

			if err := sm.Check(c.cur, s); err == nil {
				t.Fatalf("(%d -> %d) accpted but not configured; err:%+v", c.cur, s, err)
			}
		}
	}
}