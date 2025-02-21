import { render, act, screen } from "@testing-library/react";
import { useDeleteParticipant } from "../../ui/controllers/participantControllers/deleteParticipantController"; // Путь может отличаться
import api from "../api";
import React from "react";

// Компонент-обертка для тестирования хука
const DeleteParticipantWrapper = ({ participantID }: { participantID: string }) => {
    const { deleteParticipant, loading, error, success } = useDeleteParticipant();

    const handleDelete = () => deleteParticipant(participantID);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="success">{success !== null ? success.toString() : "not attempted"}</div>
            <button onClick={handleDelete} data-testid="delete-button">
                Delete Participant
            </button>
        </div>
    );
};

describe("useDeleteParticipant controller", () => {
    const participantID = "12345";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное удаление участника", async () => {
        // Мокаем успешный ответ API
        (api.delete as jest.Mock).mockResolvedValue({});

        render(<DeleteParticipantWrapper participantID={participantID} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("not attempted");

        // Нажимаем на кнопку удаления
        act(() => {
            screen.getByTestId("delete-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения действия
        await act(async () => {});

        // Проверяем успешность удаления
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("true");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.delete).toHaveBeenCalledWith(`/participants/${participantID}`);
    });

    test("ошибка при удалении участника", async () => {
        // Мокаем ошибку API
        (api.delete as jest.Mock).mockRejectedValue(new Error("Ошибка при удалении участника"));

        render(<DeleteParticipantWrapper participantID={participantID} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("not attempted");

        // Нажимаем на кнопку удаления
        act(() => {
            screen.getByTestId("delete-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("not attempted");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка при удалении участника");
    });

    test("проверка нескольких попыток удаления", async () => {
        const successMessage = "true";
        const failureMessage = "not attempted";

        // Мокаем успешный ответ для первого запроса
        (api.delete as jest.Mock).mockResolvedValueOnce({});
        // Мокаем ошибку для второго запроса
        (api.delete as jest.Mock).mockRejectedValueOnce(new Error("Ошибка при удалении участника"));

        render(<DeleteParticipantWrapper participantID={participantID} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("not attempted");

        // Нажимаем на кнопку удаления
        act(() => {
            screen.getByTestId("delete-button").click();
        });

        // Проверяем состояние после первой попытки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");
        await act(async () => {});

        // После первой попытки успешного удаления
        expect(screen.getByTestId("success").textContent).toBe(successMessage);

        // Нажимаем на кнопку удаления второй раз
        act(() => {
            screen.getByTestId("delete-button").click();
        });

        // Проверяем состояние после второй попытки (ошибки)
        expect(screen.getByTestId("loading").textContent).toBe("loading...");
        await act(async () => {});

        expect(screen.getByTestId("success").textContent).toBe(failureMessage);
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка при удалении участника");
    });
});
