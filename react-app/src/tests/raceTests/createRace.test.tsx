import {waitFor, renderHook} from "@testing-library/react";
import api from "../api";
import { handleError } from "../errorHandler";
import {useCreateRace} from "../../ui/controllers/raceControllers/createRaceController";
jest.mock("../api");
jest.mock("../errorHandler");

describe("useCreateRace", () => {
    it("должен инициализироваться с дефолтным состоянием", () => {
        const { result } = renderHook(() => useCreateRace());

        expect(result.current.input).toEqual({
            date: "",
            number: 0,
            class: 0,
        });
        expect(result.current.success).toBeNull();
        expect(result.current.loading).toBeFalsy();
    });

    it("должен корректно обрабатывать успешную отправку", async () => {
        const mockPost = jest.fn().mockResolvedValue({ data: {} });
        api.post = mockPost;
        const { result } = renderHook(() => useCreateRace());

        await result.current.handleSubmit("12345");

        await waitFor(() => expect(result.current.success).toBe("Гонка успешно создана"));
        expect(mockPost).toHaveBeenCalledWith("/ratings/12345/races", result.current.input);
        expect(result.current.input).toEqual({
            date: "",
            number: 0,
            class: 0,
        });
    });

    it("должен обрабатывать ошибку при отправке", async () => {
        const mockPost = jest.fn().mockRejectedValue(new Error("Network error"));
        api.post = mockPost;
        const { result } = renderHook(() => useCreateRace());

        await result.current.handleSubmit("12345");

        await waitFor(() => expect(result.current.loading).toBeFalsy());
        expect(handleError).toHaveBeenCalledWith(new Error("Network error"), expect.any(Function));
    });

});
