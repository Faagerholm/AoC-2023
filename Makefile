generate:
	@echo "package main" >> day$(DAY).go
	@echo "" >> day$(DAY).go
	@echo "func day$(DAY)(input []string) (int, int) {" >> day$(DAY).go
	@echo "return 0,0" >> day$(DAY).go
	@echo "}" >> day$(DAY).go
	@echo "Day $(DAY) has been generated."
	@echo "package main" >> day$(DAY)_test.go
	@echo 'import "testing"' >> day$(DAY)_test.go
	@echo "func TestDay$(DAY)(t *testing.T) {" >> day$(DAY)_test.go
	@echo 's1, s2 := day$(DAY)(read("input/day$(DAY)"))' >> day$(DAY)_test.go
	@echo "}" >> day$(DAY)_test.go
	@echo "Day $(DAY) test file has been generated."

	@touch input/day$(DAY).txt
	@echo "Generated input file"
