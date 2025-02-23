import { renderHook, act } from "@testing-library/react";
import { handleError } from "../errorHandler";
import api from "../api";
import {ProtestCreate, ProtestFormData} from "../../ui/models/protestModel";
import {useCreateProtest} from "../../ui/controllers/protestControllers/createProtestController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useCreateProtest", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestCreateData: ProtestCreate = {
        judgeId: "1",
        reviewDate: "23.04.2003",
        protestee: 1,
        protestor: 2,
        witnesses: [3, 4],
        ruleNum: 42,
        comment: "Нарушение правил",
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное создание протеста", async () => {
        const protestResponse: ProtestFormData = {
            ID: "protest789",
            JudgeID: "judge001",
            RatingID: ratingID,
            RaceID: raceID,
            RuleNum: 42,
            ReviewDate: "2025-02-23",
            Status: "Pending",
            Comment: "Нарушение правил",
        };

        (api.post as jest.Mock).mockResolvedValue({ data: protestResponse });

        const { result } = renderHook(() => useCreateProtest());

        await act(async () => {
            await result.current.createProtest(ratingID, raceID, protestCreateData);
        });

        expect(result.current.newProtest).toEqual(protestResponse);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBeFalsy();
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests`,
            protestCreateData
        );
    });

    test("ошибка при создании протеста", async () => {
        const errorMessage = "Ошибка создания протеста";
        (api.post as jest.Mock).mockRejectedValue(new Error(errorMessage));
        (handleError as jest.Mock).mockImplementation((err, setError) => {
            setError(err.message);
        });

        const { result } = renderHook(() => useCreateProtest());

        await act(async () => {
            await result.current.createProtest(ratingID, raceID, protestCreateData);
        });

        expect(result.current.error).toBe(errorMessage);
        expect(result.current.newProtest).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });

    test("показ индикатора загрузки во время запроса", async () => {
        (api.post as jest.Mock).mockImplementation(() => new Promise(() => {}));

        const { result } = renderHook(() => useCreateProtest());

        act(() => {
            result.current.createProtest(ratingID, raceID, protestCreateData);
        });

        expect(result.current.loading).toBeTruthy();
    });
});
