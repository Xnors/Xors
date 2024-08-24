# 开发文档

## 目录
```sh
./ 
├─.vscode
├─docs
├─src
│  ├─runBytecodes
│  ├─stack
│  ├─tokenizer
│  └─tokens2bytecodes
└─test
```

## 项目原理
- `tokenizer`: 将代码字符串转换为 [token](./tokens/tokens.md) 序列。
- `tokens2bytecodes`: 将 [token](./tokens/tokens.md) 序列转换为[字节码](./bytecodes/bytecodes.md) 序列
- `XVM(Xors Virtual Machine)`: [虚拟机](./XVM/XVM.md) , 执行字节码序列。

## 技术栈
- Go 语言 (1.22.5)

## 参考
- [Go 语言官方文档](https://go.dev/doc/)