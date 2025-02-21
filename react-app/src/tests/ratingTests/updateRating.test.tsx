import { render, act, screen } from "@testing-library/react";
import { useUpdateRatingController } from "../../ui/controllers/ratingControllers/updateRatingController"; // Замените путь на правильный
import api from "../api";
import React from "react";

// Компонент-обертка для тестирования хука
const UpdateRatingWrapper = ({ ratingID }: { ratingID: string }) => {
    const { success, loading, handleUpdate } = useUpdateRatingController();

    return (
        <div>
            <button
                onClick={() =>
                    handleUpdate(ratingID, {
                        name: "Иван Иванов",
                        class: 1,
                        blowout_cnt: 3,
                    })
                }
                data-testid="update-button"
            >
                Update Rating
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success}</div>
        </div>
    );
};

describe("useUpdateRatingController", () => {
    const ratingID = "rating123";

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное обновление рейтинга", async () => {
        const mockUpdatedRating = {
            id: ratingID,
            name: "Иван Иванов",
            class: 1,
            blowout_cnt: 3,
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedRating });

        render(<UpdateRatingWrapper ratingID={ratingID} />);

        // Нажимаем кнопку для обновления данных рейтинга
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после успешного обновления
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("Рейтинг успешно обновлён");
        expect(api.put).toHaveBeenCalledWith(`/ratings/${ratingID}`, {
            name: "Иван Иванов",
            class: 1,
            blowout_cnt: 3,
        });
    });

    test("ошибка при обновлении рейтинга", async () => {
        // Мокаем ошибку API
        (api.put as jest.Mock).mockRejectedValue(new Error("Ошибка обновления рейтинга"));

        render(<UpdateRatingWrapper ratingID={ratingID} />);

        // Нажимаем кнопку для обновления данных рейтинга
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
        });

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("");
    });

    test("проверка обновления при повторном вызове", async () => {
        const mockUpdatedRating = {
            id: ratingID,
            name: "Иван Иванов",
            class: 1,
            blowout_cnt: 3,
        };

        // Мокаем успешный ответ API
        (api.put as jest.Mock).mockResolvedValue({ data: mockUpdatedRating });

        render(<UpdateRatingWrapper ratingID={ratingID} />);

        // Нажимаем кнопку для обновления данных рейтинга дважды
        const updateButton = screen.getByTestId("update-button");
        await act(async () => {
            updateButton.click();
            updateButton.click(); // Дважды для проверки
        });

        expect(api.put).toHaveBeenCalledTimes(2);
    });
});
