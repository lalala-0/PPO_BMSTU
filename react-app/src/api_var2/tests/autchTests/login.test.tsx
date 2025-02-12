import api from "../../controllers/api";
import {login} from "../../controllers/autchControllers/login";

jest.mock("../../controllers/api");

describe("login controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
        sessionStorage.clear();
        window.location = { reload: jest.fn() } as any;
    });

    test("успешный вход: сохраняет токен, judgeID и перезагружает страницу", async () => {
        const mockToken = "mockToken123";
        const mockJudgeID = "judge123";

        (api.post as jest.Mock).mockResolvedValue({
            data: { token: mockToken, judge: { ID: mockJudgeID } },
        });

        const result = await login("testUser", "testPass");

        expect(api.post).toHaveBeenCalledWith("/login", { login: "testUser", password: "testPass" });
        expect(sessionStorage.getItem("token")).toBe(mockToken);
        expect(sessionStorage.getItem("judgeID")).toBe(mockJudgeID);
        expect(result).toBe(mockToken);
    });

    test("ошибка при неудачном входе: выбрасывает ошибку", async () => {
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка авторизации"));

        await expect(login("testUser", "wrongPass")).rejects.toThrow("Ошибка авторизации");

        expect(api.post).toHaveBeenCalledWith("/login", { login: "testUser", password: "wrongPass" });
        expect(sessionStorage.getItem("token")).toBeNull();
        expect(sessionStorage.getItem("judgeID")).toBeNull();
    });
});
