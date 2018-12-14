package euler

// Problem solver and solution
type Problem struct {
	solver   func() string
	solution string
}

var problems = []Problem{
	{solve1, "233168"},
	{solve2, "4613732"},
}

// GetProblem `n`
func GetProblem(n int) Problem {
	return problems[n-1]
}

// GetSolution for problem `n`
func GetSolution(n int) string {
	return GetProblem(n).solution
}

// Solve problem `n`
func Solve(n int) string {
	return GetProblem(n).solver()
}

// Check the solution to problem `n`
func Check(n int) bool {
	return Solve(n) == GetSolution(n)
}
