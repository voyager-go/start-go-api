package schedule_task

import (
	"crypto/tls"
	"github.com/RichardKnop/machinery/v1/log"
	"gopkg.in/gomail.v2"
)

// GetTaskLists 获取需要注册的任务列表
func GetTaskLists() map[string]interface{} {
	return map[string]interface{}{
		"send_email": SendEmailSample,
		"add":        Add,
		"mul":        Multiply,
		"concat":     Concat,
	}
}

func Concat(strs []string) (string, error) {
	var res string
	for _, s := range strs {
		res += s
	}
	return res, nil
}

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func Multiply(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	return sum, nil
}

type EmailInfo struct {
	User, Password, Host, To, Subject, Body, MailType string
}

// SendEmailSample 发送邮件示例
func SendEmailSample() (err error) {
	log.INFO.Print("Before send email ...")
	info := EmailInfo{
		User:     "", // 请自行配置
		Password: "", // 请自行配置
		Host:     "smtp.163.com",
		To:       "", // 请自行配置
		Subject:  "Hello Gopher",
		Body:     "Golang is the greatest language in the world!",
		MailType: "html",
	}
	err = SendEmail(info)
	if err != nil {
		log.INFO.Print("send email failed: ", err.Error())
	}
	return err
}

// SendEmail 任务方法
func SendEmail(emailInfo EmailInfo) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailInfo.User)
	m.SetHeader("To", emailInfo.To)
	//m.SetAddressHeader("Cc", emailInfo.User, "Dan")
	m.SetHeader("Subject", emailInfo.Subject)
	var contentType string
	if emailInfo.MailType == "html" {
		contentType = "Content-Type: text/" + emailInfo.MailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	m.SetBody(contentType, emailInfo.Body)
	d := gomail.NewDialer(emailInfo.Host, 25, emailInfo.User, emailInfo.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}
