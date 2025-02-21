import { render, act, screen } from "@testing-library/react";
import { useCreateParticipant } from "../../ui/controllers/participantControllers/createParticipantController"; // Путь может отличаться
import api from "../api";
import React from "react";

// Компонент-обертка для тестирования хука
const CreateParticipantWrapper = ({ participantData }: { participantData: any }) => {
    const { createParticipant, loading, error, createdParticipant } = useCreateParticipant();

    const handleCreate = () => createParticipant(participantData);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="created">
                {createdParticipant ? JSON.stringify(createdParticipant) : "not created"}
            </div>
            <button onClick={handleCreate} data-testid="create-button">
                Create Participant
            </button>
        </div>
    );
};

describe("useCreateParticipant controller", () => {
    const participantData = {
        firstName: "John",
        lastName: "Doe",
        email: "john.doe@example.com",
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное создание участника", async () => {
        // Мокаем успешный ответ API
        (api.post as jest.Mock).mockResolvedValue({
            data: {
                ...participantData,
                id: "12345",
            },
        });

        render(<CreateParticipantWrapper participantData={participantData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("created").textContent).toBe("not created");

        // Нажимаем на кнопку создания участника
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения процесса
        await act(async () => {});

        // Проверяем успешное создание участника
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("created").textContent).toBe(
            '{"firstName":"John","lastName":"Doe","email":"john.doe@example.com","id":"12345"}'
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.post).toHaveBeenCalledWith("/participants", participantData);
    });

    test("ошибка при создании участника", async () => {
        // Мокаем ошибку API
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка при создании участника"));

        render(<CreateParticipantWrapper participantData={participantData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("created").textContent).toBe("not created");

        // Нажимаем на кнопку создания участника
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("created").textContent).toBe("not created");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка при создании участника");
    });

    test("проверка состояния загрузки", async () => {
        // Мокаем успешный ответ API
        (api.post as jest.Mock).mockResolvedValue({
            data: {
                ...participantData,
                id: "12345",
            },
        });

        render(<CreateParticipantWrapper participantData={participantData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");

        // Нажимаем на кнопку создания участника
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки во время выполнения запроса
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения процесса
        await act(async () => {});

        // Проверяем состояние после завершения запроса
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
    });
});
