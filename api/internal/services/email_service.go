package services

import (
	"bookmark-api/internal/config"
	"fmt"
	"log"
	"net/smtp"
)

type EmailService struct {
	smtpEmail    string
	smtpPassword string
	frontendURL  string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpEmail:    config.AppConfig.SMTPEmail,
		smtpPassword: config.AppConfig.SMTPPassword,
		frontendURL:  config.AppConfig.FrontendURL,
	}
}

func (es *EmailService) SendPasswordResetEmail(userEmail string, resetToken string) error {
	if es.smtpEmail == "" || es.smtpPassword == "" {
		log.Println("SMTP credentials not configured, skipping email send")
		return nil
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", es.frontendURL, resetToken)

	subject := "Password Reset Request"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<style>
		body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
		.container { max-width: 600px; margin: 0 auto; padding: 20px; background: #f9fafb; border-radius: 8px; }
		.header { background: #111827; color: white; padding: 20px; border-radius: 6px 6px 0 0; text-align: center; }
		.content { background: white; padding: 20px; border-radius: 0 0 6px 6px; }
		.button { display: inline-block; background: #111827; color: white; padding: 12px 24px; border-radius: 6px; text-decoration: none; font-weight: bold; margin: 20px 0; }
		.footer { text-align: center; margin-top: 20px; font-size: 12px; color: #6b7280; }
		.warning { color: #dc2626; font-size: 12px; margin-top: 20px; }
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>Password Reset Request</h1>
		</div>
		<div class="content">
			<p>Hello,</p>
			<p>We received a request to reset your password. Click the button below to set a new password:</p>
			
			<a href="%s" class="button">Reset Password</a>
			
			<p>Or copy this link in your browser:</p>
			<p style="word-break: break-all; background: #f3f4f6; padding: 10px; border-radius: 4px; font-size: 12px;">%s</p>
			
			<p class="warning">This link will expire in 24 hours.</p>
			<p class="warning">If you didn't request this, please ignore this email.</p>
		</div>
		<div class="footer">
			<p>&copy; 2026 Bookmark Manager. All rights reserved.</p>
		</div>
	</div>
</body>
</html>
	`, resetLink, resetLink)

	err := es.sendEmail(userEmail, subject, body)
	if err != nil {
		log.Printf("Failed to send password reset email: %v", err)
		return err
	}

	log.Printf("Password reset email sent to %s", userEmail)
	return nil
}

func (es *EmailService) sendEmail(to, subject, body string) error {
	from := es.smtpEmail
	password := es.smtpPassword

	// Gmail SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "465"

	// Construct headers
	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n", from, to, subject)

	// Send email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		[]byte(headers+body),
	)

	return err
}
