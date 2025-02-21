import api from "../api";
import {RankingResponse, Rating} from "../../ui";
import {fetchRankingData, fetchRatingInfo} from "../../ui/controllers/ratingControllers/rankingRatingController"; // Импортируем API

jest.mock("../api");

describe("RatingService", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешный запрос информации о рейтинге", async () => {
        const mockRatingData: Rating = {
            id: "123",
            Name: "rating",
            Class: "Laser",
            BlowoutCnt: 1,
        };

        // Мокаем успешный ответ от API
        (api.get as jest.Mock).mockResolvedValue({ data: mockRatingData });

        const rating = await fetchRatingInfo("123");

        expect(rating).toEqual(mockRatingData);
        expect(api.get).toHaveBeenCalledWith("/ratings/123");
    });

    test("ошибка при запросе информации о рейтинге", async () => {
        // Мокаем ошибку от API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка запроса"));

        // Проверяем, что ошибка будет выброшена
        await expect(fetchRatingInfo("123")).rejects.toThrow("Ошибка запроса");
        expect(api.get).toHaveBeenCalledWith("/ratings/123");
    });

    test("успешный запрос данных о рангах", async () => {
        const mockRankingData: RankingResponse = {
            RankingTable: [
                {
                    crewID: "crew1",
                    SailNum: 101,
                    ParticipantNames: ["John Doe", "Jane Smith"],
                    ParticipantBirthDates: ["2000-01-01", "1998-05-23"],
                    ResInRace: { "race1": 1, "race2": 2 },
                    PointsSum: 15,
                    Rank: 1,
                    CoachNames: ["Coach A"],
                },
            ],
            Races: [
                { RaceNum: 1, RaceID: "race1" },
                { RaceNum: 2, RaceID: "race2" },
            ],
        };

        // Мокаем успешный ответ от API
        (api.get as jest.Mock).mockResolvedValue({ data: mockRankingData });

        const ranking = await fetchRankingData("123");

        expect(ranking).toEqual(mockRankingData);
        expect(api.get).toHaveBeenCalledWith("/ratings/123/rankings");
    });

    test("ошибка при запросе данных о рангах", async () => {
        // Мокаем ошибку от API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка запроса"));

        // Проверяем, что ошибка будет выброшена
        await expect(fetchRankingData("123")).rejects.toThrow("Ошибка запроса");
        expect(api.get).toHaveBeenCalledWith("/ratings/123/rankings");
    });
});
