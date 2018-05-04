package alidayu

import (
	"time"
)

// SendSMS send sms by alidayu
func SendSMS(recNum, smsFreeSignName, smsTemplateCode, smsParam string) (success bool, response string) {
	if recNum == "" || smsFreeSignName == "" || smsTemplateCode == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = MethodSendSMS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = smsFreeSignName
	params["rec_num"] = recNum
	params["sms_template_code"] = smsTemplateCode
	params["sms_param"] = smsParam

	return DoPost(params)

}
