import { render, act, screen } from "@testing-library/react";
import { useGetRatingsController, filterAndGroupRatings } from "../../ui/controllers/ratingControllers/getRatingsController"; // Путь может отличаться
import api from "../api";
import React from "react";

// Компонент-обертка для тестирования хука
const GetRatingsWrapper = () => {
    const { ratings, error, setRatings } = useGetRatingsController();

    return (
        <div>
            <div data-testid="error">{error}</div>
            <div data-testid="ratings">{JSON.stringify(ratings)}</div>
            <button onClick={() => setRatings([])} data-testid="clear-button">
                Clear Ratings
            </button>
        </div>
    );
};

describe("useGetRatingsController", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение рейтингов", async () => {
        // Мокаем успешный ответ API
        const mockRatings = [
            { Name: "Good", Class: "1", BlowoutCnt: 5 },
            { Name: "Bad", Class: "2", BlowoutCnt: 3 },
        ];
        (api.get as jest.Mock).mockResolvedValue({
            data: mockRatings,
        });

        render(<GetRatingsWrapper />);

        // Ожидаем загрузки рейтингов
        await act(async () => {});

        // Проверяем, что данные рейтингов отобразились
        expect(screen.getByTestId("ratings").textContent).toBe(
            JSON.stringify(mockRatings)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith("http://go-server:8081/api/ratings");
    });

    test("ошибка при получении рейтингов", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка при получении данных"));

        render(<GetRatingsWrapper />);

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем, что ошибка отображается
        expect(screen.getByTestId("error").textContent).toBe("Ошибка при получении данных");
        expect(screen.getByTestId("ratings").textContent).toBe("[]");
        expect(api.get).toHaveBeenCalledWith("http://go-server:8081/api/ratings");
    });

    test("проверка очистки рейтингов", async () => {
        // Мокаем успешный ответ API
        const mockRatings = [
            { Name: "Good", Class: "1", BlowoutCnt: 5 },
            { Name: "Bad", Class: "2", BlowoutCnt: 3 },
        ];
        (api.get as jest.Mock).mockResolvedValue({
            data: mockRatings,
        });

        render(<GetRatingsWrapper />);

        // Ожидаем загрузки рейтингов
        await act(async () => {});

        // Проверяем, что данные рейтингов отобразились
        expect(screen.getByTestId("ratings").textContent).toBe(
            JSON.stringify(mockRatings)
        );

        // Нажимаем на кнопку очистки
        act(() => {
            screen.getByTestId("clear-button").click();
        });

        // Проверяем, что рейтинги очищены
        expect(screen.getByTestId("ratings").textContent).toBe("[]");
    });
});

describe("filterAndGroupRatings", () => {
    const ratings = [
        { id: "", Name: "Good", Class: "1", BlowoutCnt: 5 },
        { id: "", Name: "Bad", Class: "2", BlowoutCnt: 3 },
        { id: "", Name: "Excellent", Class: "1", BlowoutCnt: 7 },
    ];

    test("фильтрация по имени", () => {
        const filtered = filterAndGroupRatings(ratings, { name: "Go", class: "", blowoutCnt: "" });
        expect(filtered).toEqual([
            { Name: "Good", Class: "1", BlowoutCnt: 5, id:"" },
        ]);
    });

    test("фильтрация по классу", () => {
        const filtered = filterAndGroupRatings(ratings, { name: "", class: "1", blowoutCnt: "" });
        expect(filtered).toEqual([
            { Name: "Excellent", Class: "1", BlowoutCnt: 7, id:"" },
            { Name: "Good", Class: "1", BlowoutCnt: 5, id:"" },
        ]);
    });

    test("фильтрация по количеству выбросов", () => {
        const filtered = filterAndGroupRatings(ratings, { name: "", class: "", blowoutCnt: "5" });
        expect(filtered).toEqual([
            { Name: "Good", Class: "1", BlowoutCnt: 5, id:"" },
        ]);
    });

    test("группировка и сортировка по имени", () => {
        const filtered = filterAndGroupRatings(ratings, { name: "", class: "", blowoutCnt: "" });
        expect(filtered).toEqual([
            { Name: "Bad", Class: "2", BlowoutCnt: 3, id:"" },
            { Name: "Excellent", Class: "1", BlowoutCnt: 7, id:"" },
            { Name: "Good", Class: "1", BlowoutCnt: 5, id:"" },
        ]);
    });
});
