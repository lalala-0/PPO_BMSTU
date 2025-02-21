import { render, act, screen } from "@testing-library/react";
import { useGetJudge } from "../../ui/controllers/judgeControllers/getJudgeController";
import api from "../api";
import {JudgeFormData} from "../../ui/models/judgeModel";
import React from "react";

// Компонент-обертка для тестирования хука
const GetJudgeWrapper = ({ judgeID }: { judgeID: string }) => {
    const { judge, loading, error, fetchJudge } = useGetJudge(judgeID);

    return (
        <div>
            <button onClick={() => fetchJudge()} data-testid="fetch-button">
                Fetch Judge
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="judge">{judge ? JSON.stringify(judge) : "No judge"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useGetJudge controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешная загрузка данных судьи", async () => {
        const mockJudgeID = "judge123";
        const mockResponse: JudgeFormData = {
            id: "judge123",
            fio: "Иван Иванов",
            login: "ivanov123",
            role: "main",
            post: "Судья",
        };

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(<GetJudgeWrapper judgeID={mockJudgeID} />);

        // Нажимаем кнопку для загрузки данных
        const fetchButton = screen.getByTestId("fetch-button");
        await act(async () => {
            fetchButton.click();
        });

        // Проверяем состояние после успешной загрузки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("judge").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/judges/${mockJudgeID}`);
    });

    test("ошибка при загрузке данных судьи", async () => {
        const mockJudgeID = "judge123";

        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка загрузки судьи"));

        render(<GetJudgeWrapper judgeID={mockJudgeID} />);

        // Нажимаем кнопку для загрузки данных
        const fetchButton = screen.getByTestId("fetch-button");
        await act(async () => {
            fetchButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("judge").textContent).toBe("No judge");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка загрузки судьи");
    });

    test("загрузка данных при изменении judgeID", async () => {
        const mockJudgeID = "judge123";
        const mockResponse: JudgeFormData = {
            id: "judge123",
            fio: "Иван Иванов",
            login: "ivanov123",
            role: "main",
            post: "Судья",
        };

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        const { rerender } = render(<GetJudgeWrapper judgeID={mockJudgeID} />);

        // Проверяем, что данные загружаются при изменении judgeID
        await act(async () => {
            rerender(<GetJudgeWrapper judgeID="newJudgeID" />);
        });

        // Проверяем состояние после успешной загрузки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("judge").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/judges/newJudgeID`);
    });
});
