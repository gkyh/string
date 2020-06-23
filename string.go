package String

import (
	"fmt"
	"regexp"
	"strconv"
)

func Exist(f string) bool {
	return string(f) != string(0x1E)
}

func Uint8(f string) uint8 {
	v, _ := strconv.ParseUint(f, 10, 8)
	return uint8(v)
}

func Int(f string) int {
	v, _ := strconv.ParseInt(f, 10, 0)
	return int(v)
}
func Int32(f string) int32 {
	v, _ := strconv.ParseInt(f, 10, 64)
	return int32(v)
}

func Int64(f string) int64 {
	v, _ := strconv.ParseInt(f, 10, 64)
	return int64(v)
}

func Float64(f string) float64 {
	v, _ := strconv.ParseFloat(f, 64)
	return float64(v)
}

// "12,123.45"->1212345
func Long(str string) int64 {

	re, _ := regexp.Compile(`[^\d^\.^\-]`)
	src := re.ReplaceAllString(str, "")

	f64, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return 0
	}

	s := fmt.Sprintf("%0.0f", f64*100)
	i64, _ := strconv.ParseInt(s, 10, 64)
	return int64(i64)

	/*
		strs := strings.Split(src, ".")

		s64 := strs[0]

		if len(strs) == 1 {
			s64 += "00"
		} else {

			ys := strs[1]

			if len(ys) == 2 {

				s64 += ys
			} else if len(ys) == 1 {

				s64 += ys
				s64 += "0"
			} else if len(ys) == 0 {

				s64 += "00"
			} else {

				s64 += string(ys[0:2])
			}
		}

		i64, _ := strconv.ParseInt(s64, 10, 64)
	*/
}

func Ncy(i int64) (s string) {

	if i == 0 {
		return "0.00"
	}
	var f float64 = float64(i)
	f = f / 100
	s = fmt.Sprintf("%0.2f", f)
	return
}

// Convert any type to string.
func Parse(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

/*
// HexStr2int converts hex format string to decimal number.
func HexStr2int(hexStr string) (int, error) {
	num := 0
	length := len(hexStr)
	for i := 0; i < length; i++ {
		char := hexStr[length-i-1]
		factor := -1

		switch {
		case char >= '0' && char <= '9':
			factor = int(char) - '0'
		case char >= 'a' && char <= 'f':
			factor = int(char) - 'a' + 10
		default:
			return -1, fmt.Errorf("invalid hex: %s", string(char))
		}

		num += factor * PowInt(16, i)
	}
	return num, nil
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 16

		c := "?"
		if r >= 0 && r <= 9 {
			c = string(r + '0')
		} else {
			c = string(r + 'a' - 10)
		}
		hex = c + hex
		num = num / 16
	}
	return hex
}
*/
