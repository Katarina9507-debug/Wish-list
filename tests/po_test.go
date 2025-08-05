package tests

import (
	"testing"

	"Wish-list/pages"
)

const (
	defaultTimeout = 5000 // 5 секунд
)

func TestLoginAndWishList(t *testing.T) {
	context, err := browser.NewContext()
	if err != nil {
		t.Fatalf("could not create context: %v", err)
	}
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}
	defer page.Close()

	loginPage := pages.NewLoginPage(page)
	wishListPage := pages.NewWishListPage(page)

	_, err = page.Goto("https://wishlist.otus.kartushin.su/")
	if err != nil {
		t.Fatalf("could not navigate to page: %v", err)
	}

	if err := loginPage.WaitForLoginInput(); err != nil {
		t.Fatalf("login field is not visible: %v", err)
	}

	if err := loginPage.FullLogin("MyNewName123", "MyNewName123MyNewName123"); err != nil {
		t.Fatalf("login failed: %v", err)
	}

	if isVisible, err := wishListPage.IsLoaded(defaultTimeout); !isVisible || err != nil {
		t.Fatalf("wishlist page is not load: %s", err)
	}
}
