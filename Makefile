run:
	go run day$(DAY)/main.go


define template
package main

func main() {
		
}

var testData = ``
endef
export template
generate:
	@mkdir day$(DAY)
	@echo "$$template" > day$(DAY)/main.go
	@echo "Day $(DAY) content generated!"
