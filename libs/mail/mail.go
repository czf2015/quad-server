package mail

import (
	"encoding/base64"
	"log"
	"strings"
	"io/ioutil"
	"net/http"

	"google.golang.org/api/gmail/v1"
	"golang.org/x/oauth2/google"

	"goserver/libs/conf"
	"goserver/libs/oauth2"
)

var (
	adminEmail = conf.GetSectionKey("app", "ADMIN_EMAIL").String()
	supportEmail = conf.GetSectionKey("app", "SUPPORT_EMAIL").String()
)

func getSendMailClient() *http.Client {
	b, err := ioutil.ReadFile("conf/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return oauth2.GetClient(config)
}

func SendEmail(to string, cc string, bcc string, replyTo string, subject string, htmlBody string) {
	srv, err := gmail.New(getSendMailClient())
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}

	if len(replyTo) == 0 {
		replyTo = "no-reply@rev-stream.com"
	}
	var other = ""
	if len(cc) > 0 {
		other += "Cc:" + cc + "\r\n"
	}
	if len(bcc) == 0 {
		// just to monitor all emails through adminEmail by default
		bcc = adminEmail
	}
	if len(bcc) > 0 {
		other += "Bcc:" + bcc + "\r\n"
	}

	var message gmail.Message
	temp := []byte("From: " + adminEmail + "\r\n" +
		"reply-to: " + replyTo + "\r\n" +
		"To:  " + to + "\r\n" +
		other +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html;charset=utf-8\r\n" +
		"\r\n" +
		"<html><head><link rel=\"Shortcut Icon\" type=\"image/x-icon\" href=\"https://rev-stream.com/favicon.png\"/></head><body>" +
		htmlBody +
		"<br><div><div><font face=\"arial, sans-serif\" color=\"#000000\">Rev Stream LLC</font></div><div><a href=\"http://www.rev-stream.com\" target=\"_blank\">www.rev-stream.com</a></div><div><font color=\"#000000\"><img src=\"https://docs.google.com/uc?export=download&id=1W6UL9in16djTHw7kkpLiPQbervgydw-0&revid=0B8BoPWrntmRuQWRmT3BVYVJZWkJrUWwvMDdpajdGTHV4eWI4PQ\" width=\"200\" height=\"25\"><br></font></div></div></body></html>")

	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)

	_, err = srv.Users.Messages.Send("me", &message).Do()
}

func NotifySupport(replyTo string, title string, msg string) {
	SendEmail(supportEmail, "", "", replyTo, title, msg)
}

func SendContactUsEmail(replyTo string, title string, msg string) {
	message := "<pre>" + msg + "</pre>"
	SendEmail(supportEmail, "", "", replyTo, title, message)
	receipt := "<div>We will get back to you after receiving your email.</div>" +
		"<br><div>------Below is your message to us-----</div>" +
		message
	SendEmail(replyTo, "", "", "", "Thanks for contacting us!", receipt)
}

func SendWelcomeEmail(to string, confirmLink string) {
	msg := "<div>You have signed up successfully.</div>" +
		"<div>Please click the following link to confirm your email and activate your account.</div>" +
		"<div>" + confirmLink + "</div>"
	SendEmail(to, "", "", "", "Welcome to Rev-Stream", msg)
	NotifySupport(to, to+" just joined us!", "")
}

func SendResetPasswordEmail(to string, resetLink string) {
	msg := "<div>Please click the following link to reset your password.</div>" +
		"<div>" + resetLink + "</div>"
	SendEmail(to, "", "", "", "Reset your password", msg)
}
