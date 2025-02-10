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

export const JudgeRoleMap: Record<number, string> = {
  1: "Главный судья",
  2: "Судья",
};
