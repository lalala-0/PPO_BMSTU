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

// Преобразование CrewModel в CrewFormData
export function fromCrewModelToStringData(
  crew: { id: string; ratingId: string; sailNum: number; class: string },
  classToString: (cls: string) => string,
): CrewFormData {
  return {
    id: crew.id,
    ratingId: crew.ratingId,
    SailNum: crew.sailNum,
    Class: classToString(crew.class),
  };
}

// Преобразование массива CrewModel в массив CrewFormData
export function fromCrewModelsToStringData(
  crews: { id: string; ratingId: string; sailNum: number; class: string }[],
  classToString: (cls: string) => string,
): CrewFormData[] {
  return crews.map((crew) => fromCrewModelToStringData(crew, classToString));
}

// Преобразование ProtestParticipantModel в ProtestCrewFormData
export function fromProtestParticipantModelToStringData(
  crew: { id: string; ratingId: string; sailNum: number; class: string },
  role: number,
  classToString: (cls: string) => string,
  roleToString: (role: number) => string,
): ProtestCrewFormData {
  return {
    id: crew.id,
    ratingId: crew.ratingId,
    SailNum: crew.sailNum,
    Class: classToString(crew.class),
    role: roleToString(role),
  };
}

// Преобразование массива ProtestParticipantModel в массив ProtestCrewFormData
export function fromProtestParticipantModelsToStringData(
  crews: { id: string; ratingId: string; sailNum: number; class: string }[],
  roles: number[],
  classToString: (cls: string) => string,
  roleToString: (role: number) => string,
): ProtestCrewFormData[] {
  return crews.map((crew, index) =>
    fromProtestParticipantModelToStringData(
      crew,
      roles[index],
      classToString,
      roleToString,
    ),
  );
}

// Преобразование CrewModel в CrewInput
export function fromCrewModelToInputData(crew: { sailNum: number }): CrewInput {
  return { sailNum: crew.sailNum };
}

// Преобразование ProtestParticipantModel в CrewInput
export function fromProtestParticipantModelToInputData(
  crew: { sailNum: number },
  role: number,
): CrewInput {
  return { sailNum: crew.sailNum };
}
