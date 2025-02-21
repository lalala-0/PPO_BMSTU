import { render, act, screen } from "@testing-library/react";
import { useCreateRatingController } from "../../ui/controllers/ratingControllers/createRatingController"; // Путь может отличаться
import api from "../api";
import React from "react";

// Компонент-обертка для тестирования хука
const CreateRatingWrapper = ({ ratingData }: { ratingData: any }) => {
    const { input, success, loading, handleSubmit } = useCreateRatingController();

    const handleCreate = () => handleSubmit(ratingData);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success}</div>
            <div data-testid="created">
                {JSON.stringify(input)}
            </div>
            <button onClick={handleCreate} data-testid="create-button">
                Create Rating
            </button>
        </div>
    );
};

describe("useCreateRatingController", () => {
    const ratingData = {
        name: "Good",
        class: 1,
        blowout_cnt: 5,
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное создание рейтинга", async () => {
        // Мокаем успешный ответ API
        (api.post as jest.Mock).mockResolvedValue({
            data: {
                ...ratingData,
                id: "12345",
            },
        });

        render(<CreateRatingWrapper ratingData={ratingData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("");

        // Нажимаем на кнопку создания рейтинга
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения процесса
        await act(async () => {});

        // Проверяем успешное создание рейтинга
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("Рейтинг успешно создан");
        expect(api.post).toHaveBeenCalledWith("/ratings/", ratingData);
    });

    test("ошибка при создании рейтинга", async () => {
        // Мокаем ошибку API
        (api.post as jest.Mock).mockRejectedValue(new Error("Ошибка при создании рейтинга"));

        render(<CreateRatingWrapper ratingData={ratingData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("");

        // Нажимаем на кнопку создания рейтинга
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("success").textContent).toBe("");
    });

    test("проверка состояния загрузки", async () => {
        // Мокаем успешный ответ API
        (api.post as jest.Mock).mockResolvedValue({
            data: {
                ...ratingData,
                id: "12345",
            },
        });

        render(<CreateRatingWrapper ratingData={ratingData} />);

        // Проверяем начальное состояние
        expect(screen.getByTestId("loading").textContent).toBe("not loading");

        // Нажимаем на кнопку создания рейтинга
        act(() => {
            screen.getByTestId("create-button").click();
        });

        // Проверяем состояние загрузки во время выполнения запроса
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем завершения процесса
        await act(async () => {});

        // Проверяем состояние после завершения запроса
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
    });
});
