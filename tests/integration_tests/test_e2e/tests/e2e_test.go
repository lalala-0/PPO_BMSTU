package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (suite *e2eTestSuite) TestJudgeE2E() {

	// 1. Создаем рейтинг
	reqBody := strings.NewReader(`{
	  "blowout_cnt": 1,
	  "class": 1,
	  "name": "Test Rating"
	}`)
	req, _ := http.NewRequest("POST", "/api/ratings", reqBody) // Обратите внимание на изменение пути
	req.Header.Set("accept", "application/json")               // Добавлен заголовок 'accept'
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)
	assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание рейтинга не удалось")

	//// 2. Создаем трех участников
	//participantNames := []string{"Alice", "Bob", "Charlie"}
	//for _, name := range participantNames {
	//	reqBody = strings.NewReader(`{"name": "` + name + `", "age": 30}`)
	//	req, _ = http.NewRequest("POST", "/api/participants/", reqBody)
	//	req.Header.Set("Content-Type", "application/json")
	//	resp = httptest.NewRecorder()
	//	suite.router.ServeHTTP(resp, req)
	//	assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание участника "+name+" не удалось")
	//}
	//
	//// 3. Создаем три команды
	//teamNames := []string{"Team 1", "Team 2", "Team 3"}
	//for _, team := range teamNames {
	//	reqBody = strings.NewReader(`{"name": "` + team + `", "sail_number": "1234"}`)
	//	req, _ = http.NewRequest("POST", "/api/ratings/1/crews", reqBody)
	//	req.Header.Set("Content-Type", "application/json")
	//	resp = httptest.NewRecorder()
	//	suite.router.ServeHTTP(resp, req)
	//	assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание команды "+team+" не удалось")
	//}
	//
	//// 4. Добавляем по одному участнику в каждую команду
	//for i := 1; i <= 3; i++ {
	//	reqBody = strings.NewReader(`{"participantID": ` + uuid.New().String() + `}`)
	//	req, _ = http.NewRequest("POST", "/api/ratings/1/crews/"+string(i)+"/members", reqBody)
	//	req.Header.Set("Content-Type", "application/json")
	//	resp = httptest.NewRecorder()
	//	suite.router.ServeHTTP(resp, req)
	//	assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Добавление участника в команду не удалось")
	//}
	//
	//// 5. Создаем гонку
	//reqBody = strings.NewReader(`{"name": "Race 1", "location": "Bay Area"}`)
	//req, _ = http.NewRequest("POST", "/api/ratings/1/races", reqBody)
	//req.Header.Set("Content-Type", "application/json")
	//resp = httptest.NewRecorder()
	//suite.router.ServeHTTP(resp, req)
	//assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание гонки не удалось")
	//
	//// 6. Проводим стартовую процедуру гонки
	//req, _ = http.NewRequest("POST", "/api/ratings/1/races/1/start", nil)
	//resp = httptest.NewRecorder()
	//suite.router.ServeHTTP(resp, req)
	//assert.Equal(suite.T(), http.StatusOK, resp.Code, "Стартовая процедура не удалась")
	//
	//// 7. Проводим финишную процедуру гонки
	//req, _ = http.NewRequest("POST", "/api/ratings/1/races/1/finish", nil)
	//resp = httptest.NewRecorder()
	//suite.router.ServeHTTP(resp, req)
	//assert.Equal(suite.T(), http.StatusOK, resp.Code, "Финишная процедура не удалась")
	//
	//// 8. Создаем протест
	//reqBody = strings.NewReader(`{"reason": "Foul play", "details": "A team broke the rules"}`)
	//req, _ = http.NewRequest("POST", "/api/ratings/1/races/1/protests", reqBody)
	//req.Header.Set("Content-Type", "application/json")
	//resp = httptest.NewRecorder()
	//suite.router.ServeHTTP(resp, req)
	//assert.Equal(suite.T(), http.StatusCreated, resp.Code, "Создание протеста не удалось")
	//
	//// 9. Завершаем рассмотрение протеста
	//req, _ = http.NewRequest("PATCH", "/api/ratings/1/races/1/protests/1/complete", nil)
	//resp = httptest.NewRecorder()
	//suite.router.ServeHTTP(resp, req)
	//assert.Equal(suite.T(), http.StatusOK, resp.Code, "Завершение рассмотрения протеста не удалось")
}
