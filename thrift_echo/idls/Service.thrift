include "User.thrift"

namespace go T_go_RPC.models.rpc  //格式：namespace 语言名 路径，注意末尾不能有分号

typedef map<string, string> Data  // 类型定义

struct Response {
    1:required i32 errCode; //错误码
    2:required string errMsg; //错误信息
    3:required Data data;
}

//定义服务
service Greeter {
    Response SayHello(
        1:required User.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}
