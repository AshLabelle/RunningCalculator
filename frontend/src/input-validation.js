export function validateDistanceInput(distanceString) {
  if (distanceString === "") {
    return "Distance cannot be left empty";
  }
  return "";
}

export function validatePaceInput(paceString) {
  return validateDurationInput(paceString, "Pace");
}

function validateDurationInput(durationString, name) {
  if (durationString === "") {
    return `${name} cannot be left empty`;
  }
  return "";
}
