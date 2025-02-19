import { render, act, screen } from "@testing-library/react";
import { useUpdateCrewMember } from "../../ui/controllers/crewMemberControllers/updateCrewMemberController";
import api from "../api";
import {ParticipantFormData, ParticipantInput} from "../../ui/models/participantModel";

// Компонент-обертка для тестирования хука
const UpdateCrewMemberWrapper = ({
                                     ratingID,
                                     crewID,
                                     participantID,
                                     participantData,
                                 }: {
    ratingID: string;
    crewID: string;
    participantID: string;
    participantData: ParticipantInput;
}) => {
    const { updateCrewMember, loading, error, data } = useUpdateCrewMember();

    return (
        <div>
            <button
                onClick={() => updateCrewMember(ratingID, crewID, participantID, participantData)}
                data-testid="update-button"
            >
                Update Crew Member
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="data">{data ? JSON.stringify(data) : "No data"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useUpdateCrewMember controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление данных участника", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant1";

        const mockInput: ParticipantInput = {
            id: mockParticipantID,
            fio: "Иван Иванов",
            category: 1,
            gender: 0,
            birthday: "2000-01-01",
            coach: "Петров П.П.",
        };

        const mockResponse: ParticipantFormData = {
            id: mockParticipantID,
            FIO: mockInput.fio,
            Category: String(mockInput.category),
            Gender: String(mockInput.gender),
            Birthday: mockInput.birthday,
            Coach: mockInput.coach,
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(
            <UpdateCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
                participantData={mockInput}
            />
        );

        // Нажимаем кнопку, чтобы инициировать обновление данных
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после успешного обновления данных
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.put).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews/${mockCrewID}/members/${mockParticipantID}`,
            mockInput
        );
    });

    test("ошибка при обновлении данных участника", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant1";

        const mockInput: ParticipantInput = {
            fio: "Иван Иванов",
            category: 1,
            gender: 0,
            birthday: "2000-01-01",
            coach: "Петров П.П.",
        };

        // Мокаем ошибку API
        (api.put as jest.Mock).mockRejectedValue(new Error("Ошибка обновления данных"));

        render(
            <UpdateCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
                participantData={mockInput}
            />
        );

        // Нажимаем кнопку, чтобы инициировать обновление данных
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe("No data");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка обновления данных");
    });
});
