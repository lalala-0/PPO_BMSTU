import { render, act, screen } from "@testing-library/react";
import {useAttachCrewMember} from "../../ui/controllers/crewMemberControllers/attachCrewMemberController";
import api from "../api";

// Компонент для теста
const AttachCrewMemberWrapper = ({ ratingID, crewID }: { ratingID: string, crewID: string }) => {
    const { attachCrewMember, success, loading, error } = useAttachCrewMember(ratingID, crewID);

    return (
        <div>
            <button onClick={() => attachCrewMember({ participantID: "string",  helmsman: 1 })} data-testid="attach-button">
                Attach Crew Member
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useAttachCrewMember controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное добавление члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockCrewParticipant = { participantID: "string",  helmsman: 1 };

        // Мокаем успешный ответ
        (api.post as jest.Mock).mockResolvedValue({ data: { participantID: "string",  helmsman: 1 } });

        const { getByTestId } = render(<AttachCrewMemberWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать добавление
        const attachButton = getByTestId("attach-button");
        await act(async () => {
            attachButton.click();
        });

        // Проверяем, что флаг success установлен в true
        expect(getByTestId("success").textContent).toBe("Success");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews/${mockCrewID}/members`,
            mockCrewParticipant
        );
    });

    test("ошибка при добавлении члена экипажа", async () => {
        const mockRatingID = "rating123";
        const mockCrewID = "crew123";
        const mockCrewParticipant = { memberID: "member123" };

        // Мокаем ошибку в API
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка добавления"));

        const { getByTestId } = render(<AttachCrewMemberWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать добавление
        const attachButton = getByTestId("attach-button");
        await act(async () => {
            attachButton.click();
        });

        // Проверяем, что флаг success установлен в false и ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка добавления");
    });
});
