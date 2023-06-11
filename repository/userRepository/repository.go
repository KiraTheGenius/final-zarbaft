package userRepository

import (
	"bankai/models"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByUserId(userId uint) (*models.User, error)
	DeleteUser(user *models.User) error
	GetSlider(locale string, tag string) ([]map[string]interface{}, error)
	GetCollection(locale string, slug string) (*models.Collection, error)
}

type userGormRepository struct {
	db *gorm.DB
}

func NewGormUserRepository() UserRepository {
	return &userGormRepository{
		db: getDbConnection(),
	}
}

func (ur *userGormRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *userGormRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userGormRepository) GetSlider(locale string, tag string) ([]map[string]interface{}, error) {
	var carpet models.Carpet
	results := []map[string]interface{}{}
	selectQuery, orderQuery := "", ""
	switch locale {
	case "fa":
		selectQuery = "carpets.name_fa as name, carpets.slug , collections.name_fa as collection_name , carpet_media.image "
	case "en":
		selectQuery = "carpets.name_en as name, carpets.slug , collections.name_en as collection_name , carpet_media.image "
	}

	switch tag {
	case "newest":
		orderQuery = "carpets.created_at desc"
	case "mostـpopular":
		orderQuery = "carpets.most_popular desc"
	case "best_selling":
		orderQuery = "carpets.best_selling desc"

	}

	ur.db.Model(&carpet).Select(selectQuery).Order(orderQuery).
		Joins("join collections on collections.id = carpets.collection_id").
		Joins("join carpet_colors on carpet_colors.carpet_id = carpets.id").
		Joins("join carpet_media on carpet_media.carpet_color_id = carpet_colors.id").
		Where("carpet_colors.default", "1").
		Where("carpet_media.feature", "main").Limit(8).Scan(&results)

	return results, nil
}

func (ur *userGormRepository) GetCollection(locale string, slug string) (*models.Collection, error) {
	var carpet models.Collection

	result := ur.db.Where("slug = ?", slug).Preload("Carpets", func(db *gorm.DB) *gorm.DB {
		return db.Select("CollectionID", "name_fa")
	}).First(&carpet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &carpet, nil
}

func (ur *userGormRepository) DeleteUser(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur *userGormRepository) GetUserByUserId(userId uint) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func getDbConnection() *gorm.DB {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	dbURI := "admin:13771377Ab?@tcp(localhost:3306)/zarbaft?charset=utf8&parseTime=True&loc=Local"
	// Connect to the database

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Set up connection pool and other configuration options

	// Enable logging in development mode
	// Migrate the User model to the database (if necessary)
	db.AutoMigrate(&models.Collection{}, &models.Carpet{}, &models.CarpetColor{}, &models.CarpetMedia{})
	// db.Model(&models.Carpet{}).AddForeignKey("collection_id", "collections(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.CarpetColor{}).AddForeignKey("carpet_id", "carpets(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.CarpetMedia{}).AddForeignKey("carpet_color_id", "carpet_colors(id)", "RESTRICT", "RESTRICT")
	// db.Model(&models.CarpetMedia{}).AddForeignKey("carpet_id", "carpets(id)", "RESTRICT", "RESTRICT")

	// Use the db instance to interact with the database in your application
	return db
}
