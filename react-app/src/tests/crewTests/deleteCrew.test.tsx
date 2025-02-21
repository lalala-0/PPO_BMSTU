import { render, act } from "@testing-library/react";
import { useDeleteCrew } from "../../ui/controllers/crewControllers/deleteCrewController";
import api from "../api";
import React from "react";

jest.mock("../api"); // Мокаем api.get

const DeleteCrewWrapper = ({ ratingID, crewID }: { ratingID: string, crewID: string }) => {
    const { loading, error, success, deleteCrewByID } = useDeleteCrew();

    return (
        <div>
            <button onClick={() => deleteCrewByID(ratingID, crewID)} data-testid="delete-crew-button">
                Delete Crew
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useDeleteCrew hook", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное удаление команды", async () => {
        const mockRatingID = "123";
        const mockCrewID = "1";
        (api.delete as jest.Mock).mockResolvedValue({}); // Мокаем успешное удаление

        const { getByTestId } = render(<DeleteCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteCrewButton = getByTestId("delete-crew-button");
        await act(async () => {
            deleteCrewButton.click();
        });

        // Проверяем, что успешное сообщение отображается
        expect(getByTestId("success").textContent).toBe("Команда успешно удалена");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
        expect(api.delete).toHaveBeenCalledWith(`/ratings/${mockRatingID}/crews/${mockCrewID}`);
    });

    test("ошибка при удалении команды", async () => {
        const mockRatingID = "123";
        const mockCrewID = "1";
        const errorMessage = "Ошибка удаления команды";
        (api.delete as jest.Mock).mockRejectedValue(new Error(errorMessage)); // Мокаем ошибку при удалении

        const { getByTestId } = render(<DeleteCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteCrewButton = getByTestId("delete-crew-button");
        await act(async () => {
            deleteCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe(`Ошибка запроса: ${errorMessage}`);
    });

    test("ошибка при сетевом запросе", async () => {
        const mockRatingID = "123";
        const mockCrewID = "1";
        (api.delete as jest.Mock).mockRejectedValue({ request: {} }); // Мокаем ошибку сети

        const { getByTestId } = render(<DeleteCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteCrewButton = getByTestId("delete-crew-button");
        await act(async () => {
            deleteCrewButton.click();
        });

        // Проверяем, что ошибка сети отображается
        expect(getByTestId("success").textContent).toBe("");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка сети. Проверьте подключение к интернету.");
    });

    test("непредвиденная ошибка при удалении", async () => {
        const mockRatingID = "123";
        const mockCrewID = "1";
        (api.delete as jest.Mock).mockRejectedValue(null); // Мокаем непредвиденную ошибку

        const { getByTestId } = render(<DeleteCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteCrewButton = getByTestId("delete-crew-button");
        await act(async () => {
            deleteCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Произошла неизвестная ошибка.");
    });
});
