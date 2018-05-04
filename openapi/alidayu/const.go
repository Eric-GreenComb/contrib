package alidayu

const (
	// URL alidayu url
	URL string = "http://gw.api.taobao.com/router/rest"
	// MethodSendSMS send sms
	MethodSendSMS string = "alibaba.aliqin.fc.sms.num.send"
	// MethodCallTTS call tts
	MethodCallTTS string = "alibaba.aliqin.fc.tts.num.singlecall"
	// MethodCallVoice call voice
	MethodCallVoice string = "alibaba.aliqin.fc.voice.num.singlecall"
	// MethodCallDouble call double
	MethodCallDouble string = "alibaba.aliqin.fc.voice.num.doublecall"
)

// AppKey app key
var AppKey string

// AppSecret app secret
var AppSecret string
