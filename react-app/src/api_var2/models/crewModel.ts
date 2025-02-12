// crewModel.ts

// Типы для данных Crew
export interface CrewFormData {
  id: string; // UUID в виде строки
  ratingId: string; // UUID в виде строки
  SailNum: number;
  Class: string;
}

// Типы для данных ProtestCrew
export interface ProtestCrewFormData extends CrewFormData {
  role: string;
}

// Типы для ввода данных Crew
export interface CrewInput {
  sailNum: number;
}

export interface CrewParticipantDetachInput {
  participantID: string;
}

export interface CrewParticipantAttachInput {
  participantID: string;
  helmsman?: number;
}

// HelmsmanMap для удобства отображения
export const HelmsmanMap: Record<number, string> = {
  0: "Не рулевой",
  1: "Рулевой",
};
