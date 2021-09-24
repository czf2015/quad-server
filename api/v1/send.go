package apiv1

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

	"goserver/libs/conf"
)

func getMailableClient() *http.Client {
	b, err := ioutil.ReadFile("conf/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return getClient(config)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "conf/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func NotifySupport(replyTo string, title string, msg string) {
	SendEmail(conf.GetSectionKey("app", "SUPPORT_EMAIL").String(), "", "", replyTo, title, msg)
}

func SendContactUsEmail(replyTo string, title string, msg string) {
	message := "<pre>" + msg + "</pre>"
	SendEmail(conf.GetSectionKey("app", "SUPPORT_EMAIL").String(), "", "", replyTo, title, message)
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

func SendEmail(to string, cc string, bcc string, replyTo string, subject string, htmlBody string) {
	srv, err := gmail.New(getMailableClient())
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
		bcc = conf.GetSectionKey("app", "ADMIN_EMAIL").String()
	}
	if len(bcc) > 0 {
		other += "Bcc:" + bcc + "\r\n"
	}

	var message gmail.Message
	temp := []byte("From: " + conf.GetSectionKey("app", "ADMIN_EMAIL").String() + "\r\n" +
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