package services

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/repository/repository_errors"
	"PPO_BMSTU/internal/repository/repository_interfaces"
	"PPO_BMSTU/internal/services/service_errors"
	"PPO_BMSTU/internal/services/service_interfaces"
	"PPO_BMSTU/password_hash"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type JudgeService struct {
	JudgeRepository repository_interfaces.IJudgeRepository
	hash            password_hash.PasswordHash
	logger          *log.Logger
	JWTService      service_interfaces.IJWTService // Новый сервис для работы с JWT

}

func NewJudgeService(JudgeRepository repository_interfaces.IJudgeRepository, hash password_hash.PasswordHash, logger *log.Logger, jwtService service_interfaces.IJWTService) service_interfaces.IJudgeService {
	return &JudgeService{
		JudgeRepository: JudgeRepository,
		hash:            hash,
		logger:          logger,
		JWTService:      jwtService, // Инициализируем новый JWTService

	}
}

func (j JudgeService) checkIfJudgeWithLoginExists(login string) (*models.Judge, error) {
	j.logger.Info("SERVICE: Checking if Judge with login exists", "login", login)
	tempJudge, err := j.JudgeRepository.GetJudgeDataByLogin(login)

	if err != nil && errors.Is(err, repository_errors.DoesNotExist) {
		j.logger.Info("SERVICE: Judge with login does not exist", "login", login)
		return nil, nil
	} else if err != nil {
		j.logger.Error("SERVICE: GetJudgeBylogin method failed", "login", login, "error", err)
		return nil, err
	} else {
		j.logger.Info("SERVICE: Judge with login exists", "login", login)
		return tempJudge, nil
	}
}

func (j JudgeService) Login(login, password string) (*models.Judge, string, error) {
	j.logger.Infof("SERVICE: Checking if Judge with login %s exists", login)
	// Проверяем, существует ли судья с таким логином
	tempJudge, err := j.checkIfJudgeWithLoginExists(login)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during checking if Judge with login exists")
		return nil, "", err
	} else if tempJudge == nil {
		j.logger.Info("SERVICE: Judge with login does not exist", "login", login)
		return nil, "", repository_errors.DoesNotExist
	}

	j.logger.Infof("SERVICE: Checking if password is correct for Judge with login %s", login)
	// Проверяем пароль
	isPasswordCorrect := j.hash.CompareHashAndPassword(tempJudge.Password, password)
	if !isPasswordCorrect {
		j.logger.Info("SERVICE: Password is incorrect for Judge with login", "login", login)
		return nil, "", service_errors.MismatchedPassword
	}

	j.logger.Info("SERVICE: Successfully logged in Judge with login", "login", login)

	// Генерация JWT токена после успешного логина
	token, err := j.JWTService.GenerateToken(tempJudge.ID)
	if err != nil {
		j.logger.Error("SERVICE: Failed to generate JWT token", "error", err)
		return nil, "", err
	}

	// Возвращаем пользователя и сгенерированный JWT токен
	return tempJudge, token, nil
}

func (j JudgeService) CreateProfile(judgeID uuid.UUID, fio string, login string, password string, role int, post string) (*models.Judge, error) {
	j.logger.Info("SERVICE: Validating data")
	if !validFIO(fio) || !validLogin(login) || !validRole(role) || !validPassword(password) {
		j.logger.Error("SERVICE: Invalid input data", "fio", fio, "login", login, "role", role, "password", password)
		return nil, fmt.Errorf("SERVICE: Invalid input data")
	}

	j.logger.Infof("SERVICE: Checking if judgeView with login %s exists", login)
	tempJudge, err := j.checkIfJudgeWithLoginExists(login)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during checking if judgeView with login exists")
		return nil, err
	} else if tempJudge != nil {
		j.logger.Info("SERVICE: Judge with login exists", "login", login)
		return nil, service_errors.NotUnique
	}

	j.logger.Infof("SERVICE: Creating new judgeView: %s", fio)
	hashedPassword, err := j.hash.GetHash(password)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during password hashing")
		return nil, err
	} else {
		password = hashedPassword
	}

	// creating judgeView
	var judge = &models.Judge{
		ID:       judgeID,
		FIO:      fio,
		Login:    login,
		Password: password,
		Role:     role,
		Post:     post,
	}

	createdJudge, err := j.JudgeRepository.CreateProfile(judge)
	if err != nil {
		j.logger.Error("SERVICE: Create method failed", "error", err)
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully created new user with ", "id", createdJudge.ID)

	return createdJudge, nil
}

func (j JudgeService) DeleteProfile(id uuid.UUID) error {
	_, err := j.JudgeRepository.GetJudgeDataByID(id)
	if err != nil {
		j.logger.Error("SERVICE: GetJudgeDataByID method failed", "id", id, "error", err)
		return err
	}

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeDataByID method failed", "id", id, "error", err)
		return err
	}

	err = j.JudgeRepository.DeleteProfile(id)
	if err != nil {
		j.logger.Error("SERVICE: Delete method failed", "error", err)
	}

	j.logger.Info("SERVICE: Successfully deleted judgeView", "id", id)
	return nil
}

func (j JudgeService) GetJudgeDataByID(id uuid.UUID) (*models.Judge, error) {
	judge, err := j.JudgeRepository.GetJudgeDataByID(id)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByID method failed", "id", id, "error", err)
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got user with GetJudgeByID", "id", id)
	return judge, nil
}

func (j JudgeService) GetJudgeDataByProtestID(protestID uuid.UUID) (*models.Judge, error) {
	judge, err := j.JudgeRepository.GetJudgeDataByProtestID(protestID)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByID method failed", "id", protestID, "error", err)
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got user with GetJudgeByID", "id", protestID)
	return judge, nil
}

func (j JudgeService) GetJudgesDataByRatingID(ratingID uuid.UUID) ([]models.Judge, error) {
	judges, err := j.JudgeRepository.GetJudgesDataByRatingID(ratingID)

	if err != nil {
		j.logger.Error("SERVICE: GetJudgesByRatingID method failed", "id", ratingID, "error", err)
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got judges with GetJudgesByRatingID", "id", ratingID)
	return judges, nil
}

func (j JudgeService) GetAllJudges() ([]models.Judge, error) {
	judges, err := j.JudgeRepository.GetAllJudges()

	if err != nil {
		j.logger.Error("SERVICE: GetAllJudges method failed", "error", err)
		return nil, err
	}

	j.logger.Info("SERVICE: Successfully got All Judges")
	return judges, nil
}

func (j JudgeService) UpdateProfile(judgeID uuid.UUID, fio string, login string, password string, role int) (*models.Judge, error) {
	judge, err := j.JudgeRepository.GetJudgeDataByID(judgeID)
	judgeCopy := judge

	if err != nil {
		j.logger.Error("SERVICE: GetJudgeByID method failed", "id", judgeID, "error", err)
		return judge, err
	}

	j.logger.Info("SERVICE: Validating data")
	if !validFIO(fio) {
		j.logger.Error("SERVICE: Invalid fio", "fio", fio)
		return judge, service_errors.InvalidFIO
	}

	if !validLogin(login) {
		j.logger.Error("SERVICE: Invalid login", "login", login)
		return judge, service_errors.InvalidLogin
	}

	if !validRole(role) {
		j.logger.Error("SERVICE: Invalid role", "role", role)
		return judge, service_errors.InvalidRole
	}

	if !validPassword(password) {
		j.logger.Error("SERVICE: Invalid password", "password", password)
		return judge, service_errors.InvalidPassword
	}

	j.logger.Infof("SERVICE: Creating new judgeView: %s", fio)
	hashedPassword, err := j.hash.GetHash(password)
	if err != nil {
		j.logger.Error("SERVICE: Error occurred during password hashing")
		return judge, err
	} else {
		password = hashedPassword
	}

	judge.Role = role
	judge.FIO = fio
	judge.Password = password
	judge.Login = login

	judge, err = j.JudgeRepository.UpdateProfile(judge)
	if err != nil {
		judge = judgeCopy
		j.logger.Error("SERVICE: UpdateJudge method failed", "error", err)
		return judge, err
	}

	j.logger.Info("SERVICE: Successfully updated judgeView coach", "judgeView", judge)
	return judge, nil
}
