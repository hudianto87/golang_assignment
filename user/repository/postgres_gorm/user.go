package postgresgorm

import (
	"context"
	"errors"
	"golangassignment/user/entity"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userRepository struct {
	db GormDBIface
}

func NewUserRepository(db GormDBIface) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("Error creating user : %v\n", err)
	}

	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "address", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		log.Printf("Error get user : %v\n", err)
		return entity.User{}, nil
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	var userExisting entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "address", "created_at", "updated_at").First(&userExisting, id).Error; err != nil {
		log.Printf("Error finding user : %v\n", err)
		return entity.User{}, nil
	}
	userExisting.Name = user.Name
	userExisting.Email = user.Email
	userExisting.Address = user.Address

	if err := r.db.WithContext(ctx).Save(&userExisting).Error; err != nil {
		return entity.User{}, nil
	}

	return userExisting, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		log.Printf("Error deleting user : %v\n", err)
		return err
	}

	return nil
}

// * itu pointer insert ke memory, & itu pointer mengambil dari memory
func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "address", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error get user : %v\n", err)
		return nil, err
	}

	return users, nil
}
