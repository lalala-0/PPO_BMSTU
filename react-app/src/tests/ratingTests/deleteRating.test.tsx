import { render, screen, fireEvent } from "@testing-library/react";
import { useDeleteRatingController } from "../../ui/controllers/ratingControllers/deleteRatingController"; // Путь может отличаться
import api from "../api"; // Импортируем API
import React from "react";

jest.mock("../api");

const DeleteRatingWrapper = ({ id }: { id: string }) => {
    const { handleDelete } = useDeleteRatingController();

    return (
        <div>
            <button onClick={() => handleDelete(id)} data-testid="delete-button">
                Delete Rating
            </button>
        </div>
    );
};

describe("useDeleteRatingController", () => {
    beforeEach(() => {
        jest.clearAllMocks();
        jest.spyOn(window, "alert").mockImplementation(() => {}); // Мокаем alert
    });

    test("успешное удаление рейтинга", async () => {
        // Мокаем успешный ответ API
        (api.delete as jest.Mock).mockResolvedValue({});

        render(<DeleteRatingWrapper id="12345" />);

        // Нажимаем на кнопку удаления рейтинга
        fireEvent.click(screen.getByTestId("delete-button"));

        // Проверяем, что был вызван alert с успешным сообщением
        expect(api.delete).toHaveBeenCalledWith("/ratings/12345/");
    });

    test("ошибка при удалении рейтинга", async () => {
        // Мокаем ошибку API
        (api.delete as jest.Mock).mockRejectedValue(new Error("Ошибка при удалении"));

        render(<DeleteRatingWrapper id="12345" />);

        // Нажимаем на кнопку удаления рейтинга
        fireEvent.click(screen.getByTestId("delete-button"));
    });

    test("удаление без ID", async () => {
        render(<DeleteRatingWrapper id="" />);

        // Нажимаем на кнопку удаления рейтинга
        fireEvent.click(screen.getByTestId("delete-button"));

        // Проверяем, что был вызван alert с сообщением о неуказанном ID
        expect(window.alert).toHaveBeenCalledWith("ID рейтинга не указан");
    });
});
