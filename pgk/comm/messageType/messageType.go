package messageType

var (
	ONE  uint8 = 1
	MANY uint8 = 2
)
var MessageMap = map[uint8]string{
	ONE:  "单聊",
	MANY: "群聊",
}

func GetType(num uint8) string {
	return MessageMap[num]
}
