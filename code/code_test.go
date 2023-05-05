package code

import "testing"

func TestMake(t *testing.T) {

	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{Opconstast, []int{65534}, []byte{byte(Opconstast), 255, 253}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)
		if len(instruction) != len(tt.expected) {
			t.Errorf("instructions has wrong length. want = %d, got=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d", i, b, instruction[i])
			}
		}

	}

}
