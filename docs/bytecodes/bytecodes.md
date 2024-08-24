# *Bytecodes*
字节码文档

## 字节码
> 注意: [源代码](../../src/tokens2bytecodes/types.go)中定义字节码时名字前有 `OP_`, 例如 `OP_LOAD_CONST` 等.
> 
> 注意: 很多文档没有完善 <!-- TODO: 完善文档 -->

| 字节码        | 描述                                                                                                                                                                                                 | 详细信息                                                        |
| ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------- |
| IMPORT_MODULE | 导入模块到 [Modules](../运行时数据/Modules.md)                                                                                                                                                       |                                                                 |
| LOAD_MODULE   | 加载 [Modules](../运行时数据/Modules.md) 中的模块 到 [ModulesStack](../运行时数据/ModulesStack.md)  中                                                                                               |                                                                 |
| LOAD_CONST    | 加载[常量](../types/data-types.md#常量)到[MainStack](../运行时数据/MainStack.md)顶                                                                                                                   |                                                                 |
| LOAD_NAME     | 加载[名称](../运行时数据/NAME.md)到[MainStack](../运行时数据/MainStack.md)顶                                                                                                                         |                                                                 |
| STORE_NAME    | 存储[MainStack](../运行时数据/MainStack.md)顶的值到[名称](../运行时数据/名称.md)                                                                                                                     |                                                                 |
| LOAD_FAST     | 加载[FAST](../运行时数据/FAST.md)到[MainStack](../运行时数据/MainStack.md)顶                                                                                                                         |                                                                 |
| STORE_FAST    | 存储[MainStack](../运行时数据/MainStack.md)顶的值到[FAST](../运行时数据/FAST.md)                                                                                                                     |                                                                 |
| CALL_FUNCTION | 调用函数                                                                                                                                                                                             | 见[函数调用详解](../技术逻辑/函数调用.md#字节码-call_function-) |
| RETURN        | 将当前[FrameStack](../运行时数据/FrameStack.md)顶的值返回到上一级[Frame](../运行时数据/Frame.md)的 [FrameStack](../运行时数据/FrameStack.md)顶, 并清空当前 [FrameStack](../运行时数据/FrameStack.md) |                                                                 |
| PTOP          | 弹出当前[FrameStack](../运行时数据/FrameStack.md)顶的值                                                                                                                                              |                                                                 |
