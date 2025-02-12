import { render, act } from "@testing-library/react";
import { useGetCrew } from "../../controllers/crewControllers/getCrewController";
import api from "../../controllers/api"; // Путь к мокированному API

jest.mock("../../controllers/api"); // Мокаем api.get

const GetCrewWrapper = ({ ratingID, crewID }: { ratingID: string; crewID: string }) => {
    const { crewInfo, loading, error, getCrew } = useGetCrew();

    return (
        <div>
            <button onClick={() => getCrew(ratingID, crewID)} data-testid="get-crew-button">
                Get Crew
            </button>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="success">{crewInfo ? "Success" : "Failure"}</div>
            <div data-testid="error">{error}</div>
        </div>
    );
};

describe("useGetCrew", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение данных о команде", async () => {
        const mockRatingID = "123";
        const mockCrewID = "crew123";
        const mockCrewData = {
            crewID: mockCrewID,
            name: "Test Crew",
            members: 5,
            sailNum: 7,
        };

        (api.get as jest.Mock).mockResolvedValueOnce({ data: mockCrewData });

        const { getByTestId } = render(<GetCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewButton = getByTestId("get-crew-button");
        await act(async () => {
            getCrewButton.click();
        });

        // Проверяем, что данные о команде отображаются
        expect(getByTestId("success").textContent).toBe("Success");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith(`/ratings/${mockRatingID}/crews/${mockCrewID}`);
    });

    test("ошибка при получении данных о команде", async () => {
        const mockRatingID = "123";
        const mockCrewID = "crew123";

        (api.get as jest.Mock).mockRejectedValueOnce({
            message: "Ошибка получения данных",
        });

        const { getByTestId } = render(<GetCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewButton = getByTestId("get-crew-button");
        await act(async () => {
            getCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения данных");
    });

    test("ошибка при сетевом запросе", async () => {
        const mockRatingID = "123";
        const mockCrewID = "crew123";

        (api.get as jest.Mock).mockRejectedValueOnce({ request: {} }); // Мокаем ошибку сети

        const { getByTestId } = render(<GetCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewButton = getByTestId("get-crew-button");
        await act(async () => {
            getCrewButton.click();
        });

        // Проверяем, что ошибка сети отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Ошибка сети. Проверьте подключение к интернету.");
    });

    test("непредвиденная ошибка при запросе", async () => {
        const mockRatingID = "123";
        const mockCrewID = "crew123";

        (api.get as jest.Mock).mockRejectedValueOnce(null); // Мокаем непредвиденную ошибку

        const { getByTestId } = render(<GetCrewWrapper ratingID={mockRatingID} crewID={mockCrewID} />);

        // Нажимаем кнопку, чтобы инициировать запрос
        const getCrewButton = getByTestId("get-crew-button");
        await act(async () => {
            getCrewButton.click();
        });

        // Проверяем, что ошибка отображается
        expect(getByTestId("success").textContent).toBe("Failure");
        expect(getByTestId("loading").textContent).toBe("not loading");
        expect(getByTestId("error").textContent).toBe("Произошла неизвестная ошибка.");
    });
});
