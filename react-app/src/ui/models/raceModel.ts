// raceModel.ts

export interface RaceFormData {
  id: string;
  ratingId: string;
  date: string;
  number: number;
  class: string;
}

function fromRaceModelToStringData(
  race: {
    id: string;
    ratingId: string;
    date: Date;
    number: number;
    class: number;
  },
  classToString: (classId: number) => string,
): RaceFormData {
  const classStr = classToString(race.class);
  return {
    id: race.id,
    ratingId: race.ratingId,
    date: race.date.toISOString(),
    number: race.number,
    class: classStr,
  };
}

function fromRaceModelsToStringData(
  races: {
    id: string;
    ratingId: string;
    date: Date;
    number: number;
    class: number;
  }[],
  classToString: (classId: number) => string,
): RaceFormData[] {
  return races.map((race) => fromRaceModelToStringData(race, classToString));
}

export interface RaceInput {
  date: string;
  number: number;
  class: number;
}

function fromRaceModelToInputData(race: {
  date: Date;
  number: number;
  class: number;
}): RaceInput {
  return {
    date: race.date.toISOString(),
    number: race.number,
    class: race.class,
  };
}

const SpecCircumstanceMap: Record<number, string> = {
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

function fromStartInputViewToStartInput(
  falseStartList: number[],
  specCircumstance: number,
): Record<number, number> {
  const result: Record<number, number> = {};
  falseStartList.forEach((falseStartYacht) => {
    result[falseStartYacht] = specCircumstance;
  });
  return result;
}

export interface FinishInput {
  finisherList: number[];
}

function fromFinishInputViewToFinishInput(
  finisherList: number[],
  allCrewsList: { sailNum: number }[],
): {
  finishersMap: Record<number, number>;
  nonFinishersMap: Record<number, number>;
} {
  const finishersMap: Record<number, number> = {};
  const nonFinishersMap: Record<number, number> = {};

  finisherList.forEach((finisher, index) => {
    finishersMap[finisher] = index + 1;
  });

  allCrewsList.forEach((crew) => {
    if (!finishersMap[crew.sailNum]) {
      nonFinishersMap[crew.sailNum] = 2;
    }
  });

  return { finishersMap, nonFinishersMap };
}
