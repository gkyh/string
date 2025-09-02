package stringx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
    // 去掉空格和千分位逗号
    str = strings.ReplaceAll(str, ",", "")
    str = strings.TrimSpace(str)

    // 只保留数字、小数点、负号
    re := regexp.MustCompile(`[^\d.\-]`)
    src := re.ReplaceAllString(str, "")
    if src == "" {
        return 0
    }

    parts := strings.SplitN(src, ".", 2)

    // 整数部分
    intPart, _ := strconv.ParseInt(parts[0], 10, 64)
    result := intPart * 100

    // 小数部分
    if len(parts) == 2 {
        frac := parts[1]
        if len(frac) > 2 {
            frac = frac[:2] // 截断两位（如果想四舍五入，这里改逻辑）
        }
        for len(frac) < 2 {
            frac += "0" // 补足两位
        }
        fracInt, _ := strconv.ParseInt(frac, 10, 64)
        if intPart < 0 {
            result -= fracInt // 负数情况
        } else {
            result += fracInt
        }
    }

    return result
}

func Ncy(i int64) (s string) {

	yuan := i / 100
    jiaoFen := i % 100
    return fmt.Sprintf("%d.%02d", yuan, jiaoFen)
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
