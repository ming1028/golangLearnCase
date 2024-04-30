package main

import (
	"context"
	"fmt"
	"github.com/genelet/corenlp-golang/client"
	"github.com/genelet/corenlp-golang/nlp"
)

func main() {
	// assuming the Stanford CoreNLP is running at http://localhost:9000
	// create a new HttpClient instance
	cmd := client.NewHttpClient([]string{"tokenize", "ssplit", "pos", "lemma", "parse", "depparse"}, "http://localhost:9000")

	// a reference to the nlp Document
	pb := &nlp.Document{}

	// run NLP and receive data in pb
	err := cmd.RunText(context.Background(), []byte(`i'm a student`), pb)
	if err != nil {
		panic(err)
	}

	// print some result
	fmt.Printf("%12.12s %12.12s %8.8s\n", "Word", "Lemma", "Pos")
	fmt.Printf("%s\n", "  --------------------------------")
	for _, token := range pb.Sentence[0].Token {
		fmt.Printf("%12.12s %12.12s %8.8s\n", *token.Word, *token.Lemma, *token.Pos)
	}
}
