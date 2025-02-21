import { render, act, screen } from "@testing-library/react";
import { useGetCrewMembers } from "../../ui/controllers/crewMemberControllers/getCrewMembersController";
import api from "../api";
import {ParticipantFormData} from "../../ui/models/participantModel";
import React from "react";

// Компонент для теста
const GetCrewMembersWrapper = ({
                                   ratingID,
                                   crewID,
                               }: {
    ratingID: string;
    crewID: string;
}) => {
    const { data, loading, error, getCrewMembers } = useGetCrewMembers(
        ratingID,
        crewID
    );

    return (
        <div>
            <button onClick={getCrewMembers} data-testid="get-button">
                Get Crew Members
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="data">
                {data ? JSON.stringify(data) : "No data"}
            </div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useGetCrewMembers controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение списка членов экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";

        const mockResponse: ParticipantFormData[] = [
            {
                id: "participant1",
                FIO: "Иван Иванов",
                Category: "A",
                Gender: "Male",
                Birthday: "2000-01-01",
                Coach: "Петров П.П.",
            },
            {
                id: "participant2",
                FIO: "Мария Смирнова",
                Category: "B",
                Gender: "Female",
                Birthday: "1999-05-15",
                Coach: "Сидоров С.С.",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockResponse });

        render(<GetCrewMembersWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getButton = screen.getByTestId("get-button");
        await act(async () => {
            getButton.click();
        });

        // Проверяем состояние после успешного получения данных
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("data").textContent).toBe(JSON.stringify(mockResponse));
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/ratings/${mockRatingID}/crews/${mockCrewID}/members`);
    });

    test("ошибка при получении списка членов экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";

        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения данных"));

        render(<GetCrewMembersWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

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
