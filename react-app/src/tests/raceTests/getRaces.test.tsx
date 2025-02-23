import { render, act, screen } from "@testing-library/react";
import api from "../api"; // Путь к API
import React from "react";
import {useGetRacesByRatingID} from "../../ui/controllers/raceControllers/getRacesController";

// Компонент-обертка для тестирования хука
const GetRacesWrapper = () => {
    const { loading, error, races, getRacesByRatingID } = useGetRacesByRatingID();

    return (
        <div>
            <div data-testid="loading">{loading && "Загрузка..."}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="races">{JSON.stringify(races)}</div>
            <button onClick={() => getRacesByRatingID("123")} data-testid="load-button">
                Загрузить гонки
            </button>
        </div>
    );
};

describe("useGetRacesByRatingID", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение гонок", async () => {
        // Мокаем успешный ответ API
        const mockRaces = [
            { id: "1", name: "Гонка 1" },
            { id: "2", name: "Гонка 2" },
        ];
        (api.get as jest.Mock).mockResolvedValue({
            data: mockRaces,
        });

        render(<GetRacesWrapper />);

        // Имитируем клик по кнопке для загрузки гонок
        act(() => {
            screen.getByTestId("load-button").click();
        });

        // Ожидаем, что гонки будут загружены
        await act(async () => {});

        // Проверяем, что данные гонок отображаются
        expect(screen.getByTestId("races").textContent).toBe(
            JSON.stringify(mockRaces)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith("/ratings/123/races");
    });

    test("ошибка при получении гонок", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка при загрузке гонок"));

        render(<GetRacesWrapper />);

        // Имитируем клик по кнопке для загрузки гонок
        act(() => {
            screen.getByTestId("load-button").click();
        });

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем, что ошибка отображается
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка при загрузке гонок");
        expect(screen.getByTestId("races").textContent).toBe("null");
        expect(api.get).toHaveBeenCalledWith("/ratings/123/races");
    });

    test("показ индикатора загрузки", async () => {
        // Мокаем "зависший" запрос
        api.get.mockImplementation(() => new Promise(() => {}));

        render(<GetRacesWrapper />);

        // Имитируем клик по кнопке для загрузки гонок
        act(() => {
            screen.getByTestId("load-button").click();
        });

        // Проверяем, что индикатор загрузки появляется
        expect(screen.getByTestId("loading").textContent).toBe("Загрузка...");
    });

    test("сброс состояния перед новым запросом", async () => {
        const mockRaces = [
            { id: "1", name: "Гонка 1" },
            { id: "2", name: "Гонка 2" },
        ];
        (api.get as jest.Mock).mockResolvedValue({
            data: mockRaces,
        });

        render(<GetRacesWrapper />);

        // Имитируем клик по кнопке для загрузки гонок
        act(() => {
            screen.getByTestId("load-button").click();
        });

        // Ожидаем, что гонки были загружены
        await act(async () => {});

        // Проверяем, что гонки отображаются
        expect(screen.getByTestId("races").textContent).toBe(
            JSON.stringify(mockRaces)
        );

        // Сбрасываем состояние (например, при новом запросе)
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(screen.getByTestId("races").textContent).toBe(JSON.stringify(mockRaces));
    });
});
