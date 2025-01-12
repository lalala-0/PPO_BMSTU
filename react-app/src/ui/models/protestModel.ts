// protestModel.ts

export interface ProtestFormData {
  id: string;
  judgeId: string;
  ratingId: string;
  raceId: string;
  ruleNum: number;
  reviewDate: string;
  status: string;
  comment: string;
}

export function fromProtestModelToStringData(protest: any): ProtestFormData {
  const status = protestStatusToString(protest.status);
  return {
    id: protest.id,
    judgeId: protest.judgeId,
    ratingId: protest.ratingId,
    raceId: protest.raceId,
    ruleNum: protest.ruleNum,
    reviewDate: new Date(protest.reviewDate).toISOString(),
    status,
    comment: protest.comment,
  };
}

export function fromProtestModelsToStringData(
  protests: any[],
): ProtestFormData[] {
  return protests.map((protest) => fromProtestModelToStringData(protest));
}

export interface ProtestInput {
  judgeId: string;
  ruleNum: number;
  reviewDate: string;
  status: number;
  comment: string;
}

export function fromProtestModelToInputData(protest: any): ProtestInput {
  return {
    judgeId: protest.judgeId,
    ruleNum: protest.ruleNum,
    reviewDate: new Date(protest.reviewDate).toISOString(),
    status: protest.status,
    comment: protest.comment,
  };
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

// Helper function to map protest status to string
function protestStatusToString(status: number): string {
  const statusMap: Record<number, string> = {
    0: "Pending",
    1: "Approved",
    2: "Rejected",
  };
  return statusMap[status] || "Unknown";
}

export interface ProtestCrewFormData {
  sailNum: number;
  role: number;
  teamName: string; // Например, название команды
  protestID: string; // ID протеста, к которому относится команда
}
