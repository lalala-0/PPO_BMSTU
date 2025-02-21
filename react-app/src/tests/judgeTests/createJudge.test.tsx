import { render, act, screen } from "@testing-library/react";
import { useCreateJudge } from "../../ui/controllers/judgeControllers/createJudgeController";
import api from "../api";
import {JudgeInput} from "../../ui/models/judgeModel";
import React from "react";

// Компонент-обертка для тестирования хука
const CreateJudgeWrapper = ({ judgeData }: { judgeData: JudgeInput }) => {
    const { createJudge, loading, error, data } = useCreateJudge();

    return (
        <div>
            <button
                onClick={() => createJudge(judgeData)}
                data-testid="create-button"
            >
                Create Judge
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="data">{data ? JSON.stringify(data) : "No data"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useCreateJudge controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное создание судьи", async () => {
        const mockInput: JudgeInput = {
            id: "judge123",
            fio: "Иван Иванов",
            login: "ivanov123",
            password: "securePassword",
            role: 1,
            post: "Судья",
        };

        const mockResponse = { ...mockInput };

        // Мокаем успешный ответ API
        (api.post as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(
            <CreateJudgeWrapper judgeData={mockInput} />
        );

        // Нажимаем кнопку, чтобы инициировать создание судьи
        const createButton = screen.getByTestId("create-button");
        await act(async () => {
            createButton.click();
        });

        // Проверяем состояние после успешного создания судьи
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.post).toHaveBeenCalledWith(`/judges`, mockInput);
    });

    test("ошибка при создании судьи", async () => {
        const mockInput: JudgeInput = {
            id: "judge123",
            fio: "Иван Иванов",
            login: "ivanov123",
            password: "securePassword",
            role: 1,
            post: "Судья",
        };

        // Мокаем ошибку API
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка создания судьи"));

        render(
            <CreateJudgeWrapper judgeData={mockInput} />
        );

        // Нажимаем кнопку, чтобы инициировать создание судьи
        const createButton = screen.getByTestId("create-button");
        await act(async () => {
            createButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe("No data");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка создания судьи");
    });
});
