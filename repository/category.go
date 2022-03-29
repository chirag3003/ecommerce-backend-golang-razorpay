package repository

import (
	"context"
	"fmt"
	"github.com/chirag3003/ecommerce-golang-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	SaveCategory(data *models.Category) (*mongo.InsertOneResult, error)
	SaveSubcategory(data *models.Subcategory, categoryID string) (*mongo.UpdateResult, error)
	FindCategory(id string) (*models.Category, error)
	FindAll() ([]models.Category, error)
	FindSubcategory(id string) (*models.Subcategory, error)
	UpdateCategory(id string, data *models.CategoryUpdateInput) (*mongo.UpdateResult, error)
	UpdateSubcategory(id string, data *models.CategoryUpdateInput) (*mongo.UpdateResult, error)
	DeleteCategory(id string) (*mongo.DeleteResult, error)
	DeleteSubcategory(id string) (*mongo.UpdateResult, error)
	ChangeVisibility(id string, p bool) (*mongo.UpdateResult, error)
}

type categoryRepo struct {
	db *mongo.Collection
}

func NewCategoryRepo(Category *mongo.Collection) CategoryRepository {
	return &categoryRepo{
		db: Category,
	}
}

func (c *categoryRepo) SaveCategory(data *models.Category) (*mongo.InsertOneResult, error) {
	one, err := c.db.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	return one, nil
}
func (c *categoryRepo) SaveSubcategory(data *models.Subcategory, categoryID string) (*mongo.UpdateResult, error) {
	ID, err2 := primitive.ObjectIDFromHex(categoryID)
	if err2 != nil {
		return nil, err2
	}
	one, err := c.db.UpdateByID(context.TODO(), ID, bson.M{"$push": bson.M{"subcategories": data}})
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (c *categoryRepo) FindCategory(id string) (*models.Category, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	find := c.db.FindOne(context.TODO(), bson.M{"_id": ID})
	data := &models.Category{}
	err = find.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *categoryRepo) FindAll() ([]models.Category, error) {
	find, err := c.db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var data []models.Category
	err = find.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *categoryRepo) FindSubcategory(id string) (*models.Subcategory, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	find := c.db.FindOne(context.TODO(), bson.M{"subcategories._id": ID})
	data := &models.Category{}
	err = find.Decode(data)
	if err != nil {
		return nil, err
	}
	var subcategory *models.Subcategory
	for _, s := range data.Subcategories {
		if s.ID.String() == ID.String() {
			subcategory = &s
			break
		}
	}
	return subcategory, nil
}
func (c *categoryRepo) UpdateCategory(id string, data *models.CategoryUpdateInput) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	byID, err := c.db.UpdateByID(context.TODO(), ID,
		bson.M{"$set": data})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byID, nil
}
func (c *categoryRepo) UpdateSubcategory(id string, data *models.CategoryUpdateInput) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	data.ID = ID
	byID, err := c.db.UpdateOne(context.TODO(), bson.M{"subcategories._id": ID},
		bson.M{"$set": bson.M{"subcategories.$": data}})
	if err != nil {
		return nil, err
	}

	return byID, nil
}
func (c *categoryRepo) DeleteCategory(id string) (*mongo.DeleteResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	one, err := c.db.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (c *categoryRepo) DeleteSubcategory(id string) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	one, err := c.db.UpdateOne(context.TODO(), bson.M{"subcategories._id": ID}, bson.M{"$pull": bson.M{"subcategories": bson.M{"_id": ID}}})
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (c *categoryRepo) ChangeVisibility(id string, p bool) (*mongo.UpdateResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	byID, err := c.db.UpdateByID(context.TODO(), ID,
		bson.D{{"$set", bson.D{{"public", p}}}})
	if err != nil {
		return nil, err
	}

	return byID, nil
}
