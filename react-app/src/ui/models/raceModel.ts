// raceModel.ts

export interface RaceFormData {
  id: string;
  ratingId: string;
  date: string;
  number: number;
  class: string;
}

export interface RaceInput {
  date: string;
  number: number;
  class: number;
}

export const SpecCircumstanceMap: Record<number, string> = {
  0: "-",
  1: "DNS",
  2: "DNF",
  3: "DNC",
  4: "OCS",
  5: "ZFP",
  6: "UFD",
  7: "BFD",
  8: "SCP",
  9: "RET",
  10: "DSQ",
  11: "DNE",
  12: "RDG",
  13: "DPI",
};

export interface StartInput {
  specCircumstance: number;
  falseStartList: number[];
}

export interface FinishInput {
  finisherList: number[];
}
