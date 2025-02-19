import { render, act, screen } from "@testing-library/react";
import { useGetCrewMember } from "../../ui/controllers/crewMemberControllers/getCrewMemberController";
import api from "../api";
import {ParticipantFormData} from "../../ui/models/participantModel";

// Компонент для теста
const GetCrewMemberWrapper = ({
                                  ratingID,
                                  crewID,
                                  participantID,
                              }: {
    ratingID: string;
    crewID: string;
    participantID: string;
}) => {
    const { data, loading, error, getCrewMember } = useGetCrewMember();

    return (
        <div>
            <button
                onClick={() => getCrewMember(ratingID, crewID, participantID)}
                data-testid="get-button"
            >
                Get Crew Member
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="data">
                {data ? JSON.stringify(data) : "No data"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useGetCrewMember controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение данных члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant123";

        const mockResponse: ParticipantFormData = {
            id: mockParticipantID,
            FIO: "Иван Иванов",
            Category: "A",
            Gender: "Male",
            Birthday: "2000-01-01",
            Coach: "Петров П.П.",
        };

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(
            <GetCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
            />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const getButton = screen.getByTestId("get-button");
        await act(async () => {
            getButton.click();
        });

        // Проверяем состояние после успешного получения данных
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews/${mockCrewID}/members/${mockParticipantID}`
        );
    });

    test("ошибка при получении данных члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant123";

        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения данных"));

        render(
            <GetCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
            />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const getButton = screen.getByTestId("get-button");
        await act(async () => {
            getButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe("No data");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения данных");
    });
});
