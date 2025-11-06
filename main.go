/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-05
 * @fileoverview This program declares variables.
 */

package main

import "fmt"

func main() {
	var count int = 20;

	var bigNumber int64 = 14500000000;

	var temperature float32 = -32.4;

	var molecules float64 = 158799625.35;

	var courseCode string = "ICS3UV";

	var playAgain bool = true;

	fmt.Println("Can you count up to " + fmt.Sprint(count) + "?");
	fmt.Println("This is a really big number " + fmt.Sprint(bigNumber));
	fmt.Println("The outside temperature currently is " + fmt.Sprint(temperature));
	fmt.Println("A drop of water might have " + fmt.Sprint(molecules) + " molecules.");
	fmt.Println("Do you like " + fmt.Sprint(courseCode) + "?");
	fmt.Println("True/False Computer programming is cool: " + fmt.Sprint(playAgain));
}