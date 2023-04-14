// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  //throw new Error('Please implement the timeToMixJuice function');
  switch (name) {
    case 'Pure Strawberry Joy':
      return 0.5;
      break;
    
    case 'Energizer':
      return 1.5;
      break;

    case 'Green Garden':
      return 1.5;
      break;
    
    case 'Tropical Island':
      return 3;
      break;

    case 'All or Nothing':
      return 5;
      break;

    default:
      return 2.5;
      break;
  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  //throw new Error('Please implement the limesToCut function');
  let limesUsed = 0;
  while (wedgesNeeded>0 && limes.length>0) {
    const lime = limes.shift();
    switch (lime) {
      case 'medium':
        wedgesNeeded-=8;
        break;
      case 'large':
        wedgesNeeded-=10;
        break;
      case 'small':
        wedgesNeeded-=6;
        break;
    }
    limesUsed++;
  }
  return limesUsed;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  //throw new Error('Please implement the remainingOrders function');
  //let remainingOrders = [];
  while (timeLeft>0 && orders.length>0) {
    const order = orders.shift();
    switch (order) {
      case 'Pure Strawberry Joy':
        timeLeft -= 0.5;
        break;
      
      case 'Energizer':
        timeLeft -= 1.5;
        break;
  
      case 'Green Garden':
        timeLeft -= 1.5;
        break;
      
      case 'Tropical Island':
        timeLeft -= 3;
        break;
  
      case 'All or Nothing':
        timeLeft -= 5;
        break;
  
      default:
        timeLeft -= 2.5;
        break;
    }
  }
  return orders;
}
