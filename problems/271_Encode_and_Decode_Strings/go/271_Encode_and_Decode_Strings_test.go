package problems

import (
	"reflect"
	"testing"
)

func TestCodec_EncodeAndDecode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				strs: []string{"neet", "code", "love", "you"},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Codec{}
			encode := c.Encode(tt.args.strs)
			decode := c.Decode(encode)
			if !reflect.DeepEqual(decode, tt.args.strs) {
				t.Errorf("Encode() = %v, Decode() = %v", encode, decode)
			}
		})
	}
}
