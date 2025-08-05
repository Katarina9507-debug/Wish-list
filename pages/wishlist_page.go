package pages

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
)

const (
	createNewListButton  = "button.btn-primary:has-text('Создать новый список')"
	wishListTextSelector = "div.d-flex.justify-content-between.align-items-center.mb-4 h2"
)

type WishList struct {
	page playwright.Page
}

func NewWishListPage(p playwright.Page) *WishList {
	return &WishList{
		page: p,
	}
}

func (w *WishList) WaitForCreateButton(timeout float64) error {
	_, err := w.page.WaitForSelector(createNewListButton, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateVisible,
		Timeout: playwright.Float(timeout),
	})
	return err
}

func (w *WishList) WaitForWishListText(timeout float64) error {
	_, err := w.page.WaitForSelector(wishListTextSelector, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateVisible,
		Timeout: playwright.Float(timeout),
	})
	return err
}

func (w *WishList) IsLoaded(timeout float64) (bool, error) {
	if err := w.WaitForCreateButton(timeout); err != nil {
		return false, fmt.Errorf("create button not visible: %w", err)
	}
	if _, err := w.page.WaitForSelector(wishListTextSelector, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateVisible,
		Timeout: playwright.Float(timeout),
	}); err != nil {
		return false, fmt.Errorf("wishlist header not visible: %w", err)
	}
	return true, nil
}
