// protestModel.ts

export interface ProtestFormData {
  ID: string;
  JudgeID: string;
  RatingID: string;
  RaceID: string;
  RuleNum: number;
  ReviewDate: string;
  Status: string;
  Comment: string;
}

export interface ProtestInput {
  judgeId: string;
  ruleNum: number;
  reviewDate: string;
  status: number;
  comment: string;
}

export interface ProtestCreate {
  judgeId: string;
  ruleNum: number;
  reviewDate: string;
  comment: string;
  protesteeSailNum: number;
  protestorSailNum: number;
  witnessesSailNum: number[];
}

export interface ProtestParticipantDetachInput {
  sailNum: number;
}

export interface ProtestParticipantAttachInput {
  sailNum: number;
  role: number;
}

export const RoleMap: Record<number, string> = {
  1: "Протестующий",
  2: "Опротестованный",
  3: "Свидетель",
};

export interface ProtestComplete {
  resPoints: number;
  comment: string;
}

export const StatusMap: Record<number, string> = {
  1: "Ожидает рассмотрения",
  2: "Рассмотрен",
};

export interface ProtestCrewFormData {
  sailNum: number;
  role: number;
  teamName: string; // Например, название команды
  protestID: string; // ID протеста, к которому относится команда
}
