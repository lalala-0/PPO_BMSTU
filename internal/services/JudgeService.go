package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/logger"
	"PPO_BMSTU/password_hash"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type JudgeService struct {
	JudgeRepository repository_interfaces.IJudgeRepository
	hash            password_hash.PasswordHash
	logger          *logger.CustomLogger
}

func NewJudgeService(JudgeRepository repository_interfaces.IJudgeRepository, hash password_hash.PasswordHash, logger *logger.CustomLogger) service_interfaces.IJudgeService {
	return &JudgeService{
		JudgeRepository: JudgeRepository,
		hash:            hash,
		logger:          logger,
	}
}

func (j JudgeService) checkIfJudgeWithLoginExists(login string) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "CheckIfJudgeWithLoginExists")
	defer span.End()

	j.logger.Info("SERVICE: Checking if Judge with login exists", "login", login)
	tempJudge, err := j.JudgeRepository.GetJudgeDataByLogin(ctx, login)

	if err != nil && errors.Is(err, repository_errors.DoesNotExist) {
		j.logger.Info("SERVICE: Judge with login does not exist", "login", login)
		span.SetStatus(codes.Ok, "Judge does not exist")
		return nil, nil
	} else if err != nil {
		j.logger.Error("SERVICE: GetJudgeBylogin method failed", "login", login, "error", err)
		span.SetStatus(codes.Error, "GetJudgeBylogin failed")
		return nil, err
	} else {
		j.logger.Info("SERVICE: Judge with login exists", "login", login)
		span.SetStatus(codes.Ok, "Judge exists")
		return tempJudge, nil
	}
}

func (j JudgeService) Login(login, password string) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "LoginJudge")
	defer span.End()

	j.logger.Info("SERVICE: Checking if Judge with login %s exists", login)
	tempJudge, err := j.checkIfJudgeWithLoginExists(login)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during checking if Judge with login exists")
		span.SetStatus(codes.Error, "Judge check failed")
		return nil, err
	} else if tempJudge == nil {
		j.logger.Info("SERVICE: Judge with login does not exist", "login", login)
		span.SetStatus(codes.Error, "Judge does not exist")
		return nil, repository_errors.DoesNotExist
	}

	j.logger.Info("SERVICE: Checking if password is correct for Judge with login %s", login)
	isPasswordCorrect := j.hash.CompareHashAndPassword(tempJudge.Password, password)
	if !isPasswordCorrect {
		j.logger.Info("SERVICE: Password is incorrect for Judge with login", "login", login)
		span.SetStatus(codes.Error, "Password mismatch")
		return nil, service_errors.MismatchedPassword
	}

	j.logger.Info("SERVICE: Successfully logged in Judge with login", "login", login)
	span.SetStatus(codes.Ok, "Login successful")
	return tempJudge, nil
}
func (j JudgeService) CreateProfile(judgeID uuid.UUID, fio string, login string, password string, role int, post string) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "CreateJudgeProfile")
	defer span.End()

	j.logger.Info("SERVICE: Validating data")
	if !validFIO(fio) || !validLogin(login) || !validRole(role) || !validPassword(password) {
		j.logger.Error("SERVICE: Invalid input data", "fio", fio, "login", login, "role", role, "password", password)
		span.SetStatus(codes.Error, "Invalid input data")
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	j.logger.Info("SERVICE: Checking if judge with login %s exists", login)
	tempJudge, err := j.checkIfJudgeWithLoginExists(login)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during checking if judge with login exists")
		span.SetStatus(codes.Error, "Judge check failed")
		return nil, err
	} else if tempJudge != nil {
		j.logger.Info("SERVICE: Judge with login exists", "login", login)
		span.SetStatus(codes.Error, "Judge already exists")
		return nil, service_errors.NotUnique
	}

	j.logger.Info("SERVICE: Creating new judge: %s", fio)
	hashedPassword, err := j.hash.GetHash(password)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during password hashing")
		span.SetStatus(codes.Error, "Password hashing failed")
		return nil, err
	} else {
		password = hashedPassword
	}

	var judge = &models.Judge{
		ID:       judgeID,
		FIO:      fio,
		Login:    login,
		Password: password,
		Role:     role,
		Post:     post,
	}

	createdJudge, err := j.JudgeRepository.CreateProfile(ctx, judge)
	if err != nil {
		j.logger.Error("SERVICE: Create method failed", "error", err)
		span.SetStatus(codes.Error, "Profile creation failed")
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully created new user with ", "id", createdJudge.ID)
	span.SetStatus(codes.Ok, "Profile created successfully")
	return createdJudge, nil
}

func (j JudgeService) DeleteProfile(id uuid.UUID) error {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "DeleteJudgeProfile")
	defer span.End()

	_, err := j.JudgeRepository.GetJudgeDataByID(ctx, id)
	if err != nil {
		j.logger.Error("SERVICE: GetJudgeDataByID method failed", "id", id, "error", err)
		span.SetStatus(codes.Error, "GetJudgeDataByID failed")
		return err
	}

	err = j.JudgeRepository.DeleteProfile(ctx, id)
	if err != nil {
		j.logger.Error("SERVICE: Delete method failed", "error", err)
		span.SetStatus(codes.Error, "Delete failed")
		return err
	}

	j.logger.Info("SERVICE: Successfully deleted judge", "id", id)
	span.SetStatus(codes.Ok, "Profile deleted successfully")
	return nil
}

func (j JudgeService) GetJudgeDataByID(id uuid.UUID) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetJudgeByID")
	defer span.End()

	judge, err := j.JudgeRepository.GetJudgeDataByID(ctx, id)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByID method failed", "id", id, "error", err)
		span.SetStatus(codes.Error, "GetJudgeByID failed")
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got user with GetJudgeByID", "id", id)
	span.SetStatus(codes.Ok, "GetJudgeByID successful")
	return judge, nil
}

func (j JudgeService) GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetJudgeByProtestID")
	defer span.End()

	judge, err := j.JudgeRepository.GetJudgeDataByProtestID(ctx, protestID)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByProtestID method failed", "id", protestID, "error", err)
		span.SetStatus(codes.Error, "GetJudgeByProtestID failed")
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got judge by protest ID", "id", protestID)
	span.SetStatus(codes.Ok, "GetJudgeByProtestID successful")
	return judge, nil
}

func (j JudgeService) GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetJudgesByRatingID")
	defer span.End()

	judges, err := j.JudgeRepository.GetJudgesDataByRatingID(ctx, ratingID)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgesByRatingID method failed", "id", ratingID, "error", err)
		span.SetStatus(codes.Error, "GetJudgesByRatingID failed")
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got judges by rating ID", "id", ratingID)
	span.SetStatus(codes.Ok, "GetJudgesByRatingID successful")
	return judges, nil
}

func (j JudgeService) GetAllJudges() ([]models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "GetAllJudges")
	defer span.End()

	judges, err := j.JudgeRepository.GetAllJudges(ctx)

	if err != nil {
		j.logger.Error("SERVICE: GetAllJudges method failed", "error", err)
		span.SetStatus(codes.Error, "GetAllJudges failed")
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got all judges")
	span.SetStatus(codes.Ok, "GetAllJudges successful")
	return judges, nil
}

func (j JudgeService) UpdateProfile(judgeID uuid.UUID, fio string, login string, password string, role int) (*models.Judge, error) {
	ctx := context.Background()

	tracer := otel.Tracer("service")
	_, span := tracer.Start(ctx, "UpdateJudgeProfile")
	defer span.End()

	judge, err := j.JudgeRepository.GetJudgeDataByID(ctx, judgeID)
	judgeCopy := judge

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByID method failed", "id", judgeID, "error", err)
		span.SetStatus(codes.Error, "GetJudgeByID failed")
		return judge, err
	}

	j.logger.Info("SERVICE: Validating data")
	if !validFIO(fio) {
		j.logger.Error("SERVICE: Invalid fio", "fio", fio)
		span.SetStatus(codes.Error, "Invalid fio")
		return judge, service_errors.InvalidFIO
	}

	if !validLogin(login) {
		j.logger.Error("SERVICE: Invalid login", "login", login)
		span.SetStatus(codes.Error, "Invalid login")
		return judge, service_errors.InvalidLogin
	}

	if !validRole(role) {
		j.logger.Error("SERVICE: Invalid role", "role", role)
		span.SetStatus(codes.Error, "Invalid role")
		return judge, service_errors.InvalidRole
	}

	if !validPassword(password) {
		j.logger.Error("SERVICE: Invalid password", "password", password)
		span.SetStatus(codes.Error, "Invalid password")
		return judge, service_errors.InvalidPassword
	}

	j.logger.Info("SERVICE: Updating judge: %s", fio)
	hashedPassword, err := j.hash.GetHash(password)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during password hashing")
		span.SetStatus(codes.Error, "Password hashing failed")
		return judge, err
	} else {
		password = hashedPassword
	}

	judge.Role = role
	judge.FIO = fio
	judge.Password = password
	judge.Login = login

	judge, err = j.JudgeRepository.UpdateProfile(ctx, judge)
	if err != nil {
		judge = judgeCopy
		j.logger.Error("SERVICE: UpdateJudge method failed", "error", err)
		span.SetStatus(codes.Error, "Update failed")
		return judge, err
	}

	j.logger.Info("SERVICE: Successfully updated judge", "judge", judge)
	span.SetStatus(codes.Ok, "Profile updated successfully")
	return judge, nil
}
