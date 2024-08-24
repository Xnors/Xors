# Tokens

## 类型
### 变量
- 特征: 第一个字符是 `$`
- 示例: `$name`, `$age`

### 参数
- 特征: 末尾字符是 `:`
- 示例: `name:`, `age:`

### 数字
- 特征: 数字字符
- 示例: `123`, `456.789`

### 字符串
- 特征: 单双引号包裹的字符序列
- 示例: `"hello world"`, `"I'm a string"`

### 运算符
- 特征: 运算符字符
- 示例: `+`, `-`, `*`, `/`, `%`

### 分隔符
- 示例: 分号 `;` 
  
### 隐分隔符
- 特征: 不会出现在 [tokens](#tokens) 中
- 示例: `,`

### 词
- 特征: 单词或短语
- 示例: `true`, `false`, `null`

## 解析示例
### 原始代码 ( `string` )
```kotlin
use strings;

out "第一句";
out 114+514;
out strings::merge("hello" , "world");

fun int_add | a int b int | int{
    ret a+b;
}

out int_add(10, 20);
```
### 解析结果 ( `[]string` ):
```kotlin
use
strings
;
out
"第一句"
;
out
114+514
;
out
strings::
merge
(
"hello"
"world"
)
;
fun
int_add
|
a
int
b
int
|
int
{
ret
a+b
;
}
out
int_add
(
10
20
)
;
```