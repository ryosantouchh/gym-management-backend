package repository

type UserRepositoryImpl struct {
	// db *gorm.DB
	db interface{}
}

func NewUserRepositoryImpl(db interface{}) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
