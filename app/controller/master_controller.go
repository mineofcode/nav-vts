package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/masagatech/nav-vts/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Master_controller struct {
	base_controller
}

// Initialize controller constructor
func (o *Master_controller) Initr(app *models.App) {
	// setting app context
	o.super(app)

	// intialize router for controller

	// set api version group
	d := o.App.Fiber.Group("master/")
	d.Get("", func(c *fiber.Ctx) error {
		// o.App.DB.Collection("testcol").InsertOne(context.Background(), bson.M{
		// 	"a": "b",
		// })

		k, _ := o.App.DB.Collection("testcol").Find(c.Context(), bson.M{})
		var result []interface{}

		k.All(c.Context(), &result)
		//fmt.Println(result)
		d, _ := json.Marshal(result)
		return c.Send(d)
	})
	// actual route
}
