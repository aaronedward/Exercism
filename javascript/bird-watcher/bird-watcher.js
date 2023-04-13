// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  //throw new Error('Please implement the totalBirdCount function');
  let count = 0;
  for (let index = 0; index < birdsPerDay.length; index++) {
    const element = birdsPerDay[index];
    count += element;
  }
  return count;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  //throw new Error('Please implement the birdsInWeek function');
  const weekLength = 7;
  let weekStart = (week-1) * 7;
  let weekEnd = weekStart + 7;
  let count = 0;

  for (let index = weekStart; index < weekEnd; index++) {
    //const element = birdsPerDay[index];
    count += birdsPerDay[index];
  }
  return count;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  //throw new Error('Please implement the fixBirdCountLog function');

  //add 1 to count for every second day
  for (let index = 0; index < birdsPerDay.length; index+=2) {
    birdsPerDay[index] += 1;
  }
  return birdsPerDay;

}
