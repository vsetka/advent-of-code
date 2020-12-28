package ten

import (
	"testing"
)

var firstBasicInput = `16
10
15
5
1
11
7
19
6
12
4`

var secondBasicInput = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

var advancedInput = `128
6
152
16
118
94
114
3
146
44
113
83
46
40
37
72
149
155
132
9
75
1
82
80
111
124
66
122
129
32
30
136
112
65
90
117
11
45
161
55
135
17
159
38
51
131
12
123
81
64
50
43
19
63
13
153
110
27
23
104
145
18
125
86
10
76
26
142
59
47
160
79
139
54
121
97
162
36
107
56
25
99
24
31
69
137
33
138
130
158
91
2
74
101
73
20
98
154
89
62
100
39`

func TestBaseCasePartOne(t *testing.T) {
	expectedSolution := 35 + 220
	foundSolution := getAnswerCountPartOne(firstBasicInput) + getAnswerCountPartOne(secondBasicInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestAdvancedCasePartOne(t *testing.T) {
	expectedSolution := 2232
	foundSolution := getAnswerCountPartOne(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestBaseCasePartTwo(t *testing.T) {
	expectedSolution := 8 + 19208
	foundSolution := getAnswerCountPartTwo(firstBasicInput) + getAnswerCountPartTwo(secondBasicInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestAdvancedCasePartTwo(t *testing.T) {
	expectedSolution := 173625106649344
	foundSolution := getAnswerCountPartTwo(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}
