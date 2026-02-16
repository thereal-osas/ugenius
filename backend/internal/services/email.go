package services

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/ugenius/backend/internal/config"
)

type EmailService struct {
	config      *config.SMTPConfig
	frontendURL string
}

func NewEmailService(cfg *config.SMTPConfig, frontendURL string) *EmailService {
	return &EmailService{
		config:      cfg,
		frontendURL: frontendURL,
	}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	if s.config.Host == "" || s.config.User == "" {
		// Skip sending if SMTP not configured (development mode)
		fmt.Printf("[DEV] Email to %s: %s\n%s\n", to, subject, body)
		return nil
	}

	auth := smtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		s.config.From,
		to,
		subject,
		body,
	))

	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	return smtp.SendMail(addr, auth, s.config.User, []string{to}, msg)
}

func (s *EmailService) SendVerificationEmail(to, firstName, token string) error {
	verifyURL := fmt.Sprintf("%s/verify-email?token=%s", s.frontendURL, token)

	body := s.buildEmailTemplate(
		fmt.Sprintf("Welcome to U-Genius, %s!", firstName),
		`<p>Thank you for joining U-Genius! We're excited to have you on your journey to academic excellence.</p>
		<p>Please verify your email address by clicking the button below:</p>`,
		verifyURL,
		"Verify Email",
		"If you didn't create an account with U-Genius, please ignore this email.",
	)

	return s.SendEmail(to, "Verify Your U-Genius Account", body)
}

func (s *EmailService) SendPasswordResetEmail(to, firstName, token string) error {
	resetURL := fmt.Sprintf("%s/reset-password?token=%s", s.frontendURL, token)

	body := s.buildEmailTemplate(
		fmt.Sprintf("Password Reset Request"),
		fmt.Sprintf(`<p>Hi %s,</p>
		<p>We received a request to reset your password. Click the button below to create a new password:</p>`, firstName),
		resetURL,
		"Reset Password",
		"This link will expire in 1 hour. If you didn't request a password reset, please ignore this email.",
	)

	return s.SendEmail(to, "Reset Your U-Genius Password", body)
}

func (s *EmailService) SendSubmissionApprovedEmail(to, firstName, subject string, feedback string) error {
	body := s.buildEmailTemplate(
		"Reading Hours Approved! 🎉",
		fmt.Sprintf(`<p>Hi %s,</p>
		<p>Great news! Your reading hour submission for <strong>%s</strong> has been approved.</p>
		%s
		<p>Keep up the excellent work!</p>`,
			firstName,
			subject,
			formatFeedback(feedback),
		),
		fmt.Sprintf("%s/dashboard", s.frontendURL),
		"View Dashboard",
		"",
	)

	return s.SendEmail(to, "Your Reading Hours Were Approved", body)
}

func (s *EmailService) SendSubmissionRejectedEmail(to, firstName, subject string, feedback string) error {
	body := s.buildEmailTemplate(
		"Reading Hours Need Revision",
		fmt.Sprintf(`<p>Hi %s,</p>
		<p>Your reading hour submission for <strong>%s</strong> was not approved.</p>
		%s
		<p>Please review the feedback and submit again.</p>`,
			firstName,
			subject,
			formatFeedback(feedback),
		),
		fmt.Sprintf("%s/dashboard/reading-hours", s.frontendURL),
		"View Submission",
		"",
	)

	return s.SendEmail(to, "Reading Hours Submission Needs Revision", body)
}

func (s *EmailService) SendAchievementEmail(to, firstName, badgeTitle, badgeDesc string) error {
	body := s.buildEmailTemplate(
		fmt.Sprintf("New Achievement Unlocked! 🏆"),
		fmt.Sprintf(`<p>Congratulations %s!</p>
		<p>You've earned a new achievement: <strong>%s</strong></p>
		<p>%s</p>`,
			firstName,
			badgeTitle,
			badgeDesc,
		),
		fmt.Sprintf("%s/dashboard/achievements", s.frontendURL),
		"View Achievements",
		"",
	)

	return s.SendEmail(to, fmt.Sprintf("Achievement Unlocked: %s", badgeTitle), body)
}

func (s *EmailService) SendContactEmail(name, email, subject, message string) error {
	// Create contact email body
	contactBody := s.buildContactEmailTemplate(name, email, subject, message)

	// Send to admin email
	return s.SendEmail(s.config.From, fmt.Sprintf("U-Genius Contact: %s", subject), contactBody)
}

func formatFeedback(feedback string) string {
	if feedback == "" {
		return ""
	}
	return fmt.Sprintf(`<p><strong>Feedback:</strong> %s</p>`, feedback)
}

func (s *EmailService) buildContactEmailTemplate(name, email, subject, message string) string {
	return strings.TrimSpace(fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head><meta charset="UTF-8"></head>
<body style="font-family: 'Inter', Arial, sans-serif; background-color: #f5f5f5; padding: 40px 20px;">
	<div style="max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
		<div style="background: linear-gradient(135deg, #1a1a2e 0%%, #16213e 100%%); padding: 30px; text-align: center;">
			<h1 style="color: #D4A574; margin: 0; font-size: 28px;">📧 New Contact Message</h1>
			<p style="color: #ffffff; margin: 10px 0 0 0;">Someone has contacted you through the U-Genius website</p>
		</div>
		<div style="padding: 40px 30px;">
			<div style="margin-bottom: 25px;">
				<p style="color: #D4A574; font-weight: bold; margin: 0 0 5px 0;">👤 From:</p>
				<p style="color: #333; margin: 0; font-size: 16px;">%s (%s)</p>
			</div>
			
			<div style="margin-bottom: 25px;">
				<p style="color: #D4A574; font-weight: bold; margin: 0 0 5px 0;">📋 Subject:</p>
				<p style="color: #333; margin: 0; font-size: 16px; background: #f9f9f9; padding: 15px; border-radius: 8px; border-left: 4px solid #D4A574;">%s</p>
			</div>
			
			<div style="margin-bottom: 25px;">
				<p style="color: #D4A574; font-weight: bold; margin: 0 0 5px 0;">💬 Message:</p>
				<p style="color: #333; margin: 0; font-size: 16px; line-height: 1.6; background: #f9f9f9; padding: 15px; border-radius: 8px; border-left: 4px solid #D4A574;">%s</p>
			</div>
			
			<div style="text-align: center; margin-top: 30px;">
				<a href="mailto:%s" style="background-color: #D4A574; color: #1a1a2e; padding: 12px 30px; text-decoration: none; border-radius: 8px; font-weight: bold; display: inline-block;">Reply to %s</a>
			</div>
		</div>
		<div style="background: #f9f9f9; padding: 20px 30px; text-align: center; border-top: 1px solid #eee;">
			<p style="color: #888; font-size: 12px; margin: 0;">© 2026 U-Genius. Empowering Academic Excellence.</p>
		</div>
	</div>
</body>
</html>`, name, email, subject, message, email, name))
}

func (s *EmailService) buildEmailTemplate(title, content, buttonURL, buttonText, footer string) string {
	var footerHTML string
	if footer != "" {
		footerHTML = fmt.Sprintf(`<p style="color: #666; font-size: 12px; margin-top: 30px;">%s</p>`, footer)
	}

	var buttonHTML string
	if buttonURL != "" && buttonText != "" {
		buttonHTML = fmt.Sprintf(`
		<p style="text-align: center; margin: 30px 0;">
			<a href="%s" style="background-color: #D4A574; color: #1a1a2e; padding: 12px 30px; text-decoration: none; border-radius: 8px; font-weight: bold;">%s</a>
		</p>`, buttonURL, buttonText)
	}

	return strings.TrimSpace(fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head><meta charset="UTF-8"></head>
<body style="font-family: 'Inter', Arial, sans-serif; background-color: #f5f5f5; padding: 40px 20px;">
	<div style="max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
		<div style="background: linear-gradient(135deg, #1a1a2e 0%%, #16213e 100%%); padding: 30px; text-align: center;">
			<h1 style="color: #D4A574; margin: 0; font-size: 28px;">U-Genius</h1>
		</div>
		<div style="padding: 40px 30px;">
			<h2 style="color: #1a1a2e; margin-top: 0;">%s</h2>
			%s
			%s
			%s
		</div>
		<div style="background: #f9f9f9; padding: 20px 30px; text-align: center; border-top: 1px solid #eee;">
			<p style="color: #888; font-size: 12px; margin: 0;">© 2024 U-Genius. Empowering Academic Excellence.</p>
		</div>
	</div>
</body>
</html>`, title, content, buttonHTML, footerHTML))
}
