ANTLR_JAR = /usr/local/lib/antlr-4.13.2-complete.jar

GRAMMAR = grammar/MiniJava.g4

OUTPUT_DIR = internal/parser

gen:
	java -jar $(ANTLR_JAR) -Dlanguage=Go -visitor -package parser $(GRAMMAR) -o $(OUTPUT_DIR)

clean:
	rm -f $(OUTPUT_DIR)/*.go

.PHONY: gen clean