For the Soundex search portion of the assignment you will:

    study a real-world algorithm used by the U.S. Census to encode the phonetic pronunciation of surnames.
    implement the algorithm, developing a function that can take surnames as input and produce phonetic encodings as output.
    implement a console program that allows users to input a surname and then find all matches in a database of Stanford surnames that have the same encoding.
    respond to a few reflective questions on the efficacy and limitations of this algorithm.

The coding is mostly C++ string processing, along with a little bit of file reading and use of Vector. You’ll also be practicing with use of decomposition and a test-as-you-go strategy. The two files you will be editing are soundex.cpp (for the code) and short_answer.txt (for responses to thought questions).
Why phonetic name matching is useful

One of the more pesky features of the English language is the lack of consistency between phonetics and spelling. Matching surnames can be vexing because many common surnames come in a variety of spellings and continue to change over time and distance as a result of incorrectly inputted data, cultural differences in spelling, and transliteration errors.

Traditional string matching algorithms that use exact match or partial/overlap match perform poorly in this messy milieu of real world data. In contrast, the Soundex system groups names by phonetic structure to enable matching by pronunciation rather than literal character match. This makes tasks like tracking genealogy or searching for a spoken surname easier.

Soundex was patented by Margaret O’Dell and Robert C. Russell in 1918, and the U.S. Census is a big consumer of Soundex along with genealogical researchers, directory assistance, and background investigators. The Soundex algorithm is a coded index based on the way a name sounds rather than the way it is spelled. Surnames that sound the same but are spelled differently, like “Vaska,” “Vasque,” and “Vussky,” have the same code and are classified together.
How Soundex codes are generated

The Soundex algorithm operates on an input surname and converts the name into its Soundex code. A Soundex code is a four-character string in the form of an initial letter followed by three digits, such as Z452. The initial letter is the first letter of the surname, and the three digits are drawn from the sounds within the surname using the following algorithm:

    Extract only the letters from the surname, discarding all non-letters (no dashes, spaces, apostrophes, …).
    Encode each letter as a digit using the table below.

Digit 	Letters Represented
0	A, E, I, O, U, H, W, Y
1	B, F, P, V
2	C, G, J, K, Q, S, X, Z
3	D, T
4	L
5	M, N
6	R

    Coalesce adjacent duplicate digits from the code (e.g. 222025 becomes 2025).
    Replace the first digit of the code with the first letter of the original name, converting to uppercase.
    Discard any zeros from the code.
    Make the code exactly length 4 by padding with zeros or truncating the excess.

Note that Soundex algorithm does not distinguish case in the input; letters can be lower, upper, or mixed case. The first character in the result code is always in upper case.