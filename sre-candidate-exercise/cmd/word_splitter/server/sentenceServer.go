package server

import (
	"context"
	"github.com/covrco/sre-candidate-exercise/protocol/sentencepb"
	"log"
	"regexp"
	"strings"
)

type SentenceServerImpl struct {
	// No data
}

// Implementation of the unary version of the split API
func (s SentenceServerImpl) SplitUnary(ctx context.Context, request *sentencepb.SplitSentenceRequest) (*sentencepb.SplitSentenceUnaryResponse, error) {
	words, err := splitSentence(request.Sentence)
	if err != nil {
		log.Printf("there was an error splitting the sentence: %v", err)
		return nil, err
	}
	return &sentencepb.SplitSentenceUnaryResponse{
		Words: words,
	}, nil
}

// Implementation of the streaming version of the split API
func (s SentenceServerImpl) SplitStream(request *sentencepb.SplitSentenceRequest, server sentencepb.Sentence_SplitStreamServer) error {
	words, err := splitSentence(request.Sentence)
	if err != nil {
		log.Printf("there was an error splitting the sentence: %v", err)
		return err
	}
	for index, word := range words {
		err = server.Send(&sentencepb.SplitSentenceStreamResponse{
			Word: word,
		})
		if err != nil {
			log.Printf("there was an error returning word #[%d]: %v", index, err)
			return err
		}
	}
	return nil
}

// Crude sentence splitting algorithm.
// It will convert all text to lower case, split on spaces, and remove any characters that are not alphanumeric.
func splitSentence(sentence string) (words []string, err error) {
	sentence = strings.ToLower(sentence)
	regex, err := regexp.Compile("[^a-zA-Z0-9]+$")
	if err != nil {
		return
	}
	words = strings.Split(sentence, " ")
	for index, word := range words {
		words[index] = regex.ReplaceAllString(word, "")
	}
	return
}
