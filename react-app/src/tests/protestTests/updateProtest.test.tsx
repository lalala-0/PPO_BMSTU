import { render, act, screen } from "@testing-library/react";
import api from "../api";
import { ProtestFormData } from "../../ui/models/protestModel";
import React from "react";
import {useFetchProtests} from "../../ui/controllers/protestControllers/getProtestsController";

// Компонент-обертка для тестирования хука
const FetchProtestsWrapper = ({ ratingID, raceID }: { ratingID: string, raceID: string }) => {
    const { protests, loading, error } = useFetchProtests(ratingID, raceID);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="protests">
                {protests.length > 0 ? JSON.stringify(protests) : "No protests"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useFetchProtests controller", () => {
    const ratingID = "rating123";
    const raceID = "race456";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение протеста", async () => {
        const mockProtests: ProtestFormData[] = [
            { ID: "protest1", JudgeID: "judge1", RatingID: ratingID, RaceID: raceID, RuleNum: 10, Status: "Resolved", ReviewDate: "2025-02-20", Comment: "Valid protest" },
            { ID: "protest2", JudgeID: "judge2", RatingID: ratingID, RaceID: raceID, RuleNum: 5, Status: "Pending", ReviewDate: "2025-02-21", Comment: "Pending protest" }
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockProtests });

        render(<FetchProtestsWrapper ratingID={ratingID} raceID={raceID} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения асинхронной операции
        await act(async () => {
            // В ожидании API-запроса
        });

        // Проверяем конечное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("protests").textContent).toBe(JSON.stringify(mockProtests));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/ratings/${ratingID}/races/${raceID}/protests`);
    });

    test("ошибка при получении протеста", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения протестов"));

        render(<FetchProtestsWrapper ratingID={ratingID} raceID={raceID} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения асинхронной операции
        await act(async () => {
            // В ожидании API-запроса
        });

        // Проверяем конечное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("protests").textContent).toBe("No protests");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения протестов");
    });

    test("проверка повторных вызовов с разными параметрами", async () => {
        const mockProtests: ProtestFormData[] = [
            { ID: "protest1", JudgeID: "judge1", RatingID: ratingID, RaceID: raceID, RuleNum: 10, Status: "Resolved", ReviewDate: "2025-02-20", Comment: "Valid protest" }
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockProtests });

        render(<FetchProtestsWrapper ratingID={ratingID} raceID={raceID} />);

        // Нажимаем кнопку обновления данных протеста дважды
        await act(async () => {
            // В ожидании API-запроса
        });

        expect(api.get).toHaveBeenCalledTimes(1);

        const newRaceID = "race789";
        render(<FetchProtestsWrapper ratingID={ratingID} raceID={newRaceID} />);

        // Проверяем, что API был вызван повторно с новым raceID
        await act(async () => {
            // В ожидании нового API-запроса
        });

        expect(api.get).toHaveBeenCalledTimes(2);
        expect(api.get).toHaveBeenCalledWith(`/ratings/${ratingID}/races/${newRaceID}/protests`);
    });
});
