import { render, act, screen } from "@testing-library/react";
import { useUpdateJudge } from "../../ui/controllers/judgeControllers/updateJudgeController";
import api from "../api";
import {JudgeFormData} from "../../ui/models/judgeModel";
import React from "react";

// Компонент-обертка для тестирования хука
const UpdateJudgeWrapper = ({ judgeID }: { judgeID: string }) => {
    const { updateJudge, loading, error, updatedJudge } = useUpdateJudge(judgeID);

    return (
        <div>
            <button
                onClick={() =>
                    updateJudge({
                        id: judgeID,
                        fio: "Иван Иванов",
                        login: "ivanov123",
                        role: 1,
                        post: "Судья",
                    })
                }
                data-testid="update-button"
            >
                Update Judge
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="updated-judge">
                {updatedJudge ? JSON.stringify(updatedJudge) : "No judge updated"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useUpdateJudge controller", () => {
    const judgeID = "judge123";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление судьи", async () => {
        const mockUpdatedJudge: JudgeFormData = {
            id: judgeID,
            fio: "Иван Иванов",
            login: "ivanov123",
            role: "1",
            post: "Судья",
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedJudge });

        render(<UpdateJudgeWrapper judgeID={judgeID} />);

        // Нажимаем кнопку для обновления данных судьи
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после успешного обновления
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("updated-judge").textContent).toBe(
            JSON.stringify(mockUpdatedJudge)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.put).toHaveBeenCalledWith(`/judges/${judgeID}`, {
            id: judgeID,
            fio: "Иван Иванов",
            login: "ivanov123",
            role: 1,
            post: "Судья",
        });
    });

    test("ошибка при обновлении судьи", async () => {
        // Мокаем ошибку API
        (api.put as jest.Mock).mockRejectedValue(new Error("Ошибка обновления судьи"));

        render(<UpdateJudgeWrapper judgeID={judgeID} />);

        // Нажимаем кнопку для обновления данных судьи
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("updated-judge").textContent).toBe("No judge updated");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка обновления судьи");
    });

    test("проверка обновления при повторном вызове", async () => {
        const mockUpdatedJudge: JudgeFormData = {
            id: judgeID,
            fio: "Иван Иванов",
            login: "ivanov123",
            role: "1",
            post: "Судья",
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedJudge });

        render(<UpdateJudgeWrapper judgeID={judgeID} />);

        // Нажимаем кнопку для обновления данных судьи дважды
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
            updateButton.click(); // Дважды для проверки
        });

        expect(api.put).toHaveBeenCalledTimes(2); // Запрос должен быть только один
    });
});
