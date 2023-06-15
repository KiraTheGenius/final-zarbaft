package userRepository

import (
	"bankai/models"
	"bankai/utils"
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

	selectQuery := map[string]string{
		"fa": "carpets.name_fa as name, carpets.slug , collections.name_fa as collection_name , carpet_media.image",
		"en": "carpets.name_en as name, carpets.slug , collections.name_en as collection_name , carpet_media.image "}
	orderQuery := map[string]string{
		"newest":       "carpets.created_at desc",
		"mostـpopular": "carpets.most_popular desc",
		"best_selling": "carpets.best_selling desc"}

	ur.db.Model(&carpet).Select(selectQuery[locale]).Order(orderQuery[tag]).
		Joins("join collections on collections.id = carpets.collection_id").
		Joins("join carpet_colors on carpet_colors.carpet_id = carpets.id").
		Joins("join carpet_media on carpet_media.carpet_color_id = carpet_colors.id").
		Where("carpet_colors.default", "1").
		Where("carpet_media.feature", "main").Limit(8).Scan(&results)

	return results, nil
}

func (ur *userGormRepository) GetCollection(locale string, slug string) (*models.Collection, error) {
	var carpet models.Collection
	name := map[string]string{
		"fa": "name_fa",
		"en": "name_en "}
	motto := map[string]string{
		"fa": "motto_fa",
		"en": "motto_en "}
	result := ur.db.Select("ID", name[locale], motto[locale], "slug", "background", "collection").
		Where("slug = ?", slug).
		Preload("Carpets", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "collection_id", name[locale], "code_naqshe")
		}).
		Preload("Carpets.CarpetColors", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", name[locale], "default", "carpet_id").Where("default", "1")
		}).
		Preload("Carpets.CarpetColors.CarpetMedias", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "image", "carpet_color_id", "feature").Where("feature", "main").Or("feature", "background")
		}).
		First(&carpet)
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

	dbURI := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.ENV("DB_USERNAME"), utils.ENV("DB_PASSWORD"), utils.ENV("DB_PORT"), utils.ENV("DB_DATABASE"))
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
