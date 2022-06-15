package controllers

import (
	"github.com/chirag3003/ecommerce-golang-api/helpers"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"github.com/chirag3003/ecommerce-golang-api/repository"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/razorpay/razorpay-go"
	"log"
	"os"
)

type Order interface {
	NewOrder(ctx *fiber.Ctx) error
	OrderPaid(ctx *fiber.Ctx) error
	OrderEvents(ctx *fiber.Ctx) error
	GetOrders(ctx *fiber.Ctx) error
	GetOrder(ctx *fiber.Ctx) error
}

func OrderControllers() Order {
	client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY"), os.Getenv("RAZORPAY_SECRET"))
	return &orderRoutes{
		Order:       repository.NewOrderRepository(conn.DB()),
		User:        repository.NewUserRepository(conn.DB()),
		Products:    repository.NewProductsRepository(conn.DB()),
		Transaction: repository.NewTransactionRepo(conn.DB()),
		Razorpay:    client,
	}
}

type orderRoutes struct {
	Order       repository.OrderRepository
	User        repository.UserRepository
	Products    repository.ProductsRepository
	Transaction repository.TransactionsRepository
	Razorpay    *razorpay.Client
}

func (c *orderRoutes) NewOrder(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	body := &models.NewOrderInput{}
	err := ctx.BodyParser(body)
	if err != nil || !helpers.ValidateNewOrderInput(body) {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//Finding and making the products array
	var products []models.OrderProduct
	var amount float64 = 0
	for _, prod := range body.Products {
		product, err := c.Products.FindByID(prod.Product)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		if product == nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		amount += (product.Price - product.Discount) * float64(prod.Quantity)
		products = append(products, models.OrderProduct{
			Product:  *product,
			Quantity: prod.Quantity,
			Size:     prod.Size,
		})

	}

	//Finding user address
	address, err := c.User.GetAddressByID(body.Address)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	//creating orderID nad order model
	orderID, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz123456789", 10)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	order := &models.Order{
		OrderID:       orderID,
		UserID:        user.ID,
		Address:       *address,
		Products:      products,
		OrderStatus:   "pending",
		PaymentMethod: "razorpay",
		PaymentStatus: "pending",
	}
	order.SetCreatedAt()

	//creating a transaction at razorpay
	data := map[string]interface{}{
		"amount":   amount * 100,
		"currency": "INR",
		"receipt":  order.OrderID,
	}
	rResp, err := c.Razorpay.Order.Create(data, nil)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	transaction, ok := helpers.FillTransactionModel(rResp)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	//saving transaction details in the database
	_, err = c.Transaction.NewTransaction(transaction)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	//modifying order struct according to the transaction
	order.TransactionID = transaction.Razorpay.RazorpayID
	//saving order to the database
	err = c.Order.Save(order)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response := &models.NewOrderResponse{
		TransactionID: transaction.Razorpay.RazorpayID,
		OrderID:       order.OrderID,
	}
	return ctx.JSON(response)
}

func (c *orderRoutes) GetOrders(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	orders, err := c.Order.GetOrders(user.ID)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if orders == nil {
		orders = &[]models.Order{}
	}
	return ctx.JSON(orders)
}
func (c *orderRoutes) GetOrder(ctx *fiber.Ctx) error {
	user := helpers.ParseUser(ctx)
	orderID := ctx.Params("orderID")
	order, err := c.Order.GetOrder(orderID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	if order.UserID != user.ID {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.JSON(order)
}

func (c *orderRoutes) OrderPaid(ctx *fiber.Ctx) error {
	if !helpers.ValidateWebhook(ctx) {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	payload, ok := body["payload"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	order, ok := payload["order"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	entity, ok := order["entity"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	transactionID := entity["id"].(string)
	orderID := entity["receipt"].(string)
	c.Transaction.AddEvent(transactionID, body)
	c.Transaction.SetStatus(transactionID, "paid")
	c.Order.SetPaid(orderID)
	c.Order.SetStatus(orderID, "placed")

	return ctx.SendStatus(200)

}

func (c *orderRoutes) OrderEvents(ctx *fiber.Ctx) error {
	if !helpers.ValidateWebhook(ctx) {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	payload, ok := body["payload"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	payment, ok := payload["payment"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	entity, ok := payment["entity"].(map[string]interface{})
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	transactionID := entity["order_id"].(string)
	c.Transaction.AddEvent(transactionID, body)
	return ctx.SendStatus(200)
}
