package controllers

import (
	"PPO_BMSTU/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (s *Services) judgeProfile(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	if judge.Role == models.NotMainJudge {
		c.HTML(200, "judge-profile", gin.H{
			"title": "Профиль судьи",
			"judge": judge,
		})
		return
	}
}

func (s *Services) judgeMainMenu(judge *models.Judge) gin.H {
	var result = gin.H{
		"title": "Панель судьи",
		"judge": judge,
	}
	//
	//ordersWithoutJudge, _ := s.Services.OrderService.Filter(map[string]string{"judge_id": "null", "status": "1,2"})
	//ordersInProgress, _ := s.Services.OrderService.Filter(map[string]string{"judge_id": "not null", "status": "1,2"})
	//
	//ordersWithoutJudgeData := make([]orderData, len(ordersWithoutJudge))
	//for i, o := range ordersWithoutJudge {
	//	user, _ := s.Services.UserService.GetUserByID(o.UserID)
	//	ordersWithoutJudgeData[i] = orderData{
	//		ID:           o.ID,
	//		User:         user,
	//		Status:       models.OrderStatuses[o.Status],
	//		Address:      o.Address,
	//		CreationDate: o.CreationDate.Format("2006-01-02 15:04:05"),
	//		Deadline:     o.Deadline.Format("2006-01-02"),
	//		Rate:         o.Rate,
	//	}
	//}
	//
	//ordersInProgressData := make([]orderData, len(ordersInProgress))
	//for i, o := range ordersInProgress {
	//	user, _ := s.Services.UserService.GetUserByID(o.UserID)
	//	ordersInProgressData[i] = orderData{
	//		ID:           o.ID,
	//		User:         user,
	//		Status:       models.OrderStatuses[o.Status],
	//		Address:      o.Address,
	//		CreationDate: o.CreationDate.Format("2006-01-02 15:04:05"),
	//		Deadline:     o.Deadline.Format("2006-01-02"),
	//		Rate:         o.Rate,
	//	}
	//}
	//
	//result["ordersWithoutJudge"] = ordersWithoutJudgeData
	//result["ordersInProgress"] = ordersInProgressData

	return result
}

func (s *Services) mainJudgeMainMenu(judge *models.Judge) gin.H {
	var result = gin.H{
		"title": "Панель главного судьи",
		"judge": judge,
	}

	//params := map[string]string{
	//	"status":   "1,2",
	//	"judge_id": judge.ID.String(),
	//}
	//inProgressOrders, _ := s.Services.OrderService.Filter(params)
	//
	//inProgressOrdersData := make([]orderData, len(inProgressOrders))
	//for i, o := range inProgressOrders {
	//	user, _ := s.Services.UserService.GetUserByID(o.UserID)
	//	inProgressOrdersData[i] = orderData{
	//		ID:           o.ID,
	//		User:         user,
	//		Status:       models.OrderStatuses[o.Status],
	//		Address:      o.Address,
	//		CreationDate: o.CreationDate.Format("2006-01-02 15:04:05"),
	//		Deadline:     o.Deadline.Format("2006-01-02"),
	//		Rate:         o.Rate,
	//	}
	//}
	//
	//result["ordersInProgress"] = inProgressOrdersData
	return result
}

func (s *Services) menu(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	if judge.Role == models.MainJudge {
		c.HTML(200, "adminDashboard", s.mainJudgeMainMenu(judge))
		return
	}

	c.HTML(200, "masterDashboard", s.judgeMainMenu(judge))
}

// GET_ALL_JUDGES

type judgeData struct {
	ID    uuid.UUID
	FIO   string
	Role  int
	Post  string
	Login string
}

func (s *Services) judgesDirectory(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	judges, _ := s.Services.JudgeService.GetAllJudges()

	judgesData := make([]judgeData, len(judges))
	for i, w := range judges {
		judgesData[i] = judgeData{
			ID:    w.ID,
			FIO:   w.FIO,
			Role:  w.Role,
			Post:  w.Post,
			Login: w.Login,
		}
	}

	if judge.Role == models.MainJudge {
		c.HTML(200, "judgesDirectory", gin.H{
			"title":  "Список судей",
			"judge":  judge,
			"judges": judgesData,
		})
		return
	}

	c.HTML(403, "judgesDirectory", gin.H{"title": "Список судей", "error": "Доступ запрещен!"})
}

// UPDATE

type editJudgeData struct {
	FIO      string `form:"fio"`
	Login    string `form:"login"`
	Password string `form:"password"`
	Post     string `form:"post"`
	Role     int    `form:"role"`
}

func (s *Services) editJudgeGet(c *gin.Context) {
	authJudge := s.authenticatedJudge(c)
	judgeID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.HTML(400, "editJudge", gin.H{
			"title": "Редактировать профиль",
			"judge": authJudge,
			"error": "Неверный идентификатор судьи",
		})
		return
	}

	editedJudge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "editJudge", gin.H{
			"title": "Редактировать профиль",
			"judge": authJudge,
			"error": "Судья не найден",
		})
		return
	}

	c.HTML(200, "editJudge", gin.H{
		"title": "Редактировать профиль",
		"judge": authJudge,
		"formData": editJudgeData{
			FIO:      editedJudge.FIO,
			Login:    editedJudge.Login,
			Password: editedJudge.Password,
			Post:     editedJudge.Post,
			Role:     editedJudge.Role,
		},
	})
}

func (s *Services) editJudgePost(c *gin.Context) {
	authJudge := s.authenticatedJudge(c)
	judgeID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.HTML(400, "editJudge", gin.H{
			"title": "Редактировать профиль",
			"judge": authJudge,
			"error": "Неверный идентификатор cудьи",
		})
		return
	}

	editedJudge, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	if err != nil {
		c.HTML(400, "editJudge", gin.H{
			"title": "Редактировать профиль",
			"judge": authJudge,
			"error": "Судья4 не найден",
		})
		return
	}

	var data editJudgeData
	err = c.Bind(&data)
	if err != nil {
		c.HTML(400, "editJudge", gin.H{
			"title":    "Редактировать профиль",
			"judge":    authJudge,
			"error":    err.Error(),
			"formData": data,
		})
		return
	}

	_, updateErr := s.Services.JudgeService.UpdateProfile(
		editedJudge.ID,
		data.FIO,
		data.Login,
		data.Password,
		data.Role,
	)

	if updateErr != nil {
		c.HTML(400, "editJudge", gin.H{
			"title": "Редактировать профиль",
			"judge": authJudge,
			"error": updateErr.Error(),
		})
		return
	}

	c.Redirect(302, "/judge/"+judgeID.String())
}

// CREATE

func (s *Services) createJudgeGet(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	if judge.Role == models.MainJudge {
		c.HTML(200, "createJudge", gin.H{
			"title": "Добавление судьи",
			"judge": judge,
		})
		return
	}

	c.HTML(403, "createJudge", gin.H{"title": "Добавление судьи", "error": "Доступ запрещен!"})
}

type createJudgeFormData struct {
	FIO      string `form:"fio" binding:"required"`
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
	Role     string `form:"role" binding:"required"`
	Post     string `form:"post" binding:"required"`
}

func (s *Services) createJudgePost(c *gin.Context) {
	judge := s.authenticatedJudge(c)

	if judge.Role != models.MainJudge {
		c.HTML(403, "createJudge", gin.H{"title": "Добавление судьи", "error": "Доступ запрещен!"})
		return
	}

	var data createJudgeFormData
	if err := c.Bind(&data); err != nil {
		c.HTML(http.StatusBadRequest, "createJudge", gin.H{
			"title": "Добавление судьи",
			"error": err.Error(),
		})
		return
	}

	role, _ := strconv.Atoi(data.Role)

	newJudge := models.Judge{
		FIO:      data.FIO,
		Login:    data.Login,
		Post:     data.Post,
		Role:     role,
		Password: data.Password,
	}

	_, err := s.Services.JudgeService.CreateProfile(newJudge.ID, newJudge.FIO, newJudge.Login, newJudge.Password, newJudge.Role, newJudge.Post)
	if err != nil {
		c.HTML(http.StatusBadRequest, "createJudge", gin.H{
			"title": "Добавление судьи",
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/judge/directory")
}

/////////////////////

func (s *Services) judgeDetails(c *gin.Context) {
	//judge := s.authenticatedJudge(c)
	//
	//if judge.Role != models.ManagerRole {
	//	c.HTML(403, "judgeDetails", gin.H{"title": "Информация об исполнителе", "error": "Доступ запрещен!"})
	//	return
	//}
	//
	//judgeID, err := uuid.Parse(c.Param("id"))
	//if err != nil {
	//	c.HTML(http.StatusBadRequest, "judgeDetails", gin.H{
	//		"title": "Информация об исполнителе",
	//		"error": "Неверный идентификатор исполнителя",
	//	})
	//	return
	//}
	//
	//judgeDetails, err := s.Services.JudgeService.GetJudgeDataByID(judgeID)
	//if err != nil {
	//	c.HTML(http.StatusBadRequest, "judgeDetails", gin.H{
	//		"title": "Информация об исполнителе",
	//		"error": "Исполнитель не найден",
	//	})
	//	return
	//}
	//
	//params := map[string]string{
	//	"judge_id": judgeID.String(),
	//	"status":   "1,2",
	//}
	//inProgressOrders, _ := s.Services.OrderService.Filter(params)
	//
	//inProgressOrdersData := make([]orderData, len(inProgressOrders))
	//for i, o := range inProgressOrders {
	//	user, _ := s.Services.UserService.GetUserByID(o.UserID)
	//	inProgressOrdersData[i] = orderData{
	//		ID:           o.ID,
	//		User:         user,
	//		Status:       models.OrderStatuses[o.Status],
	//		Address:      o.Address,
	//		CreationDate: o.CreationDate.Format("2006-01-02 15:04:05"),
	//		Deadline:     o.Deadline.Format("2006-01-02"),
	//		Rate:         o.Rate,
	//	}
	//}
	//
	//params["status"] = "3"
	//completedOrders, _ := s.Services.OrderService.Filter(params)
	//
	//avgRate, _ := s.Services.JudgeService.GetAverageOrderRate(judgeDetails)
	//
	//c.HTML(200, "judgeDetails", gin.H{
	//	"judge":            judge,
	//	"title":            "Информация об исполнителе",
	//	"judgeDetails":     judgeDetails,
	//	"inProgressOrders": inProgressOrdersData,
	//	"completedOrders":  completedOrders,
	//	"avgRate":          avgRate,
	//})
}
