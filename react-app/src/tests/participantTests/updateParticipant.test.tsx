import { render, act, screen } from "@testing-library/react";
import { useUpdateParticipant } from "../../ui/controllers/participantControllers/updateParticipantController"; // Замените путь на правильный
import api from "../api";
import { ParticipantFormData } from "../../ui/models/participantModel";
import React from "react";

// Компонент-обертка для тестирования хука
const UpdateParticipantWrapper = ({ participantID }: { participantID: string }) => {
    const { updateParticipant, loading, error, updatedParticipant } = useUpdateParticipant(participantID);

    return (
        <div>
            <button
                onClick={() =>
                    updateParticipant({
                        fio: "Иван Иванов",
                        category: 1,
                        birthday: "1990-01-01",
                        coach: "Ирина Петрова",
                        gender: 1, // Можно оставить опциональным
                    })
                }
                data-testid="update-button"
            >
                Update Participant
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="updated-participant">
                {updatedParticipant ? JSON.stringify(updatedParticipant) : "No participant updated"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useUpdateParticipant controller", () => {
    const participantID = "participant123";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление участника", async () => {
        const mockUpdatedParticipant: ParticipantFormData = {
            id: participantID,
            FIO: "Иван Иванов",
            Category: "1",
            Birthday: "1990-01-01",
            Coach: "Ирина Петрова",
            Gender: "1",
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedParticipant });

        render(<UpdateParticipantWrapper participantID={participantID} />);

        // Нажимаем кнопку для обновления данных участника
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после успешного обновления
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("updated-participant").textContent).toBe(
            JSON.stringify(mockUpdatedParticipant)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.put).toHaveBeenCalledWith(`/participants/${participantID}`, {
            fio: "Иван Иванов",
            category: 1,
            birthday: "1990-01-01",
            coach: "Ирина Петрова",
            gender: 1,
        });
    });

    test("ошибка при обновлении участника", async () => {
        // Мокаем ошибку API
        (api.put as jest.Mock).mockRejectedValue(new Error("Ошибка обновления участника"));

        render(<UpdateParticipantWrapper participantID={participantID} />);

        // Нажимаем кнопку для обновления данных участника
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("updated-participant").textContent).toBe("No participant updated");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка обновления участника");
    });

    test("проверка обновления при повторном вызове", async () => {
        const mockUpdatedParticipant: ParticipantFormData = {
            id: participantID,
            FIO: "Иван Иванов",
            Category: "1",
            Birthday: "1990-01-01",
            Coach: "Ирина Петрова",
            Gender: "1",
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedParticipant });

        render(<UpdateParticipantWrapper participantID={participantID} />);

        // Нажимаем кнопку для обновления данных участника дважды
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
            updateButton.click(); // Дважды для проверки
        });

        expect(api.put).toHaveBeenCalledTimes(2);
    });
});
