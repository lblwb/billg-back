package service

// GetOrdersByUser - Список сделанных заказов по данной услуге
//func GetOrdersByUser(c *fiber.Ctx) error {
//	paramsSlug := c.Params("slug")
//
//	modelService, err := service.GetServiceBySlug(paramsSlug)
//	fmt.Println(modelService)
//	if err != nil {
//		return c.JSON(fiber.Map{
//			"success": false,
//		})
//	} else {
//
//		return c.JSON(fiber.Map{
//			"service":  modelService,
//			"services": []service.Services{},
//			"count":    0,
//			"success":  true,
//		})
//	}
//
//}
