package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sk25469/go-mongodb-server/pkg/config"
	"github.com/sk25469/go-mongodb-server/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllEmployee(ctx *fiber.Ctx) error {
	query := bson.D{{}}
	mg := config.GetMongoInstance()
	cursor, err := mg.Db.Collection("employees").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var employees []models.Employee = make([]models.Employee, 0)

	if err := cursor.All(ctx.Context(), &employees); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.JSON(employees)
}

func AddNewEmployee(ctx *fiber.Ctx) error {
	mg := config.GetMongoInstance()
	collection := mg.Db.Collection("employees")

	employee := new(models.Employee)

	if err := ctx.BodyParser(employee); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	insertionResult, err := collection.InsertOne(ctx.Context(), employee)

	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)

	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	return ctx.Status(201).JSON(createdEmployee)
}

func UpdateEmployeeById(ctx *fiber.Ctx) error {
	mg := config.GetMongoInstance()
	idParam := ctx.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return ctx.SendStatus(400)
	}

	employee := new(models.Employee)

	if err := ctx.BodyParser(employee); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = mg.Db.Collection("employees").FindOneAndUpdate(ctx.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.SendStatus(400)
		}
		return ctx.SendStatus(500)
	}

	employee.ID = idParam

	return ctx.Status(200).JSON(employee)
}

func DeleteEmployeeById(ctx *fiber.Ctx) error {
	mg := config.GetMongoInstance()
	employeeID, err := primitive.ObjectIDFromHex(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mg.Db.Collection("employees").DeleteOne(ctx.Context(), &query)

	if err != nil {
		return ctx.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return ctx.SendStatus(404)
	}

	return ctx.Status(200).JSON("record deleted")
}
