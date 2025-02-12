// participantModel.ts

export interface ParticipantFormData {
  id: string;
  FIO: string;
  Category: string;
  Gender: string;
  Birthday: string;
  Coach: string;
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
  0: "Муж.",
  1: "Жен.",
};
