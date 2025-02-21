import api from "../api"; // Путь к твоему мокированному API
import { act } from "react-dom/test-utils";
import { render } from "@testing-library/react";
import {useGetCrewsByRatingID} from "../../ui/controllers/crewControllers/getCrewsController";
import React from "react";

// Мокаем api.get
jest.mock("../api");

const GetCrewsWrapper = ({ ratingID }: { ratingID: string }) => {
    const { loading, error, crews, getCrewsByRatingID } = useGetCrewsByRatingID();

    // Используем useGetCrewsByRatingID внутри компонента
    return (
        <div>
            <button onClick={() => getCrewsByRatingID(ratingID)} data-testid="get-crews-button">
                Get Crews
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{crews ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("getCrewsByRatingID controller", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение данных о командах", async () => {
        const mockRatingID = "123";
        const mockCrews = [
            { id: "1", name: "Crew 1" },
            { id: "2", name: "Crew 2" },
        ];
        (api.get as jest.Mock).mockResolvedValue({ data: mockCrews });

        const { getByTestId } = render(<GetCrewsWrapper ratingID={mockRatingID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewsButton = getByTestId("get-crews-button");
        await act(async () => {
            getCrewsButton.click();
        });

        // Проверяем, что данные команд отображаются
        expect(getByTestId("success").textContent).toBe("Success");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/ratings/${mockRatingID}/crews`);
    });

    test("ошибка при получении данных о командах", async () => {
        const mockRatingID = "123";
        (api.get as jest.Mock).mockRejectedValue({
            message: "Ошибка получения данных",
        });

        const { getByTestId } = render(<GetCrewsWrapper ratingID={mockRatingID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewsButton = getByTestId("get-crews-button");
        await act(async () => {
            getCrewsButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения данных");
    });

    test("ошибка при сетевом запросе", async () => {
        const mockRatingID = "123";
        (api.get as jest.Mock).mockRejectedValue({ request: {} }); // Мокаем ошибку сети

        const { getByTestId } = render(<GetCrewsWrapper ratingID={mockRatingID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewsButton = getByTestId("get-crews-button");
        await act(async () => {
            getCrewsButton.click();
        });

        // Проверяем, что ошибка сети отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка сети. Проверьте подключение к интернету.");
    });

    test("непредвиденная ошибка при запросе", async () => {
        const mockRatingID = "123";
        (api.get as jest.Mock).mockRejectedValue(null); // Мокаем непредвиденную ошибку

        const { getByTestId } = render(<GetCrewsWrapper ratingID={mockRatingID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewsButton = getByTestId("get-crews-button");
        await act(async () => {
            getCrewsButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Произошла неизвестная ошибка.");
    });
});
