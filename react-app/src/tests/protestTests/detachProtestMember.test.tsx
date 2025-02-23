import { renderHook, act } from "@testing-library/react";
import api from "../api";
import { handleError } from "../errorHandler";
import {ProtestParticipantDetachInput} from "../../ui/models/protestModel";
import {useDetachProtestMember} from "../../ui/controllers/protestControllers/detachProtestMemberController";

jest.mock("../api");
jest.mock("../errorHandler");

describe("useDetachProtestMember", () => {
    const ratingID = "rating123";
    const raceID = "race456";
    const protestID = "protest789";
    const protestParticipantDetachInput: ProtestParticipantDetachInput = {
        ratingID: "rating123",
        sailNum: 101,
    };

    beforeEach(() => {
        jest.clearAllMocks();
    });

    test("успешное открепление участника", async () => {
        (api.delete as jest.Mock).mockResolvedValue({});

        const { result } = renderHook(() =>
            useDetachProtestMember(ratingID, raceID, protestID)
        );

        await act(async () => {
            await result.current.detachProtestMember(protestParticipantDetachInput);
        });

        expect(api.delete).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members/${protestParticipantDetachInput.sailNum}`
        );
        expect(result.current.error).toBeNull();
        expect(result.current.loading).toBe(false);
    });

    test("ошибка при откреплении участника", async () => {
        const error = new Error("Ошибка удаления");
        (api.delete as jest.Mock).mockRejectedValue(error);

        const { result } = renderHook(() =>
            useDetachProtestMember(ratingID, raceID, protestID)
        );

        await act(async () => {
            await result.current.detachProtestMember(protestParticipantDetachInput);
        });

        expect(api.delete).toHaveBeenCalledWith(
            `/ratings/${ratingID}/races/${raceID}/protests/${protestID}/members/${protestParticipantDetachInput.sailNum}`
        );
        expect(handleError).toHaveBeenCalledWith(error, expect.any(Function));
        expect(result.current.loading).toBe(false);
    });
});
