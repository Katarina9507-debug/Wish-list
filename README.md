# Автотест для Wishlist (Playwright + Go)

## Описание
Тест проверяет авторизацию на сайте Wishlist:
1. Открывает страницу входа
2. Проверяет видимость полей формы
3. Вводит тестовые данные
4. Проверяет успешный вход

## Установка

1. Установите Go (версия 1.18+)
2. Установите Playwright:

```bash
go get github.com/playwright-community/playwright-go
go install github.com/playwright-community/playwright-go/cmd/playwright
playwright install
playwright install chromium# Wish-list
