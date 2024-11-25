package arrays_and_hashing

import (
	"reflect"
	"testing"
)

func TestEncoder(t *testing.T) {
	tests := []struct {
		name   string
		strsIn []string
	}{
		{
			name:   "positive test",
			strsIn: []string{"neet", "code", "love", "you"},
		},
		{
			name:   "positive test2",
			strsIn: []string{" ", "code", "", "you"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{}
			encoded := e.Encode(tt.strsIn)
			decoded := e.Decode(encoded)

			if !reflect.DeepEqual(tt.strsIn, decoded) {
				t.Errorf("Decode() = %v, want %v", decoded, tt.strsIn)
			}
		})
	}
}
