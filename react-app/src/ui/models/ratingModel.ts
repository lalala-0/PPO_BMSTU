// ratingModel.ts
export interface Rating {
  id: string;
  Name: string;
  Class: string;
  BlowoutCnt: number;
}

export interface RatingInput {
  name: string;
  class: number;
  blowout_cnt: number;
}

export interface RatingTableLine {
  crewID: string;
  SailNum: number;
  ParticipantNames: string[];
  ParticipantBirthDates: string[];
  ResInRace: { [key: string]: number };
  PointsSum: number;
  Rank: number;
  CoachNames: string[];
}

export interface RaceInfo {
  RaceNum: number;
  RaceID: string;
}

export interface RankingResponse {
  RankingTable: RatingTableLine[];
  Races: RaceInfo[];
}
