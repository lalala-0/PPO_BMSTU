import { render } from "@testing-library/react";
import { act } from "react-dom/test-utils";
import { useUpdateCrew } from "../../ui/controllers/crewControllers/updateCrewController"; // Путь к твоему хук-файлу
import api from "../api"; // Путь к твоему мокированному API
import React from "react";

// Мокаем api.put
jest.mock("../api");

const UpdateCrewWrapper = ({ ratingID, crewID, updatedData }: { ratingID: string; crewID: string; updatedData: any }) => {
    const { loading, error, success, handleUpdate } = useUpdateCrew();

    // Используем useUpdateCrew внутри компонента
    return (
        <div>
            <button onClick={() => handleUpdate(ratingID, crewID, updatedData)} data-testid="update-crew-button">
                Update Crew
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useUpdateCrew hook", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление данных команды", async () => {
        const mockRatingID = "123";
        const mockCrewID = "10";
        const updatedData = { sailNum: 2 };

        const mockResponse = { data: { success: true } };
        (api.put as jest.Mock).mockResolvedValue(mockResponse);

        const { getByTestId } = render(
            <UpdateCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} updatedData={updatedData} />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const updateCrewButton = getByTestId("update-crew-button");
        await act(async () => {
            updateCrewButton.click();
        });

        // Проверяем, что запрос был вызван с правильными параметрами
        expect(api.put).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews/${mockCrewID}`,
            updatedData
        );

        // Проверяем состояние после успешного запроса
        expect(getByTestId("success").textContent).toBe("Success");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
    });

    test("ошибка при обновлении данных команды", async () => {
        const mockRatingID = "123";
        const mockCrewID = "10";
        const updatedData = { sailNum: 2 };

        (api.put as jest.Mock).mockRejectedValue({
            message: "Ошибка обновления данных",
        });

        const { getByTestId } = render(
            <UpdateCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} updatedData={updatedData} />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const updateCrewButton = getByTestId("update-crew-button");
        await act(async () => {
            updateCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка обновления данных");
    });

    test("ошибка при сетевом запросе", async () => {
        const mockRatingID = "123";
        const mockCrewID = "10";
        const updatedData = { sailNum: 2 };

        (api.put as jest.Mock).mockRejectedValue({ request: {} }); // Мокаем ошибку сети

        const { getByTestId } = render(
            <UpdateCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} updatedData={updatedData} />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const updateCrewButton = getByTestId("update-crew-button");
        await act(async () => {
            updateCrewButton.click();
        });

        // Проверяем, что ошибка сети отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка сети. Проверьте подключение к интернету.");
    });

    test("непредвиденная ошибка при запросе", async () => {
        const mockRatingID = "123";
        const mockCrewID = "10";
        const updatedData = { sailNum: 2 };

        (api.put as jest.Mock).mockRejectedValue(null); // Мокаем непредвиденную ошибку

        const { getByTestId } = render(
            <UpdateCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} updatedData={updatedData} />
        );

        // Нажимаем кнопку, чтобы инициировать запрос
        const updateCrewButton = getByTestId("update-crew-button");
        await act(async () => {
            updateCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Произошла неизвестная ошибка.");
    });
});
