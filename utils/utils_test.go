package utils

import (
	"reflect"
	"testing"
)

func TestByte_2_int32(t *testing.T) {
	type args struct {
		in_byte []byte
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "b_2_i",
			args: args{in_byte: []byte{0x12, 0x34, 0x56, 0x78}},
			want: 0x12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Byte_2_int32(tt.args.in_byte); got != tt.want {
				t.Errorf("Byte_2_int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32_2_byte(t *testing.T) {
	type args struct {
		in_int32 int32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "i_2_b",
			args: args{in_int32: 0x12345678},
			want: []byte{0x12, 0x34, 0x56, 0x78},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32_2_byte(tt.args.in_int32); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32_2_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}
