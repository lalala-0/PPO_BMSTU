import { renderHook, act } from "@testing-library/react";
import api from "../api";
import { handleError } from "../errorHandler";
import {ProtestFormData} from "../../ui/models/protestModel";
import {useFetchProtests} from "../../ui/controllers/protestControllers/getProtestsController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useFetchProtests", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const mockProtests: ProtestFormData[] = [
        {
            ID: "1",
            JudgeID: "judge1",
            RatingID: "rating123",
            RaceID: "race456",
            RuleNum: 10,
            ReviewDate: "2025-02-20",
            Status: "Pending",
            Comment: "Test protest 1",
        },
        {
            ID: "2",
            JudgeID: "judge2",
            RatingID: "rating123",
            RaceID: "race456",
            RuleNum: 12,
            ReviewDate: "2025-02-21",
            Status: "Resolved",
            Comment: "Test protest 2",
        },
    ];

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение списка протестов", async () => {
        (api.get as jest.Mock).mockResolvedValue({ data: mockProtests });

        const { result } = renderHook(() =>
            useFetchProtests(ratingID, raceID)
        );

        // Проверяем начальное состояние
        expect(result.current.loading).toBe(true);

        // Ждем изменения состояния
        await act(async () => {
            await new Promise(resolve => setTimeout(resolve, 0)); // Ждём асинхронные обновления
        });

        expect(api.get).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests`
        );
        expect(result.current.protests).toEqual(mockProtests);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBe(false);
    });

    test("ошибка при получении списка протестов", async () => {
        const error = new Error("Ошибка загрузки");
        (api.get as jest.Mock).mockRejectedValue(error);

        const { result } = renderHook(() =>
            useFetchProtests(ratingID, raceID)
        );

        // Проверяем начальное состояние
        expect(result.current.loading).toBe(true);

        // Ждем изменения состояния
        await act(async () => {
            await new Promise(resolve => setTimeout(resolve, 0)); // Ждём асинхронные обновления
        });

        expect(api.get).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests`
        );
        expect(handleError).toHaveBeenCalledWith(error, expect.any(Function));
        expect(result.current.protests).toEqual([]);
        expect(result.current.loading).toBe(false);
    });
});
