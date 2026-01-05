最新源码地址：hhttps://gitlab.0voice.com/2310_vip/protocol
1.  1-json
* cjson_test 			cjson测试
* jsoncpp_test  		jsoncpp测试
* rapidjson_test		rapidjson测试
1. 2-protobuf 
* 1-im				即时通讯
* 2-im_new			添加新字段对原有应用的影响
* 3-person          测试protobuf序列化性能，主要是和json、xml做对比
* 4-addressbook     protobuf自带的范例
* 5-codec           测试protobuf不同变量类型占用的字节数
1. 3-security  	数据加密，供参考，课上不做讲解
2. protobuf-cpp-3.19.6.tar.gz 下载好的protobu工具


```
# 解压
tar zxf protobuf-cpp-3.19.6.tar.gz
# 编译（make的时间有点长）
cd protobuf-3.19.6/
./configure 
make
sudo make install
sudo ldconfig
```
