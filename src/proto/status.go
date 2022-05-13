package proto

const (
	StatusCode_Success                  int32 = 0  	// 成功返回
	StatusCode_JsonMarshalError   		int32 = 1 	// JSON序列化失败
	StatusCode_JsonUnmarshalError 		int32 = 2 	// JSON反序列化失败
	StatusCode_EncodeError 				int32 = 3 	// Encode失败
	StatusCode_HttpError				int32 = 4   // http 错误
	StatusCode_ConsistentEmptyError		int32 = 5   // 节点为空
	StatusCode_MaxConnectError			int32 = 6	// 超过最大连接数
	StatusCode_StopServerError			int32 = 7   // 服务器停服
	StatusCode_TokenError				int32 = 8   // token错误
	StatusCode_DataExistError			int32 = 9   // 数据不存在
)
