import { renderHook, act } from "@testing-library/react";
import api from "../api";
import { handleError } from "../errorHandler";
import {CrewFormData} from "../../ui/models/crewModel";
import {useGetProtestMembers} from "../../ui/controllers/protestControllers/getProtestMembersController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useGetProtestMembers", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestID = "protest789";
    const mockMembers: CrewFormData[] = [
        { id: "1", ratingId: "rating123", SailNum: 101, Class: "Laser" },
        { id: "2", ratingId: "rating123", SailNum: 102, Class: "Optimist" },
    ];

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение списка участников", async () => {
        (api.get as jest.Mock).mockResolvedValue({ data: mockMembers });

        const { result } = renderHook(() =>
            useGetProtestMembers(ratingID, raceID, protestID)
        );

        // Проверяем начальное состояние
        expect(result.current.loading).toBe(true);

        // Ждем изменения состояния
        await act(async () => {
            await new Promise(resolve => setTimeout(resolve, 0)); // Ждём асинхронные обновления
        });

        expect(api.get).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members`
        );
        expect(result.current.protestMembers).toEqual(mockMembers);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBe(false);
    });

    test("ошибка при получении списка участников", async () => {
        const error = new Error("Ошибка загрузки");
        (api.get as jest.Mock).mockRejectedValue(error);

        const { result } = renderHook(() =>
            useGetProtestMembers(ratingID, raceID, protestID)
        );

        // Проверяем начальное состояние
        expect(result.current.loading).toBe(true);

        // Ждем изменения состояния
        await act(async () => {
            await new Promise(resolve => setTimeout(resolve, 0)); // Ждём асинхронные обновления
        });

        expect(api.get).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members`
        );
        expect(handleError).toHaveBeenCalledWith(error, expect.any(Function));
        expect(result.current.protestMembers).toEqual([]);
        expect(result.current.loading).toBe(false);
    });
});
