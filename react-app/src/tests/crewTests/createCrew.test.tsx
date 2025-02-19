import { render, fireEvent, waitFor } from "@testing-library/react";
import { useCreateCrew } from "../../ui/controllers/crewControllers/createCrewController";
import api from "../api";

jest.mock("../api"); // Мокаем api.get

const mockRatingID = "rating123";

// Компонент для тестирования хука
const CreateCrewComponent = () => {
    const {
        input,
        success,
        loading,
        error,
        handleChange,
        handleSubmit,
    } = useCreateCrew();

    return (
        <div>
            <input
                type="number"
                name="sailNum"
                value={input.sailNum}
                onChange={handleChange}
                data-testid="sailNum-input"
            />
            <button
                onClick={() => handleSubmit(mockRatingID, { sailNum: input.sailNum })}
                data-testid="submit-button"
            >
                Create Crew
            </button>
            {loading && <div data-testid="loading">Loading...</div>}
            {success && <div data-testid="success">{success}</div>}
            {error && <div data-testid="error">{error}</div>}
        </div>
    );
};

describe("useCreateCrew hook", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное создание команды", async () => {
        const mockResponse = { data: { id: "crew123" } };
        (api.post as jest.Mock).mockResolvedValue(mockResponse);

        const { getByTestId } = render(<CreateCrewComponent />);

        // Заполняем форму
        fireEvent.change(getByTestId("sailNum-input"), { target: { value: "5" } });

        // Нажимаем кнопку отправки
        const submitButton = getByTestId("submit-button");
        fireEvent.click(submitButton);

        // Проверяем, что состояния обновились
        expect(getByTestId("loading").textContent).toBe("Loading...");

        await waitFor(() => expect(getByTestId("success").textContent).toBe("Команда успешно создана"));
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews`,
            { sailNum: 5 }
        );
    });

    test("ошибка при создании команды", async () => {
        const errorMessage = "Ошибка создания команды";
        (api.post as jest.Mock).mockRejectedValue(new Error(errorMessage));

        const { getByTestId } = render(<CreateCrewComponent />);

        // Заполняем форму
        fireEvent.change(getByTestId("sailNum-input"), { target: { value: "5" } });

        // Нажимаем кнопку отправки
        const submitButton = getByTestId("submit-button");
        fireEvent.click(submitButton);

        // Проверяем, что состояние ошибки установлено
        expect(getByTestId("loading").textContent).toBe("Loading...");

        await waitFor(() =>
            expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка создания команды")
        );
        expect(api.post).toHaveBeenCalledWith(
            `/ratings/${mockRatingID}/crews`,
            { sailNum: 5 }
        );
    });

    test("изменение значения sailNum", () => {
        const { getByTestId } = render(<CreateCrewComponent />);

        // Проверяем начальное значение input
        const input = getByTestId("sailNum-input");

        // Меняем значение
        fireEvent.change(input, { target: { value: "10" } });
    });
});
