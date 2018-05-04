## alidayu
fork from github.com/ltt1987/alidayu, just for alidayu sms send

阿里大鱼Go语言开发包。One Golang package for alidayu service. 

> 阿里大鱼API说明文档：[http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.bkKKhG&apiId=25450](http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.bkKKhG&apiId=25450)  
官网：[http://alidayu.com](http://alidayu.com)

###使用方法：

- 发送短信：`alidayu.SendSMS`

```

func main() {
	alidayu.AppKey = "...your AppKey..."
	alidayu.AppSecret = "...your AppSecret..."

	success, resp := alidayu.SendSMS("18888888888", "身份验证", "SMS_4000328", `{"code":"1234","product":"alidayu"}`)
	fmt.Println("Success:", success)
	fmt.Println(resp)
}
```