import { renderHook, act } from "@testing-library/react";
import { handleError } from "../errorHandler";
import api from "../api";
import {useAttachProtestMember} from "../../ui/controllers/protestControllers/attachProtestMemberController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useAttachProtestMember", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestID = "protest789";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное добавление участника протеста", async () => {
        const successMessage = "Команда-участник успешно добавлена";
        (api.post as jest.Mock).mockResolvedValue({ data: { message: successMessage } });

        const { result } = renderHook(() =>
            useAttachProtestMember(ratingID, raceID, protestID)
        );

        await act(async () => {
            await result.current.attachProtestMember({ sailNum: 101, role: 2 });
        });

        expect(result.current.successMessage).toBe(successMessage);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBeFalsy();
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members`,
            { sailNum: 101, role: 2 }
        );
    });

    test("ошибка при добавлении участника", async () => {
        const errorMessage = "Ошибка добавления участника";
        (api.post as jest.Mock).mockRejectedValue(new Error(errorMessage));
        (handleError as jest.Mock).mockImplementation((err, setError) => {
            setError(err.message);
        });

        const { result } = renderHook(() =>
            useAttachProtestMember(ratingID, raceID, protestID)
        );

        await act(async () => {
            await result.current.attachProtestMember({ sailNum: 102, role: 1 });
        });

        expect(result.current.error).toBe(errorMessage);
        expect(result.current.successMessage).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });

    test("показ индикатора загрузки во время запроса", async () => {
        (api.post as jest.Mock).mockImplementation(() => new Promise(() => {}));

        const { result } = renderHook(() =>
            useAttachProtestMember(ratingID, raceID, protestID)
        );

        act(() => {
            result.current.attachProtestMember({ sailNum: 103, role: 3 });
        });

        expect(result.current.loading).toBeTruthy();
    });

    test("состояние сбрасывается перед новым запросом", async () => {
        const successMessage = "Команда-участник успешно добавлена";
        (api.post as jest.Mock).mockResolvedValue({ data: { message: successMessage } });

        const { result } = renderHook(() =>
            useAttachProtestMember(ratingID, raceID, protestID)
        );

        await act(async () => {
            await result.current.attachProtestMember({ sailNum: 104, role: 1 });
        });

        expect(result.current.successMessage).toBe(successMessage);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBeFalsy();

        await act(async () => {
            await result.current.attachProtestMember({ sailNum: 105, role: 2 });
        });

        expect(result.current.successMessage).toBe(successMessage);
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });
});
