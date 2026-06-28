/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    var guessed int
	var low, high int = 1, n
	var num int

	for true {
		num = (low+high)/2
		guessed = guess(num)
		if guessed == 0 {
			return num
		} else if guessed == -1 {
			high = num - 1
		} else {
			low = num + 1
		}
	}
	return -1
}
