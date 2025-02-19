import api from "../api"; // Путь к твоему мокированному API
import { act } from "react-dom/test-utils";
import { render } from "@testing-library/react";
import {useDeleteJudge} from "../../ui/controllers/judgeControllers/deleteJudgeController";

// Мокаем api.delete
jest.mock("../api");

const DeleteJudgeWrapper = ({ judgeID }: { judgeID: string }) => {
    const { deleteJudge, loading, error, success } = useDeleteJudge();

    // Используем useDeleteJudge внутри компонента
    return (
        <div>
            <button onClick={() => deleteJudge(judgeID)} data-testid="delete-button">
                Delete Judge
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{success ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("deleteJudge controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное удаление судьи", async () => {
        const mockJudgeID = "judge123";
        (api.delete as jest.Mock).mockResolvedValue({});

        const { getByTestId } = render(<DeleteJudgeWrapper judgeID={mockJudgeID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteButton = getByTestId("delete-button");
        await act(async () => {
            deleteButton.click();
        });

        // Проверяем, что флаг success установился в true
        expect(getByTestId("success").textContent).toBe("Success");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
        expect(api.delete).toHaveBeenCalledWith(`/judges/${mockJudgeID}`);
    });

    test("ошибка при удалении судьи", async () => {
        const mockJudgeID = "judge123";
        (api.delete as jest.Mock).mockRejectedValue(new Error("Ошибка удаления"));

        const { getByTestId } = render(<DeleteJudgeWrapper judgeID={mockJudgeID} />);

        // Нажимаем кнопку, чтобы инициировать удаление
        const deleteButton = getByTestId("delete-button");
        await act(async () => {
            deleteButton.click();
        });

        // Проверяем, что флаг success установился в false и ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка удаления");
    });
});
