package service

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/token_manager"
	"time"
)

type userService struct {
	repo            repository.User
	tokenManager    token_manager.TokenManager
	refreshTokenTtl time.Duration
}

func NewUserService(repo repository.User, tokenManager token_manager.TokenManager, refreshTokenTtl time.Duration) *userService {
	return &userService{
		repo:            repo,
		tokenManager:    tokenManager,
		refreshTokenTtl: refreshTokenTtl,
	}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*dto.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, userId int) (*dto.User, error) {
	return s.repo.GetUserByID(ctx, userId)
}

func (s *userService) TakeBook(ctx context.Context, bookId, userId int) error {
	if bookId == 0 || userId == 0 {
		return fmt.Errorf("")
	}

	ok, err := s.repo.CheckUserBook(ctx, bookId, userId)

	if !ok {
		return s.repo.TakeBook(ctx, bookId, userId)
	}
	return err
}

func (s *userService) RegisterUser(ctx context.Context, name string) (*dto.SuccessRegister, error) {
	userId, err := s.repo.RegisterUser(ctx, name)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"user_name_key\"" {
			return nil, fmt.Errorf("user already registered")
		}
		return nil, err
	}

	accessToken, err := s.tokenManager.NewJWT(userId)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()

	if err != nil {
		return nil, err
	}

	if err = s.repo.WriteRefreshToken(ctx, refreshToken, userId, s.refreshTokenTtl); err != nil {
		return nil, err
	}

	return &dto.SuccessRegister{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) UserValidate(accessToken string) (string, error) {
	/*
		Check if token expired, return false
		------------------------------------
		I took the logic of token validation to the service level to make it easier to scale the application
		in the future when extending the functionality of the application and expanding token validation
		(verification by user agent or ip)
	*/

	/* We always return userId, because in UpdateAccessToken we need it when we're updating access token */

	userId, err := s.tokenManager.Parse(accessToken)

	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *userService) GetSubFromToken(accessToken string) (string, error) {
	/*
		Check if access token is valid
		------------------------------
		func return err with userId only when token is expired
	*/
	userId, err := s.tokenManager.Parse(accessToken)
	/*
		If token is valid return that token is valid
	*/
	if err == nil {
		return "", fmt.Errorf("access token is valid")
	}

	/*
		If error with parsing jwt token
	*/
	if err != nil && userId == "" {
		return "", fmt.Errorf("access token is invalid")
	}

	return userId, nil
}

func (s *userService) UpdateAccessToken(ctx context.Context, accessToken, refreshToken string) (string, error) {
	/*
		Check if access token is valid       		 --+
		Check if access token is not expired 		 --+
		Check if user refresh token is equal given token
		Check if refresh token is not expired
		Update access token
	*/

	/*
		Check if access token is valid
		------------------------------
		func return err with userId only when token is expired
	*/
	userId, err := s.GetSubFromToken(accessToken)

	if err != nil {
		return "", err
	}

	/*
		Check if refresh token is not expired
	*/
	refreshTokenDB, err := s.repo.CheckRefreshToken(ctx, userId)

	if err != nil {
		return "", fmt.Errorf("refresh token not found")
	}

	/*
		Check if user refresh token is equal given token
	*/
	if refreshTokenDB != refreshToken {
		return "", fmt.Errorf("refresh token is not correct")
	}

	/*
		Update access token
	*/
	newAccessToken, err := s.tokenManager.NewJWT(userId)

	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *userService) UpdateRefreshToken(ctx context.Context, accessToken, refreshToken string) (*dto.UpdateTokens, error) {

	/*
		Check if access token is invalid 	  ---+
		Get userId from it				 	  ---+
		Check if refresh token exist (not expired)
		Create new refresh token
		Create new access token
		Update refresh token
	*/

	/*
		Check if access token is valid
		------------------------------
		func return err with userId only when token is expired
	*/
	userId, err := s.GetSubFromToken(accessToken)

	if err != nil {
		return nil, err
	}

	/* Check if refresh token exist (not expired) */
	refreshTokenDB, err := s.repo.CheckRefreshToken(ctx, userId)

	if err != nil {
		return nil, err
	}

	if refreshTokenDB != "" {
		return nil, fmt.Errorf("refresh token is valid")
	}

	/* Create new refresh token */
	newRefreshToken, err := s.tokenManager.NewRefreshToken()

	if err != nil {
		return nil, err
	}

	/* Create new access token */
	newAccessToken, err := s.tokenManager.NewJWT(userId)

	if err != nil {
		return nil, err
	}

	/* Update refresh token */
	if err = s.repo.WriteRefreshToken(ctx, newRefreshToken, userId, s.refreshTokenTtl); err != nil {
		return nil, err
	}

	return &dto.UpdateTokens{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
