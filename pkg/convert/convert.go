package convert

import "strconv"

type StrTo string

// String 方法实现了Stringer接口，作用是将StrTo类型转换为字符串类型
func (s StrTo) String() string {
	return string(s)
}

// Int 方法将StrTo类型转换为int类型
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt 方法将StrTo类型转换为int类型，如果转换失败则返回0
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// Uint32 方法将StrTo类型转换为uint32类型
func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUint32 方法将StrTo类型转换为uint32类型，如果转换失败则返回0
func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
