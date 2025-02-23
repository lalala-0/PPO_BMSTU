import { render, act, screen } from "@testing-library/react";
import api from "../api";
import React from "react";
import {useStartProcedure} from "../../ui/controllers/raceControllers/startProcedureController";

jest.mock("../errorHandler", () => ({
    handleError: jest.fn(),
}));

const StartProcedureWrapper = ({ ratingID, raceID }: { ratingID: string; raceID: string }) => {
    const { loading, error, success, startProcedure } = useStartProcedure(ratingID, raceID);

    return (
        <div>
            <div data-testid="loading">{loading && "Загрузка..."}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="success">{success}</div>
            <button
                onClick={() => startProcedure({ specCircumstance: 1, falseStartList: [] })}
                data-testid="start-button"
            >
                Начать процедуру
            </button>
        </div>
    );
};

describe("useStartProcedure", () => {
    const ratingID = "123";
    const raceID = "456";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное выполнение процедуры старта", async () => {
        const mockResponse = { data: "Процедура успешно завершена" };
        (api.post as jest.Mock).mockResolvedValue(mockResponse);

        render(<StartProcedureWrapper ratingID={ratingID} raceID={raceID} />);

        act(() => {
            screen.getByTestId("start-button").click();
        });

        // Ожидаем завершения запроса
        await act(async () => {});

        expect(screen.getByTestId("success").textContent).toBe("Процедура успешно завершена");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/start`,
            { specCircumstance: 1, falseStartList: [] }
        );
    });

    test("ошибка при выполнении процедуры старта", async () => {
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка при старте процедуры"));

        render(<StartProcedureWrapper ratingID={ratingID} raceID={raceID} />);

        act(() => {
            screen.getByTestId("start-button").click();
        });

        await act(async () => {});

        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("success").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/start`,
            { specCircumstance: 1, falseStartList: [] }
        );
    });

    test("показ индикатора загрузки при выполнении процедуры", async () => {
        api.post.mockImplementation(() => new Promise(() => {}));

        render(<StartProcedureWrapper ratingID={ratingID} raceID={raceID} />);

        act(() => {
            screen.getByTestId("start-button").click();
        });

        expect(screen.getByTestId("loading").textContent).toBe("Загрузка...");
    });

    test("сброс состояния перед новым запросом", async () => {
        const mockResponse = { data: "Процедура успешно завершена" };
        (api.post as jest.Mock).mockResolvedValue(mockResponse);

        render(<StartProcedureWrapper ratingID={ratingID} raceID={raceID} />);

        act(() => {
            screen.getByTestId("start-button").click();
        });

        await act(async () => {});

        expect(screen.getByTestId("success").textContent).toBe("Процедура успешно завершена");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
    });
});
