package pages

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

const (
	emailInputSelector    = "input[type='text']"
	passwordInputSelector = "input[type='password']"
	submitButtonSelector  = "button[type='submit']"
)

type LoginPage struct {
	page playwright.Page
}

func NewLoginPage(p playwright.Page) *LoginPage {
	return &LoginPage{
		page: p,
	}
}

func (l *LoginPage) FillEmail(email string) error {
	if err := l.page.Fill(emailInputSelector, email); err != nil {
		return fmt.Errorf("failed to fill email: %w", err)
	}
	return nil
}

func (l *LoginPage) FillPassword(password string) error {
	if err := l.page.Fill(passwordInputSelector, password); err != nil {
		return fmt.Errorf("failed to fill password: %w", err)
	}
	return nil
}

func (l *LoginPage) Submit() error {
	if err := l.page.Click(submitButtonSelector); err != nil {
		return fmt.Errorf("failed to submit login form: %w", err)
	}
	return nil
}

func (l *LoginPage) FullLogin(login, password string) error {
	if err := l.FillEmail(login); err != nil {
		return err
	}
	if err := l.FillPassword(password); err != nil {
		return err
	}
	return l.Submit()
}

func (l *LoginPage) WaitForLoginInput() error {
	_, err := l.page.WaitForSelector(emailInputSelector, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateVisible,
		Timeout: playwright.Float(5000),
	})
	return err
}
