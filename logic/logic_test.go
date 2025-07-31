package logic

import (
	"github.com/playwright-community/playwright-go"
	"testing"
	"time"
)

func TestWishlistLogin(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, _ := browser.NewPage()
	defer page.Close()

	if _, err = page.Goto("https://wishlist.otus.kartushin.su/login"); err != nil {
		t.Fatalf("could not goto: %v", err)
	}

	// Ожидаем загрузки страницы
	page.WaitForSelector("form")

	// Локаторы
	loginField := page.Locator("input[type='text']")                                     // Поле ввода логина
	passwordField := page.Locator("input[type='password']")                              // Поле ввода пароля
	loginButton := page.Locator("button[type='submit']")                                 //Кнопка входа
	addItemButton := page.Locator("button.btn-primary:has-text('Создать новый список')") // Кнопка "Создать новый список"
	wishList := page.Locator(".navbar-brand")                                            // Список желаний

	// Проверяем видимость элементов
	if visible, err := loginField.IsVisible(); err != nil || !visible {
		t.Errorf("Login field is not visible")
	}

	if visible, err := passwordField.IsVisible(); err != nil || !visible {
		t.Errorf("Password field is not visible")
	}

	if visible, err := loginButton.IsVisible(); err != nil || !visible {
		t.Errorf("Login button is not visible")
	}

	// Ввод логина и пароля
	if err := loginField.Fill("MyNewName123"); err != nil {
		t.Errorf("Could not fill login field: %v", err)
	}
	if err := passwordField.Fill("MyNewName123MyNewName123"); err != nil {
		t.Errorf("Could not fill password field: %v", err)
	}

	// Клик с ожиданием навигации
	if err := loginButton.Click(playwright.LocatorClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Fatalf("login click failed: %v", err)
	}

	time.Sleep(2 * time.Second)

	// Ожидание появления селектора на странице
	page.WaitForSelector(
		"button.btn-primary:has-text('Создать новый список')",
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateVisible,
			Timeout: playwright.Float(10000),
		},
	)

	// Проверяем видимость элементов после авторизации
	if visible, err := addItemButton.IsVisible(); !visible || err != nil {
		t.Error("Add item button is not visible")
	}

	if visible, err := wishList.IsVisible(); !visible || err != nil {
		t.Error("Wish list is not visible")
	}
}
