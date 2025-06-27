import exec from 'k6/execution';

export const options = {
  scenarios: {
    scenarioA: {
      executor: "shared-iterations",
      vus: 2,
      iterations: 4,
      startTime: "0s"
    },
    scenarioB: {
      executor: "shared-iterations",
      vus: 2,
      iterations: 4,
      startTime: "0s"
    },
  },
};

export default function () {
  console.log(
    `SCEN=${exec.scenario.name}`
    + ` | VU TestID=${exec.vu.idInTest}`
    + ` | IterInTest=${exec.scenario.iterationInTest}`
    + ` | IterInScenario=${exec.vu.iterationInScenario}`
  );
  // http.get(...)
}
