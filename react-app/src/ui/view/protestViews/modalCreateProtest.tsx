import React, { useState, useEffect } from "react";
import { useCreateProtest } from "../../controllers/protestControllers/createProtestController";
import { useGetJudges } from "../../controllers/judgeControllers/getJudgesController";
import { ProtestCreate } from "../../models/protestModel";
import { useParams } from "react-router-dom";
import { useGetCrewsByRatingID } from "../../controllers/crewControllers/getCrewsController";

interface ProtestModalProps {
    onClose: () => void;
}

const ProtestModal: React.FC<ProtestModalProps> = ({ onClose }) => {
    const { createProtest, loading } = useCreateProtest();
    const { judges } = useGetJudges();
    const { ratingID } = useParams<{ ratingID: string }>();
    const { raceID } = useParams<{ raceID: string }>();
    const { crews, getCrewsByRatingID } = useGetCrewsByRatingID();

    const [formData, setFormData] = useState<ProtestCreate>({
        judgeId: "",
        ruleNum: 0,
        reviewDate: "",
        comment: "",
        protestee: 0,
        protestor: 0,
        witnesses: [],
    });

    useEffect(() => {
        if (ratingID) {
            getCrewsByRatingID(ratingID);
        }
    }, [ratingID, getCrewsByRatingID]);

    const handleChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
    ) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: ["ruleNum", "protestee", "protestor"].includes(name)
                ? Number(value) // Преобразуем в число
                : value,
        }));
    };

    const handleWitnessChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedValues = Array.from(e.target.selectedOptions, (option) =>
            Number(option.value) // Преобразуем строки обратно в числа
        );
        setFormData((prev) => ({
            ...prev,
            witnesses: selectedValues,
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        await createProtest(ratingID || '', raceID || '', formData);
        onClose();
    };

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h2>Создать протест</h2>
                <form onSubmit={handleSubmit}>
                    <label>
                        Судья:
                        <select name="judgeId" value={formData.judgeId} onChange={handleChange} required>
                            <option value="">Выберите судью</option>
                            {judges.map((judge) => (
                                <option key={judge.id} value={judge.id}>{judge.fio}</option>
                            ))}
                        </select>
                    </label>
                    <label>
                        Номер правила:
                        <input type="number" name="ruleNum" value={formData.ruleNum} onChange={handleChange} required />
                    </label>
                    <label>
                        Дата рассмотрения:
                        <input type="date" name="reviewDate" value={formData.reviewDate} onChange={handleChange} required />
                    </label>
                    <label>
                        Комментарий:
                        <input type="text" name="comment" value={formData.comment} onChange={handleChange} required />
                    </label>
                    <label>
                        Протестуемый парусный номер:
                        <select name="protestee" value={formData.protestee} onChange={handleChange} required>
                            <option value="">Выберите номер</option>
                            {crews?.map((crew) => (
                                <option key={crew.id} value={crew.SailNum}>{crew.SailNum}</option>
                            ))}
                        </select>
                    </label>
                    <label>
                        Протестующий парусный номер:
                        <select name="protestor" value={formData.protestor} onChange={handleChange} required>
                            <option value="">Выберите номер</option>
                            {crews?.map((crew) => (
                                <option key={crew.id} value={crew.SailNum}>{crew.SailNum}</option>
                            ))}
                        </select>
                    </label>

                    <label>
                        Свидетели:
                        <select
                            name="witnessesSailNum"
                            multiple
                            value={formData.witnesses.map(String)} // Сохраняем выбор в виде массива строк
                            onChange={handleWitnessChange}
                        >
                            {crews?.map((crew) => (
                                <option key={crew.id} value={crew.SailNum}>
                                    {crew.SailNum}
                                </option>
                            ))}
                        </select>
                    </label>

                    {/* Отображение выбранных свидетелей */}
                    {formData.witnesses.length > 0 && (
                        <div>
                            <label>Выбранные свидетели:</label>
                            <ul>
                                {formData.witnesses.map((witnessId) => {
                                    const crew = crews?.find((c) => c.SailNum === witnessId);
                                    return <li key={witnessId}>{crew?.SailNum}</li>;
                                })}
                            </ul>
                        </div>
                    )}


                    <div className="modal-buttons">
                        <button type="submit" disabled={loading}>Создать</button>
                        <button type="button" onClick={onClose}>Отмена</button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default ProtestModal;
