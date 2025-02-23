import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import api from "../api";
import {useDeleteRace} from "../../ui/controllers/raceControllers/deleteRaceController";
import React from "react";

jest.mock("../api");
jest.mock("../errorHandler");

const MockComponent = ({ ratingID, raceID }: { ratingID: string, raceID: string }) => {
    const { loading, error, success, deleteRace } = useDeleteRace(ratingID, raceID);

    return (
        <div>
            <button onClick={deleteRace}>Удалить гонку</button>
            {loading && <div>Загрузка...</div>}
            {error && <div>{error}</div>}
            {success && <div>Гонка удалена успешно</div>}
        </div>
    );
};

describe("useDeleteRace", () => {
    it("должен инициировать загрузку при удалении гонки", async () => {
        const mockDelete = jest.fn().mockResolvedValue({});
        api.delete = mockDelete;

        render(<MockComponent ratingID="123" raceID="456" />);

        fireEvent.click(screen.getByText("Удалить гонку"));
        expect(screen.getByText("Загрузка...")).toBeInTheDocument();

        await waitFor(() => expect(mockDelete).toHaveBeenCalledWith("/ratings/123/races/456"));
    });

    it("должен показывать сообщение об успешном удалении", async () => {
        const mockDelete = jest.fn().mockResolvedValue({});
        api.delete = mockDelete;

        render(<MockComponent ratingID="123" raceID="456" />);

        fireEvent.click(screen.getByText("Удалить гонку"));

        await waitFor(() => expect(screen.getByText("Гонка удалена успешно")).toBeInTheDocument());
    });

});
