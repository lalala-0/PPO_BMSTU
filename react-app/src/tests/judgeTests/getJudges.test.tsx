import { render, act, screen } from "@testing-library/react";
import { useGetJudges } from "../../ui/controllers/judgeControllers/getJudgesController";
import api from "../api";
import {JudgeFormData} from "../../ui/models/judgeModel";

// Компонент-обертка для тестирования хука
const GetJudgesWrapper = () => {
    const { judges, loading, error, fetchJudges } = useGetJudges();

    return (
        <div>
            <button onClick={() => fetchJudges()} data-testid="fetch-button">
                Fetch Judges
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="judges">
                {judges.length > 0 ? JSON.stringify(judges) : "No judges"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useGetJudges controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешная загрузка списка судей", async () => {
        const mockResponse: JudgeFormData[] = [
            {
                id: "judge123",
                fio: "Иван Иванов",
                login: "ivanov123",
                role: "1",
                post: "Судья",
            },
            {
                id: "judge456",
                fio: "Мария Петрова",
                login: "petrova456",
                role: "2",
                post: "Судья",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(<GetJudgesWrapper />);

        // Нажимаем кнопку для загрузки данных
        const fetchButton = screen.getByTestId("fetch-button");
        await act(async () => {
            fetchButton.click();
        });

        // Проверяем состояние после успешной загрузки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("judges").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith("/judges");
    });

    test("ошибка при загрузке списка судей", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка загрузки судей"));

        render(<GetJudgesWrapper />);

        // Нажимаем кнопку для загрузки данных
        const fetchButton = screen.getByTestId("fetch-button");
        await act(async () => {
            fetchButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("judges").textContent).toBe("No judges");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка загрузки судей");
    });

    test("загрузка данных только один раз", async () => {
        const mockResponse: JudgeFormData[] = [
            {
                id: "judge123",
                fio: "Иван Иванов",
                login: "ivanov123",
                role: "1",
                post: "Судья",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(<GetJudgesWrapper />);

        // Проверяем, что API вызывается только один раз
        expect(api.get).toHaveBeenCalledTimes(1);

        // Нажимаем кнопку для загрузки данных
        const fetchButton = screen.getByTestId("fetch-button");
        await act(async () => {
            fetchButton.click();
        });

        // Проверяем, что API не вызывается повторно
        expect(api.get).toHaveBeenCalledTimes(2); // Запрос не должен повторяться
    });
});
