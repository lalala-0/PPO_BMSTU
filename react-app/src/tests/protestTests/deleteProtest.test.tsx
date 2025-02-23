import { renderHook, act } from "@testing-library/react";
import api from "../api";
import {useDeleteProtestController} from "../../ui/controllers/protestControllers/deleteProtestController";

jest.mock("../api");

describe("useDeleteProtestController", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestID = "protest789";

    beforeEach(() => {
        jest.clearAllMocks();
        jest.spyOn(window, "alert").mockImplementation(() => {});
    });

    test("успешное удаление протеста", async () => {
        (api.delete as jest.Mock).mockResolvedValue({});

        const { result } = renderHook(() => useDeleteProtestController());

        await act(async () => {
            await result.current.handleDelete(ratingID, raceID, protestID);
        });

        expect(api.delete).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}`
        );
        expect(window.alert).toHaveBeenCalledWith("Протест успешно удалён");
    });

    test("ошибка при удалении протеста", async () => {
        (api.delete as jest.Mock).mockRejectedValue(new Error("Ошибка"));

        const { result } = renderHook(() => useDeleteProtestController());

        await act(async () => {
            await result.current.handleDelete(ratingID, raceID, protestID);
        });

        expect(api.delete).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}`
        );
        expect(window.alert).toHaveBeenCalledWith("Ошибка при удалении протеста");
    });

    test("проверка некорректных параметров", async () => {
        const { result } = renderHook(() => useDeleteProtestController());

        await act(async () => {
            await result.current.handleDelete("", raceID, protestID);
        });

        expect(api.delete).not.toHaveBeenCalled();
        expect(window.alert).toHaveBeenCalledWith(
            "Некорректные параметры для удаления протеста"
        );
    });
});
