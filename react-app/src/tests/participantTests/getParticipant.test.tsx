import { render, act, screen } from "@testing-library/react";
import { useGetParticipant } from "../../ui/controllers/participantControllers/getParticipantController"; // Путь может отличаться
import api from "../api";
import { ParticipantFormData } from "../../ui/models/participantModel";
import React from "react";

// Компонент-обертка для тестирования хука
const GetParticipantWrapper = ({ participantID }: { participantID: string }) => {
    const { getParticipant, loading, error, participant } = useGetParticipant(participantID);

    React.useEffect(() => {
        getParticipant();
    }, [participantID]);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="participant">
                {participant ? JSON.stringify(participant) : "No participant found"}
            </div>
        </div>
    );
};

describe("useGetParticipant controller", () => {
    const participantID = "12345";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение данных участника", async () => {
        const mockParticipant: ParticipantFormData = {
            id: "12345",
            FIO: "Иван Иванов",
            Category: "1",
            Birthday: "1990-01-01",
            Coach: "Ирина Петрова",
            Gender: "1",
        };

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockParticipant });

        render(<GetParticipantWrapper participantID={participantID} />);

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем получения данных участника
        await act(async () => {});

        // Проверяем, что данные участника отображаются корректно
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participant").textContent).toBe(
            JSON.stringify(mockParticipant)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/participants/${participantID}`);
    });

    test("ошибка при получении данных участника", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения участника"));

        render(<GetParticipantWrapper participantID={participantID} />);

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participant").textContent).toBe("No participant found");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения участника");
    });

    test("проверка с неправильным participantID", async () => {
        const invalidID = "wrongID";

        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения участника"));

        render(<GetParticipantWrapper participantID={invalidID} />);

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participant").textContent).toBe("No participant found");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения участника");
    });

    test("проверка обновления participantID", async () => {
        const initialID = "12345";
        const updatedID = "67890";

        const mockParticipant: ParticipantFormData = {
            id: "12345",
            FIO: "Иван Иванов",
            Category: "1",
            Birthday: "1990-01-01",
            Coach: "Ирина Петрова",
            Gender: "1",
        };

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockParticipant });

        render(<GetParticipantWrapper participantID={initialID} />);

        // Проверяем, что запрос был отправлен с первоначальным ID
        expect(api.get).toHaveBeenCalledWith(`/participants/${initialID}`);

        // Обновляем participantID
        render(<GetParticipantWrapper participantID={updatedID} />);

        // Проверяем, что запрос был повторно отправлен с обновленным ID
        expect(api.get).toHaveBeenCalledWith(`/participants/${updatedID}`);
    });
});
