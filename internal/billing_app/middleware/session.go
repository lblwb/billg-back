package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Создайте новый экземпляр хранилища сессий
var store = session.New()

func setSessionValue(c *fiber.Ctx, key string, value interface{}) {
	// Получите сессию из контекста
	sess := c.Locals("session").(*session.Session)
	// Установите значение в сессию
	sess.Set(key, &value)
	// Сохраните сессию
	sess.Save()
}

// SessionMiddleware Middleware для работы с сессией
func SessionMiddleware(c *fiber.Ctx) error {
	//Установка значения
	setSessionValue(c, "balance", 100)

	// Получите сессию из хранилища
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	// Установите сессию в контекст
	c.Locals("session", sess)

	// Продолжите выполнение цепочки middleware и обработку запроса
	return c.Next()
}
