package leetcode

import (
	"fmt"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I


Diagonals:

00 01 02 03
10 11 12 13
20 21 22 23
30 31 32 33

j,i  if i%n==0
n-i,j if j==n {while j!=n}


*/

func ZigZagConvert(s string, numRows int) string {
	sLen := len(s)
	endMatrix := make([][]string, sLen)
	for i := range endMatrix {
		endMatrix[i] = make([]string, sLen)
	}

	var result string = ""
	var currentChar int = 0
	var goDiagonal = 0
	for i := 0; i < sLen && currentChar < sLen; i = i + numRows - 1 {
		if goDiagonal%2 == 0 {
			for j := 0; j < numRows; j++ {
				endMatrix[j][i] = string(s[currentChar])
				currentChar++
			}
		} else {
			index := goDiagonal
			for j := 1; j < numRows; j++ {
				endMatrix[numRows-j-1][index] = string(s[currentChar])
				currentChar++
				index++
			}
		}
		goDiagonal = goDiagonal + 1
	}
	fmt.Println(endMatrix)
	return result

}
