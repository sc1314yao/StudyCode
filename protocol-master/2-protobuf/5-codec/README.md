
# 项目说明
该项目主要测试protobuf不同类型的变量占用字节长度

# 编译
在当前工程下先生成.proto文件对应的.cc和.h文件
```
protoc -I=./ --cpp_out=./  codec.proto
```

然后
```
mkdir build
cd build 
cmake ..
make
```
得到执行文件code_test