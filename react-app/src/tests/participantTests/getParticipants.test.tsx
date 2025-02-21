import { render, act, screen } from "@testing-library/react";
import api from "../api";
import { ParticipantFormData, ParticipantFilters } from "../../ui/models/participantModel";
import React from "react";
import {useGetAllParticipants} from "../../ui/controllers/participantControllers/getParticipantsController";

// Компонент-обертка для тестирования хука
const GetParticipantsWrapper = ({ filters }: { filters: ParticipantFilters }) => {
    const { getParticipants, loading, error, participants } = useGetAllParticipants(filters);

    React.useEffect(() => {
        getParticipants();
    }, [filters]);

    return (
        <div>
            <div data-testid="loading">{loading ? "loading..." : "not loading"}</div>
            <div data-testid="error">{error}</div>
            <div data-testid="participants">
                {participants.length > 0 ? JSON.stringify(participants) : "No participants found"}
            </div>
        </div>
    );
};

describe("useGetAllParticipants controller", () => {
    const filters: ParticipantFilters = {
        category: "1",
        gender: "",
        coach: "Ирина Петрова",
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное получение участников", async () => {
        const mockParticipants: ParticipantFormData[] = [
            {
                id: "1",
                FIO: "Иван Иванов",
                Category: "1",
                Birthday: "1990-01-01",
                Coach: "Ирина Петрова",
                Gender: "1",
            },
            {
                id: "2",
                FIO: "Анна Петрова",
                Category: "2",
                Birthday: "1992-03-03",
                Coach: "Ирина Петрова",
                Gender: "2",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockParticipants });

        render(<GetParticipantsWrapper filters={filters} />);

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем получения участников
        await act(async () => {});

        // Проверяем, что участники отображаются корректно
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participants").textContent).toBe(
            JSON.stringify(mockParticipants)
        );
        expect(screen.getByTestId("error").textContent).toBe("");
        expect(api.get).toHaveBeenCalledWith("/participants", {
            params: { category: "1", coach: "Ирина Петрова" },
        });
    });

    test("ошибка при получении участников", async () => {
        // Мокаем ошибку API
        (api.get as jest.Mock).mockRejectedValue(new Error("Ошибка получения участников"));

        render(<GetParticipantsWrapper filters={filters} />);

        // Проверяем состояние загрузки
        expect(screen.getByTestId("loading").textContent).toBe("loading...");

        // Ожидаем появления ошибки
        await act(async () => {});

        // Проверяем состояние после ошибки
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participants").textContent).toBe("No participants found");
        expect(screen.getByTestId("error").textContent).toBe("Ошибка запроса: Ошибка получения участников");
    });

    test("проверка с пустыми фильтрами", async () => {
        const filtersEmpty: ParticipantFilters = {};

        const mockParticipants: ParticipantFormData[] = [
            {
                id: "1",
                FIO: "Иван Иванов",
                Category: "1",
                Birthday: "1990-01-01",
                Coach: "Ирина Петрова",
                Gender: "1",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockParticipants });

        render(<GetParticipantsWrapper filters={filtersEmpty} />);

        // Ожидаем получение участников без фильтров
        await act(async () => {});

        // Проверяем, что запрос был отправлен с пустыми фильтрами
        expect(api.get).toHaveBeenCalledWith("/participants", { params: {} });

        // Проверяем корректное отображение полученных участников
        expect(screen.getByTestId("loading").textContent).toBe("not loading");
        expect(screen.getByTestId("participants").textContent).toBe(
            JSON.stringify(mockParticipants)
        );
    });

    test("проверка обновления фильтров", async () => {
        const initialFilters: ParticipantFilters = { category: "1", coach: "Ирина Петрова" };
        const updatedFilters: ParticipantFilters = { category: "2", coach: "Александр Смирнов" };

        const mockParticipants: ParticipantFormData[] = [
            {
                id: "1",
                FIO: "Иван Иванов",
                Category: "1",
                Birthday: "1990-01-01",
                Coach: "Ирина Петрова",
                Gender: "1",
            },
        ];

        // Мокаем успешный ответ API
        (api.get as jest.Mock).mockResolvedValue({ data: mockParticipants });

        render(<GetParticipantsWrapper filters={initialFilters} />);

        // Проверяем, что фильтры были отправлены корректно
        expect(api.get).toHaveBeenCalledWith("/participants", { params: initialFilters });

        // Обновляем фильтры
        render(<GetParticipantsWrapper filters={updatedFilters} />);

        // Проверяем, что запрос был обновлен с новыми фильтрами
        expect(api.get).toHaveBeenCalledWith("/participants", { params: updatedFilters });
    });
});
