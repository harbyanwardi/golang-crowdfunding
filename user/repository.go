package user

import "gorm.io/gorm"

type Repository interface { //*huruf depan kapital berarti bersifat public
	Save(user User) (User, error)
}

type repository struct { //*huruf depan kecil berarti bersifat privat
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
