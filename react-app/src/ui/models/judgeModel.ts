// judgeModel.ts

export interface JudgeFormData {
  id: string;
  fio: string;
  login: string;
  role: string;
  post: string;
}

// Интерфейс для фильтров
export interface JudgeFilters {
  fio?: string;
  login?: string;
  role?: string;
  post?: string;
}

export function fromJudgeModelToStringData(judge: models.Judge): JudgeFormData {
  const role = modelTables.judgeRoleToString(judge.role);
  return {
    id: judge.id,
    fio: judge.fio,
    login: judge.login,
    role,
    post: judge.post,
  };
}

export function fromJudgeModelsToStringData(
  judges: models.Judge[],
): JudgeFormData[] {
  return judges.map((judge) => fromJudgeModelToStringData(judge));
}

export interface PasswordInput {
  password: string;
}

export interface JudgeInput {
  id: string;
  fio: string;
  login: string;
  password?: string;
  role?: number;
  post?: string;
}

export function fromJudgeModelToInputData(judge: models.Judge): JudgeInput {
  return {
    id: judge.id,
    fio: judge.fio,
    login: judge.login,
    password: judge.password,
    role: judge.role,
    post: judge.post,
  };
}

export function fromJudgeModelsToInputData(
  judges: models.Judge[],
): JudgeInput[] {
  return judges.map((judge) => fromJudgeModelToInputData(judge));
}

export const JudgeRoleMap: Record<number, string> = {
  1: "Главный судья",
  2: "Судья",
};

// Example implementation of models and modelTables (adjust as needed)
namespace models {
  export interface Judge {
    id: string;
    fio: string;
    login: string;
    password: string;
    role: number;
    post: string;
  }
}

namespace modelTables {
  export function judgeRoleToString(role: number): string {
    return JudgeRoleMap[role] || "Неизвестная роль";
  }
}
