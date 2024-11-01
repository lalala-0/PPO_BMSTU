package tests

import (
	"PPO_BMSTU/server/api/modelsViewApi"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (suite *e2eTestSuite) TestJudgeE2E() {
	// Создаем рейтинг
	reqBody := strings.NewReader(`{
	  "blowout_cnt": 1,
	  "class": 1,
	  "name": "Test Rating"
	}`)
	req, _ := http.NewRequest("POST", "/api/ratings/", reqBody) // Обратите внимание на изменение пути
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)
	assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание рейтинга не удалось")

	// Получаем информацию о созданном рейтинга
	req, _ = http.NewRequest("GET", "/api/ratings/", nil) // nil, так как GET запрос не имеет тела
	req.Header.Set("accept", "application/json")          // Устанавливаем заголовок 'accept'
	resp = httptest.NewRecorder()                         // Создаем новый recorder для записи ответа
	suite.router.ServeHTTP(resp, req)                     // Отправляем запрос через роутер
	assert.Equal(suite.T(), http.StatusOK, resp.Code, "Получение рейтингов не удалось")
	var ratings []modelsViewApi.RatingFormData // Предположим, у вас есть структура Rating
	err := json.Unmarshal(resp.Body.Bytes(), &ratings)
	assert.NoError(suite.T(), err, "Ошибка при разборе тела ответа")

	// Создаем трех участников
	participantNames := []string{"Test fio1", "Test fio2", "Test fio3"}
	for _, name := range participantNames {
		reqBody = strings.NewReader(`{
			"birthday": "2003-04-23",
			"category": 1,
			"coach": "Test coach",
			"fio": "` + name + `",
			"gender": 1,
			"id": "` + uuid.New().String() + `"
   		}`)
		req, _ = http.NewRequest("POST", "/api/participants/", reqBody) // Изменил путь на "/api/participants"
		req.Header.Set("accept", "application/json")                    // Установил заголовок 'accept'
		req.Header.Set("Content-Type", "application/json")
		resp = httptest.NewRecorder()
		suite.router.ServeHTTP(resp, req)
		assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание участника "+name+" не удалось")
	}

	// Получаем информацию о созданных участниках
	req, _ = http.NewRequest("GET", "/api/participants/", nil) // nil, так как GET запрос не имеет тела
	req.Header.Set("accept", "application/json")               // Устанавливаем заголовок 'accept'
	resp = httptest.NewRecorder()                              // Создаем новый recorder для записи ответа
	suite.router.ServeHTTP(resp, req)                          // Отправляем запрос через роутер
	assert.Equal(suite.T(), http.StatusOK, resp.Code, "Получение участников не удалось")
	var participants []modelsViewApi.ParticipantFormData // Предположим, у вас есть структура Rating
	err = json.Unmarshal(resp.Body.Bytes(), &participants)
	assert.NoError(suite.T(), err, "Ошибка при разборе тела ответа")

	// Создаем три команды
	for _ = range 3 {
		reqBody = strings.NewReader(`{
        	"sailNum": 1
    	}`)
		req, _ = http.NewRequest("POST", "/api/ratings/"+ratings[0].ID.String()+"/crews/", reqBody)
		req.Header.Set("Content-Type", "application/json")
		resp = httptest.NewRecorder()
		suite.router.ServeHTTP(resp, req)
		assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание команды не удалось")
	}

	// Получаем список созданных команд
	req, _ = http.NewRequest("GET", "/api/ratings/"+ratings[0].ID.String()+"/crews/", nil) // nil, так как GET запрос не имеет тела
	req.Header.Set("accept", "application/json")                                           // Устанавливаем заголовок 'accept'
	resp = httptest.NewRecorder()                                                          // Создаем новый recorder для записи ответа
	suite.router.ServeHTTP(resp, req)                                                      // Отправляем запрос через роутер
	assert.Equal(suite.T(), http.StatusOK, resp.Code, "Получение команд не удалось")
	var crews []modelsViewApi.CrewFormData // Предположим, у вас есть структура Rating
	err = json.Unmarshal(resp.Body.Bytes(), &crews)
	assert.NoError(suite.T(), err, "Ошибка при разборе тела ответа")

	// Добавляем по одному участнику в каждую команду
	for i := range 3 {
		reqBody = strings.NewReader(`{
		  "helmsman": 1,
		  "participantID": "` + participants[i].ID.String() + `"
		}`)
		req, _ = http.NewRequest("POST", "/api/ratings/"+ratings[0].ID.String()+"/crews/"+crews[i].ID.String()+"/members", reqBody)
		req.Header.Set("Content-Type", "application/json")
		resp = httptest.NewRecorder()
		suite.router.ServeHTTP(resp, req)
		assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Добавить участника в команду не удалось")
	}

}
