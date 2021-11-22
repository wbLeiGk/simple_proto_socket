package utils

func Byte_2_int32(in_byte []byte) int32 {
	var result_int32 int32
	for i := 0; i < 4; i++ {
		result_int32 ^= int32(in_byte[i]&0xff) << (8 * (3 - i))
	}
	return result_int32
}

func Int32_2_byte(in_int32 int32) []byte {
	result_byte := make([]byte, 4)
	for i := 0; i < 4; i++ {
		uint8_t := in_int32 >> ((3 - i) * 8) & 0xff
		result_byte[i] = byte(uint8_t)
	}
	return result_byte
}
