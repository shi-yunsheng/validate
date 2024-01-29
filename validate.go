package validate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 对结构体的属性或方法进行验证
//
// data可以是结构体、结构体切片、结构体映射、结构体指针
//
// true 验证通过
//
// false 验证失败
func Struct(data interface{}) (bool, error) {
	var err error

	record := reflect.ValueOf(data)

	switch record.Kind() {
	// 切片
	case reflect.Slice:
		for i := 0; i < record.Len(); i++ {
			el := record.Index(i)
			p, e := Struct(el.Interface())
			if !p {
				msg := fmt.Sprintf("第%d项验证失败，失败原因：%v", i+1, e)
				return p, errors.New(msg)
			}
		}
	// 映射
	case reflect.Map:
		iter := record.MapRange()
		for iter.Next() {
			el := iter.Value()
			p, e := Struct(el.Interface())
			if !p {
				msg := fmt.Sprintf("键值%s验证失败，失败原因：%v", iter.Key().Interface(), e)
				return p, errors.New(msg)
			}
		}
	// 指针
	case reflect.Ptr:
		return Struct(record.Elem().Interface())
	// 结构体
	case reflect.Struct:
		stu := record.Type()

		for i := 0; i < stu.NumField(); i++ {
			el := stu.Field(i)
			value := record.Field(i)

			// 获取属性的验证规则
			temp, ok := el.Tag.Lookup("validate")
			if ok { // 存在验证规则
				rules := strings.Split(temp, ";")
				for _, r := range rules {
					k := r
					v := ""
					if strings.Contains(r, "=") {
						rs := strings.SplitN(r, "=", 2)
						k = rs[0]
						v = rs[1]
					}
					switch k {
					case "required":
						if len(v) == 0 || len(v) != 0 && computedRE(record, el.Name, v) {
							if value.IsZero() {
								msg := fmt.Sprintf("%s不能为空！", el.Name)
								err = errors.New(msg)
								return false, err
							}
						}
					case "email":
						if len(v) != 0 {
							msg := fmt.Sprintf("%s的邮件验证规则写法错误！", el.Name)
							panic(msg)
						}
						if !value.IsZero() {
							if !IsMail(value.String()) {
								msg := fmt.Sprintf("%s的值“%s”不是正确的邮件地址！", el.Name, value.String())
								err = errors.New(msg)
								return false, err
							}
						}
					case "phone":
						// 没有指定国家码，默认为中国大陆
						if len(v) == 0 {
							v = "+86"
						}
						if !value.IsZero() {
							if !IsPhone(v + value.String()) {
								msg := fmt.Sprintf("%s的值“%s”不是电话号码！", el.Name, value.String())
								err = errors.New(msg)
								return false, err
							}
						}
					case "password":
						// 没有指定密码复杂度，默认中等
						if len(v) == 0 {
							v = "m"
						}
						if v != "l" && v != "m" && v != "h" {
							msg := fmt.Sprintf(
								"%s的密码验证规则写法错误！密码复杂度只能是低（l）、中（m）、高（h）。",
								el.Name,
							)
							panic(msg)
						}
						if !value.IsZero() {
							if !passPasswordComplexity(value.String()) {
								msg := fmt.Sprintf("%s的值不满足密码复杂度检测！", el.Name)
								err = errors.New(msg)
								return false, err
							}
						}
					case "url":
						if len(v) != 0 {
							msg := fmt.Sprintf("%s的URL验证规则写法错误！", el.Name)
							panic(msg)
						}
						// 写URL验证
						if !value.IsZero() {
							if !IsURL(value.String()) {
								msg := fmt.Sprintf("%s的值“%s”不是一个有效的URL地址！", el.Name, value.String())
								err = errors.New(msg)
								return false, err
							}
						}
					case "length":
						if len(v) == 0 {
							msg := fmt.Sprintf("%s没有指定长度或长度范围！", el.Name)
							panic(msg)
						}
						if !value.IsZero() {
							if !isValidLength(value.String(), v) {
								msg := fmt.Sprintf("%s的值“%s”长度不符合预设长度！长度要求为：%s", el.Name, value.String(), v)
								err = errors.New(msg)
								return false, err
							}
						}
					case "range":
						if len(v) == 0 {
							msg := fmt.Sprintf("%s没有指定取值范围！", el.Name)
							panic(msg)
						}
						if !value.IsZero() {
							val := fmt.Sprintf("%v", value.Interface())
							if !isValidRange(val, v) {
								msg := fmt.Sprintf("%s的值“%v”不在预设范围之内！范围要求为：%s", el.Name, val, v)
								err = errors.New(msg)
								return false, err
							}
						}
					case "enum":
						if len(v) == 0 {
							msg := fmt.Sprintf("%s没有列出枚举值！", el.Name)
							panic(msg)
						}
						if !value.IsZero() {
							val := fmt.Sprintf("%v", value.Interface())
							if !existValue(v, val) {
								msg := fmt.Sprintf("%s的值“%v”不在所枚举之中！枚举值为：%s", el.Name, val, v)
								err = errors.New(msg)
								return false, err
							}
						}
					case "date", "datetime":
						if len(v) != 0 {
							msg := fmt.Sprintf("%s验证规则写法错误！", el.Name)
							panic(msg)
						}
						if !value.IsZero() {
							p := IsDatetime(value.String())
							if k == "date" {
								p = IsDate(value.String())
							}
							if !p {
								msg := fmt.Sprintf("%s的值“%s”不是一个正确的时间！", el.Name, value.String())
								err = errors.New(msg)
								return false, err
							}
						}
					case "prefix":
						if len(v) == 0 {
							msg := fmt.Sprintf("%s验证规则写法错误，必须指定一个前缀！", el.Name)
							panic(msg)
						}
						if value.Len() < len(v) || value.Len() >= len(v) && value.String()[:len(v)] != v {
							msg := fmt.Sprintf("%s的值“%s”没有包含预设前缀！要求包含前缀“%s”", el.Name, value.String(), v)
							err = errors.New(msg)
							return false, err
						}
					case "suffix":
						if len(v) == 0 {
							msg := fmt.Sprintf("%s验证规则写法错误，必须指定一个后缀！", el.Name)
							panic(msg)
						}
						if value.Len() < len(v) || value.Len() >= len(v) && value.String()[value.Len()-len(v):value.Len()] != v {
							msg := fmt.Sprintf("%s的值“%s”没有包含预设后缀！要求包含后缀“%s”", el.Name, value.String(), v)
							err = errors.New(msg)
							return false, err
						}
					}
				}
			}

			// 如果值是可递归的
			switch value.Kind() {
			case reflect.Slice, reflect.Map, reflect.Struct, reflect.Ptr:
				// 有值或者有规则
				p, e := Struct(value.Interface())
				if !p {
					msg := fmt.Sprintf("%s验证失败，失败原因：%v", el.Name, e)
					err = errors.New(msg)
					return p, err
				}
			}
		}
	default:
		panic("暂不支持验证该数据类型！")
	}
	return true, err
}
