package contentType

var (
	TEXT      uint8 = 1
	FILE      uint8 = 2
	IMG       uint8 = 3
	AUDIO     uint8 = 4
	VIDEO     uint8 = 5
	AUDIOCHAT uint8 = 6
	VIDEOCHAT uint8 = 7
)
var ContentMap = map[uint8]string{
	TEXT:      "文本",
	FILE:      "文件",
	IMG:       "图片",
	AUDIO:     "音频",
	VIDEO:     "视频",
	AUDIOCHAT: "语音聊天",
	VIDEOCHAT: "视频聊天",
}

func GetType(num uint8) string {
	return ContentMap[num]
}
