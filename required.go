package validate

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// 解析Required表达式并返回表达式的结果
func computedRE(record reflect.Value, key string, expression string) bool {
	// 先切割或运算
	orE := strings.Split(expression, "|")

	orRes := false

	// 遍历或运算每一项
	for _, e := range orE {
		// 切割与运算
		andE := strings.Split(e, "&")
		// 与运算结果
		andRes := true
		// 遍历与运算每一项
		for _, a := range andE {
			symbol := []string{"==", "!=", "<=", ">=", "<", ">"}
			for _, sb := range symbol {
				if strings.Contains(a, sb) {
					temp := strings.Split(a, sb)

					re := regexp.MustCompile("'([^']*)'")
					fields := re.FindStringSubmatch(temp[0])
					// 没有拿到字段
					if len(fields) <= 1 {
						msg := fmt.Sprintf("属性%s的`required`表达式不正确！", key)
						panic(msg)
					}
					// 判断结构体是否存在属性
					if _, ok := record.Type().FieldByName(fields[1]); ok {
						r := compare(sb, temp[1], record.FieldByName(fields[1]))
						andRes = r
					} else {
						msg := fmt.Sprintf("结构体不存在%v的属性！", fields[1])
						panic(msg)
					}

					break
				}
			}
			// 如果与运算有一项结果为假，直接执行或运算的另一项
			if !andRes {
				break
			}
		}
		// 如果或运算有一项结果为真，直接返回真
		if andRes {
			return andRes
		}
	}

	return orRes
}

// 比较运算符两侧
func compare(symbol string, a string, b reflect.Value) bool {
	_type := b.Type()

	switch symbol {
	case "==":
		switch _type.Kind() {
		// 布尔值
		case reflect.Bool:
			return a == "true" && b.Interface() == true || a == "false" && b.Interface() == false
		// 数值类
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() == numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() == numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() == numA
			} else {
				panic(ok)
			}
		// 字符串类
		case reflect.String:
			// none为值的情况，判断指定属性是否为空并返回结果
			if a == "none" {
				return b.IsZero()
			}
			// 否则就是比较字符串
			return b.String() == a
		// 其他
		default:
			// none为值的情况，判断指定属性是否为空，并返回
			if a == "none" && b.IsZero() {
				return true
			}
		}
	case "!=":
		switch _type.Kind() {
		case reflect.Bool:
			return a == "true" && b.Interface() != true || a == "false" && b.Interface() != false
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() != numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() != numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() != numA
			} else {
				panic(ok)
			}
		case reflect.String:
			if a == "none" {
				return !b.IsZero()
			}
			return b.String() != a
		default:
			if a == "none" && !b.IsZero() {
				return true
			}
		}
	case "<=":
		switch _type.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() <= numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() <= numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() <= numA
			} else {
				panic(ok)
			}
		default:
			panic("只能数值类可以进行比较！")
		}
	case ">=":
		switch _type.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() >= numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() >= numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() >= numA
			} else {
				panic(ok)
			}
		default:
			panic("只能数值类可以进行比较！")
		}

	case "<":
		switch _type.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() < numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() < numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() < numA
			} else {
				panic(ok)
			}
		default:
			panic("只能数值类可以进行比较！")
		}
	case ">":
		switch _type.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if numA, ok := strconv.ParseUint(a, 0, 0); ok == nil {
				return b.Uint() > numA
			} else {
				panic(ok)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if numA, ok := strconv.ParseInt(a, 0, 0); ok == nil {
				return b.Int() > numA
			} else {
				panic(ok)
			}
		case reflect.Float32, reflect.Float64:
			if numA, ok := strconv.ParseFloat(a, 64); ok == nil {
				return b.Float() > numA
			} else {
				panic(ok)
			}
		default:
			panic("只能数值类可以进行比较！")
		}
	}
	return false
}
