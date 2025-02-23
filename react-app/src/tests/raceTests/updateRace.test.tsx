import { render, act, screen } from "@testing-library/react";
import api from "../api";
import React from "react";
import { handleError } from "../errorHandler";
import {useUpdateRace} from "../../ui/controllers/raceControllers/updateRaceController";

jest.mock("../api");
jest.mock("../errorHandler");

const MockComponent = () => {
    const { success, error, loading, handleUpdate } = useUpdateRace();
    const ratingID = "123"; // Просто пример
    const raceID = "1"; // Просто пример
    const updatedData = { date: "2025-02-21", number: 5, class: 1 }; // Пример обновленных данных

    return (
        <div>
            <div data-testid="loading">{loading && "Загрузка..."}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="success">{success}</div>
            <button
                onClick={() => handleUpdate(ratingID, raceID, updatedData)}
                data-testid="update-button"
            >
                Обновить гонку
            </button>
        </div>
    );
};

describe("useUpdateRace", () => {
    const mockRatingID = "123";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление гонки", async () => {
        const mockResponse = { data: "Гонка успешно обновлена" };
        (api.put as jest.Mock).mockResolvedValue(mockResponse);

        render(<MockComponent />);

        act(() => {
            screen.getByTestId("update-button").click();
        });

        // Ожидаем завершения запроса
        await act(async () => {});

        expect(screen.getByTestId("success").textContent).toBe("Гонка успешно обновлена");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
        expect(api.put).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/races/1`,
            { date: "2025-02-21", number: 5, class: 1 }
        );
    });

    test("ошибка при обновлении гонки", async () => {
        const errorMessage = "Ошибка при обновлении гонки";
        (api.put as jest.Mock).mockRejectedValue(new Error(errorMessage));
        (handleError as jest.Mock).mockImplementation((err, setError) => {
            setError(err.message);
        });

        render(<MockComponent />);

        act(() => {
            screen.getByTestId("update-button").click();
        });

        await act(async () => {});

        expect(screen.getByTestId("error").textContent).toBe(errorMessage);
        expect(screen.getByTestId("success").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
    });

    test("не обновляется гонка, если отсутствует идентификатор рейтинга", async () => {
        render(<MockComponent />);

        act(() => {
            screen.getByTestId("update-button").click();
        });

        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("success").textContent).toBe("");
    });

    test("показ индикатора загрузки при обновлении гонки", async () => {
        api.put.mockImplementation(() => new Promise(() => {}));

        render(<MockComponent />);

        act(() => {
            screen.getByTestId("update-button").click();
        });

        expect(screen.getByTestId("loading").textContent).toBe("Загрузка...");
    });

    test("состояние сбрасывается перед новым запросом", async () => {
        const mockResponse = { data: "Гонка успешно обновлена" };
        (api.put as jest.Mock).mockResolvedValue(mockResponse);

        render(<MockComponent />);

        act(() => {
            screen.getByTestId("update-button").click();
        });

        await act(async () => {});

        expect(screen.getByTestId("success").textContent).toBe("Гонка успешно обновлена");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("loading").textContent).toBe("");
    });
});
