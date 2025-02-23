import { renderHook, act } from "@testing-library/react";
import { handleError } from "../errorHandler";
import api from "../api";
import {ProtestComplete, ProtestFormData} from "../../ui/models/protestModel";
import {useCompleteProtest} from "../../ui/controllers/protestControllers/completeProtestController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useCompleteProtest", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestID = "protest789";
    const protestCompleteData: ProtestComplete = {
        resPoints: 10,
        comment: "Протест удовлетворен",
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное завершение протеста", async () => {
        const protestResponse: ProtestFormData = {
            ID: protestID,
            JudgeID: "judge001",
            RatingID: ratingID,
            RaceID: raceID,
            RuleNum: 42,
            ReviewDate: "2025-02-23",
            Status: "Completed",
            Comment: "Протест рассмотрен",
        };

        (api.post as jest.Mock).mockResolvedValue({ data: protestResponse });

        const { result } = renderHook(() => useCompleteProtest());

        await act(async () => {
            await result.current.completeProtest(ratingID, raceID, protestID, protestCompleteData);
        });

        expect(result.current.protestData).toEqual(protestResponse);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBeFalsy();
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/complete`,
            protestCompleteData
        );
    });

    test("ошибка при завершении протеста", async () => {
        const errorMessage = "Ошибка завершения протеста";
        (api.post as jest.Mock).mockRejectedValue(new Error(errorMessage));
        (handleError as jest.Mock).mockImplementation((err, setError) => {
            setError(err.message);
        });

        const { result } = renderHook(() => useCompleteProtest());

        await act(async () => {
            await result.current.completeProtest(ratingID, raceID, protestID, protestCompleteData);
        });

        expect(result.current.error).toBe(errorMessage);
        expect(result.current.protestData).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });

    test("показ индикатора загрузки во время запроса", async () => {
        (api.post as jest.Mock).mockImplementation(() => new Promise(() => {}));

        const { result } = renderHook(() => useCompleteProtest());

        act(() => {
            result.current.completeProtest(ratingID, raceID, protestID, protestCompleteData);
        });

        expect(result.current.loading).toBeTruthy();
    });

    test("ошибка при отсутствии идентификаторов", async () => {
        const { result } = renderHook(() => useCompleteProtest());

        await act(async () => {
            await result.current.completeProtest("", raceID, protestID, protestCompleteData);
        });

        expect(result.current.error).toBe("Недостаточно данных.");
        expect(result.current.protestData).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });
});
