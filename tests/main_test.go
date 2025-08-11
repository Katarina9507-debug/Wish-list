package tests

import (
	"github.com/playwright-community/playwright-go"
	"log"
	"os"
	"testing"
)

var (
	pw      *playwright.Playwright
	browser playwright.Browser
)

func TestMain(m *testing.M) {
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	code := m.Run()

	if browser != nil {
		browser.Close()
	}
	if pw != nil {
		pw.Stop()
	}

	os.Exit(code)
}
