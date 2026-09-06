package services

import (
	"context"
	"errors"
	"strconv"
	"user-service/config"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/ports"
	"user-service/pkg/jwtutil"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type UserService struct {
	config      config.Config
	repository  ports.UserRepository
	roleService ports.RoleService
}

// Save implements [ports.UserService].
func (e *UserService) Save(ctx context.Context, user entity.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return e.repository.Create(ctx, user)
}

// Login implements [ports.UserService].
func (e *UserService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := e.repository.GetByUsername(ctx, username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	userID, err := strconv.Atoi(user.ID)
	if err != nil {
		return "", err
	}

	roleIDs, err := e.repository.GetUserRoleIDs(ctx, userID)
	if err != nil {
		return "", err
	}

	permissions, err := e.roleService.GetPermissionsByRoleIDs(ctx, roleIDs)
	if err != nil {
		return "", err
	}

	permissionCodes := make([]string, len(permissions))
	for i, permission := range permissions {
		permissionCodes[i] = permission.Resource + ":" + permission.Action
	}

	return jwtutil.GenerateToken(e.config.JWT.Secret, e.config.JWT.ExpiresIn, user.ID, user.Username, roleIDs, permissionCodes)
}

func NewUserService(cfg config.Config, repository ports.UserRepository, roleService ports.RoleService) ports.UserService {
	return &UserService{
		repository:  repository,
		config:      cfg,
		roleService: roleService,
	}
}
