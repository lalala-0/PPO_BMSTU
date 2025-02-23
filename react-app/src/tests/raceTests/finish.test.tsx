import { render, fireEvent, screen, waitFor } from "@testing-library/react";
import api from "../api";
import {FinishInput} from "../../ui/models/raceModel";
import {useFinishProcedure} from "../../ui/controllers/raceControllers/finishProcedureController"; // Путь к API
import React from "react";

jest.mock("../api"); // Мокаем API

describe("useFinishProcedure", () => {
    it("должен показывать успешный результат при успешном запросе", async () => {
        // Мокаем успешный ответ от API
        const mockResponse = { data: "Процедура финиша успешно выполнена" };
        api.post.mockResolvedValue(mockResponse);

        render(<MockComponent ratingID="123" raceID="456" />);

        // Вызываем функцию финиша с нужным input
        fireEvent.click(screen.getByText("Завершить гонку"));

        await waitFor(() => {
            expect(screen.getByText("Процедура финиша успешно выполнена")).toBeInTheDocument();
        });
    });

    it("должен показывать состояние загрузки при отправке запроса", async () => {
        // Мокаем запрос, чтобы он не завершался сразу
        api.post.mockImplementation(() => new Promise(() => {})); // Запрос "зависает"

        render(<MockComponent ratingID="123" raceID="456" />);

        // Вызываем функцию финиша с нужным input
        fireEvent.click(screen.getByText("Завершить гонку"));

        // Проверяем, что отображается индикатор загрузки
        expect(screen.getByText("Загрузка...")).toBeInTheDocument();
    });

    it("должен сбрасывать состояния error и success перед отправкой запроса", async () => {
        const mockResponse = { data: "Процедура финиша успешно выполнена" };
        api.post.mockResolvedValue(mockResponse);

        render(<MockComponent ratingID="123" raceID="456" />);

        fireEvent.click(screen.getByText("Завершить гонку"));

        await waitFor(() => {
            expect(screen.queryByText("Что-то пошло не так")).not.toBeInTheDocument();
            expect(screen.queryByText("Процедура финиша успешно выполнена")).toBeInTheDocument();
        });
    });
});

// Компонент, который использует хук
const MockComponent = ({ ratingID, raceID }: { ratingID: string, raceID: string }) => {
    const { loading, error, success, finishProcedure } = useFinishProcedure(ratingID, raceID);

    const input: FinishInput = { finisherList: [1, 2, 3] };

    return (
        <div>
            <button onClick={() => finishProcedure(input)}>Завершить гонку</button>
            {loading && <div>Загрузка...</div>}
            {error && <div>{error}</div>}
            {success && <div>{success}</div>}
        </div>
    );
};
