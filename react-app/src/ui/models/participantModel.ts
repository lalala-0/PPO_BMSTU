// participantModel.ts

export interface ParticipantFormData {
  id: string;
  fio: string;
  category: string;
  gender: string;
  birthday: string;
  coach: string;
}

export interface ParticipantInput {
  id?: string;
  fio: string;
  category: number;
  gender?: number;
  birthday: string;
  coach: string;
}

export interface ParticipantFilters {
  fio?: string;
  category?: string;
  gender?: string;
  birthday?: string;
  coach?: string;
}

export const CategoryMap: Record<number, string> = {
  1: "Мастер спорта России международного класса",
  2: "Мастер спорта России",
  3: "Кандидат в мастера спорта",
  4: "1-ый спортивный разряд",
  5: "2-ой спортивный разряд",
  6: "3-ий спортивный разряд",
  7: "1-ый юношеский разряд",
  8: "2-ой юношеский разряд",
};

export const GenderMap: Record<number, string> = {
  1: "Муж.",
  2: "Жен.",
};

function participantCategoryToString(category: number): string {
  return CategoryMap[category] || "Неизвестная категория";
}

function genderToString(gender: number): string {
  return GenderMap[gender] || "Неизвестный пол";
}

function fromParticipantModelToStringData(participant: {
  id: string;
  fio: string;
  category: number;
  gender: number;
  birthday: Date;
  coach: string;
}): ParticipantFormData {
  const category = participantCategoryToString(participant.category);
  const gender = genderToString(participant.gender);
  return {
    id: participant.id,
    fio: participant.fio,
    category,
    gender,
    birthday: participant.birthday.toISOString().split("T")[0],
    coach: participant.coach,
  };
}

function fromParticipantModelsToStringData(
  participants: {
    id: string;
    fio: string;
    category: number;
    gender: number;
    birthday: Date;
    coach: string;
  }[],
): ParticipantFormData[] {
  return participants.map(fromParticipantModelToStringData);
}

function fromParticipantFormDataToInputData(participant: {
  id: string;
  fio: string;
  category: string;
  gender: string;
  birthday: Date;
  coach: string;
}): ParticipantInput {
  return {
    id: participant.id,
    fio: participant.fio,
    category: participant.category ? parseInt(participant.category) : 0, // Если категория не найдена, ставим 0
    gender: participant.gender ? parseInt(participant.gender) : 0, // Если пол не найден, ставим 0
    birthday: participant.birthday.toISOString().split("T")[0],
    coach: participant.coach,
  };
}
