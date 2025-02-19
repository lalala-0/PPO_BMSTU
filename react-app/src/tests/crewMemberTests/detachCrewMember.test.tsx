import { render, act, screen } from "@testing-library/react";
import { useDetachCrewMember } from "../../ui/controllers/crewMemberControllers/detachCrewMemberController";
import api from "../api";

// Компонент для теста
const DetachCrewMemberWrapper = ({
                                     ratingID,
                                     crewID,
                                     participantID,
                                 }: {
    ratingID: string;
    crewID: string;
    participantID: string;
}) => {
    const { detachCrewMember, success, loading, error } = useDetachCrewMember();

    return (
        <div>
            <button
                onClick={() => detachCrewMember(ratingID, crewID, participantID)}
                data-testid="detach-button"
            >
                Detach Crew Member
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useDetachCrewMember controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное удаление члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant123";

        // Мокаем успешный ответ API
        (api.delete as jest.Mock).mockResolvedValue({});

        render(
            <DetachCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
            />
        );

        // Нажимаем кнопку, чтобы инициировать удаление
        const detachButton = screen.getByTestId("detach-button");
        await act(async () => {
            detachButton.click();
        });

        // Проверяем состояние после удаления
        expect(screen.getByTestId("success").textContent).toBe("Success");
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.delete).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews/${mockCrewID}/members/${mockParticipantID}`
        );
    });

    test("ошибка при удалении члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockParticipantID = "participant123";

        // Мокаем ошибку API
        (api.delete as jest.Mock).mockRejectedValue(new Error("Ошибка удаления"));

        render(
            <DetachCrewMemberWrapper
                ratingID={mockRatingID}
                crewID={mockCrewID}
                participantID={mockParticipantID}
            />
        );

        // Нажимаем кнопку, чтобы инициировать удаление
        const detachButton = screen.getByTestId("detach-button");
        await act(async () => {
            detachButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("success").textContent).toBe("Failure");
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("error").textContent).toBe(
            "Ошибка запроса: Ошибка удаления"
        );
    });

    test("ошибка из-за недостаточных данных", async () => {
        render(
            <DetachCrewMemberWrapper ratingID="" crewID="" participantID="" />
        );

        // Нажимаем кнопку, чтобы инициировать удаление
        const detachButton = screen.getByTestId("detach-button");
        await act(async () => {
            detachButton.click();
        });

        // Проверяем, что отображается ошибка
        expect(screen.getByTestId("success").textContent).toBe("Failure");
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("error").textContent).toBe(
            "Недостаточно данных для выполнения запроса"
        );
        expect(api.delete).not.toHaveBeenCalled(); // Запрос не должен быть отправлен
    });
});
