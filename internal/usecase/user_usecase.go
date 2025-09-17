package usecase

import "github.com/HericVirgilio/api-go/internal/domain"

// UserUseCase é a interface que define os casos de uso para o usuário.
// Embora não seja estritamente necessário ter uma interface aqui para um CRUD simples,
// é uma boa prática para projetos maiores e para facilitar os testes.
type UserUseCase interface {
	CreateUser(name, email string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	UpdateUser(id uint, name, email string) (*domain.User, error)
	DeleteUser(id uint) error
}

// userUseCase é a implementação da interface UserUseCase.
type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase cria uma nova instância do caso de uso do usuário.
// Recebe o repositório como uma dependência (seguindo a inversão de dependência).
func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (uc *userUseCase) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		Name:  name,
		Email: email,
	}
	err := uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) GetAllUsers() ([]domain.User, error) {
	return uc.userRepo.FindAll()
}

func (uc *userUseCase) GetUserByID(id uint) (*domain.User, error) {
	return uc.userRepo.FindByID(id)
}

func (uc *userUseCase) UpdateUser(id uint, name, email string) (*domain.User, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	err = uc.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) DeleteUser(id uint) error {
	return uc.userRepo.Delete(id)
}