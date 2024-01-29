# Validate

验证结构体的属性或方法，只需在需要验证的属性或方法后面添加`validate`标签即可。例如：

```go
type User struct {
    ID uint `validate:"required;length=9"`
    Username string `validate:"required;length=[5, 10]"`
    Password string `validate:"required;password=h"`
    Nickname string
    Age uint `validate:"range=[15,]"`
    Sex string `validate:"enum=(男,女)"`
    RegTime string `validate:"datetime"`
    Email string `validate:"email;required='IsAdmin'==true"`
    IsAdmin bool
}
```

> **释义：** 
>
> `ID` 必填；长度为9
>
> `Username` 必填；长度最短为5，最长为10
>
> `Password` 必填；密码复杂度为高
>
> `Age` 最小值为15
>
> `Sex` 只能是`男,女`的其中一个
>
> `RegTime` 值是日期时间格式
>
> `Email` 值是邮件格式；当属性`IsaAdmin`为`true`时，该属性为必填

`validate`支持以下验证：

| 字段名     |          值          | 描述                                                         | 示例                                                   |
| ---------- | :------------------: | ------------------------------------------------------------ | ------------------------------------------------------ |
| `required` | `布尔值或布尔表达式` | 是否必填，如果有`required`但没有指定值，默认为`true`         | `validate:"required"`                                  |
| `email`    |                      | 电子邮件格式                                                 | `validate:"email"`                                     |
| `datetime` |                      | 日期时间格式                                                 | `validate:"datetime"`                                  |
| `date`     |                      | 日期格式                                                     | `validate:"date"`                                      |
| `phone`    |      `国际区号`      | 合法电话号码，国际区号默认是+86                              | `validate:"phone"`<br />`validate:"phone=+852"`        |
| `length`   |  `uint或区间表达式`  | 长度或最大长度、最小长度，可以只限制最小或最大<br />仅适用于`len()`能计算的类型 | `validate:"length=15"`<br />`validate:"length=[5,15]"` |
| `enum`     |   `(值1,值2,...)`    | 只能是枚举的值                                               | `validate:"enum=(男,女)"`                              |
| `password` |     `h | m | l`      | 属性的值是否符合对应的密码复杂度，默认为`m`<br />`h` 高复杂度，必须包含大小写字母、数字以及特殊符号，且长度最少8位<br />`m` 标准复杂度，包含大小写字母、数字、特殊符号，长度最少6位<br />`l` 低复杂度，几乎没有要求 | `validate:"password"`<br />`validate:"password=h"`     |
| `url`      |                      | URL格式的字符串                                              | `validate:"url"`                                       |
| `range`    |     `区间表达式`     | 取值范围，可以只限制最小或最大；数值类                       | `validate:"range=[3.24, 8.74]"`                        |
| `prefix`   |       `string`       | 规定字符串以什么开始                                         | `validate:"prefix=/"`                                  |
| `suffix`   |       `string`       | 规定字符串以什么结束                                         | `validate:"suffix=."`                                  |

