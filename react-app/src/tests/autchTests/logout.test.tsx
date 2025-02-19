import { logout } from "../../ui/controllers/autchControllers/logout";  // Путь к функции logout
jest.mock("../api"); // Мокаем api.get

describe("logout function", () => {
    beforeEach(() => {
        // Мокаем sessionStorage
        sessionStorage.setItem("token", "mockToken123");  // Устанавливаем токен в sessionStorage
    });

    afterEach(() => {
        jest.clearAllMocks();  // Очищаем моки после каждого теста
    });

    test("удаляет токен и перезагружает страницу", () => {
        // Вызываем logout
        logout();

        // Проверяем, что токен был удален
        expect(sessionStorage.getItem("token")).toBeNull();

    });
});
