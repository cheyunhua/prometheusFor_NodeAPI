package mailX

import (
	"awesomeProject/logger"
	"awesomeProject/mode"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
)

//发送
func SendX() {
	nm := gomail.NewMessage()
	nm.SetHeader("From", model.C.Mail.From)
	nm.SetHeader("To", model.C.Mail.To)
	//抄送
	nm.SetAddressHeader("To", model.C.Mail.RTo, "test")
	//主题
	nm.SetHeader("Subject", "发送巡检报告")
	//正文
	nm.SetBody("text/html", "2022/11/2-2022/11/09")
	//所有附件
	nm.Attach(model.C.Job.JobName[0] + ".xlsx")
	nm.Attach(model.C.Job.JobName[1] + ".xlsx")
	//	nm.Attach("E:/xxxx.JPG")

	//ikfrnpwowsqgbibf:pop3密钥，374290910@qq.com为发件人
	nd := gomail.NewDialer(model.C.Mail.Host, model.C.Mail.Port, model.C.Mail.User, model.C.Mail.Pwd)
	nd.TLSConfig = &tls.Config{InsecureSkipVerify: model.C.Mail.IsSsl}
	if err := nd.DialAndSend(nm); err != nil {
		logger.DefaultLogger.Error(err)
	} else {
		fmt.Println("******************发送附件邮件成功*****************")
		fmt.Println("***************邮件服务器:smtp.qq.com******************")
		fmt.Println("************邮件发件人:************")
		fmt.Println("************邮件收件人:************")
		fmt.Println("************邮件主题:************")
		fmt.Println("************邮件内容:************")
		logger.DefaultLogger.Info(err)

	}
}

//	var stime string
//这里是我的需求为定时发送
//	flag.StringVar(&stime, "stime", "30 30 11 25 2021 2", "定时时间表达式 * * * * * ?")
//	flag.Parse()
//	fmt.Println("开始发送邮件时间:", stime)
//
