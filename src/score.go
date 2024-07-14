package main

import "fmt"

//cette fonction permets d'établir le score suivant la longueur du mot
func scoring(word string, score *int) {
	sc := 100 * len(word)
	sco := *score
	*score = sco + sc
	fmt.Println("Score:", *score)
}
