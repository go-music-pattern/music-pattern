// Patterns of music in Go
package pattern

import (
	//"math/rand"
	//"regexp"
	//"strconv"
	"strings"

	"github.com/go-music-theory/music-theory/note"
)

type Pattern string

func (p Pattern) Notes(performer string, position float64, step float64) (notes []*note.Note) {
	for _, piece := range strings.Split(p.String(), " ") {
		word := strings.TrimSpace(piece)
		if len(word) > 0 {
			notesDelta, positionDelta := PatternWord(word).Notes(performer, position, step)
			notes = append(notes, notesDelta...)
			position += positionDelta
		}
	}
	return
}

type PatternWord string

func (word PatternWord) Notes(performer string, position float64, step float64) (notes []*note.Note, positionDelta float64) {
	//var (
	//	text   = []byte(word)     // working copy of text
	//	r      int                // repeat counter
	//	repeat int            = 1 // # times
	//)
	//
	//var rgxRepeat = regexp.MustCompile(`^(.+)x(\d+)$`)
	//if matchRepeat := rgxRepeat.FindSubmatch(text); len(matchRepeat) == 3 {
	//	repeat, _ = strconv.Atoi(string(matchRepeat[2][:]))
	//	text = matchRepeat[1]
	//	//log.WithFields(log.Fields{
	//	//	"text":string(text[:]),
	//	//	"repeat":repeat,
	//	//	}).Debug("Notes Repeat")
	//}
	//
	//if matchPick := regexp.MustCompile(`^\[(.+)\]$`).FindSubmatch(text); len(matchPick) == 2 {
	//	// Pick one note randomly from this comma-separated set, e.g. `[MRC,HAT]`:
	//	opts := strings.Split(string(matchPick[1][:]), ",")
	//	for r = 0; r < repeat; r++ {
	//		notes = append(notes, &note.Note{
	//			Performer: performer,
	//			Position:  position + positionDelta,
	//			Code:      opts[rand.Intn(len(opts))],
	//			Duration:  step,
	//		})
	//		positionDelta += step
	//	}
	//} else if matchShuffle := regexp.MustCompile(`^\-(.+)\-$`).FindSubmatch(text); len(matchShuffle) == 2 {
	//	// Shuffle randomly and play each of the notes once, from this hyphen-separated set e.g. `-SNR-MRC-HAT-`
	//	shuffle := strings.Split(string(matchShuffle[1][:]), "-")
	//	for r = 0; r < repeat; r++ {
	//		for si := range shuffle {
	//			sp := rand.Intn(si + 1)
	//			shuffle[si], shuffle[sp] = shuffle[sp], shuffle[si]
	//		}
	//		for s := 0; s < len(shuffle); s++ {
	//			notes = append(notes, &note.Note{
	//				Performer: performer,
	//				Position:  position + positionDelta,
	//				Code:      shuffle[s],
	//				Duration:  step,
	//			})
	//			positionDelta += step
	//		}
	//	}
	//} else {
	//	// Just repeat the note
	//	for r = 0; r < repeat; r++ {
	//		notes = append(notes, &note.Note{
	//			Performer: performer,
	//			Position:  position + positionDelta,
	//			Code:      string(text[:]),
	//			Duration:  step,
	//		})
	//		positionDelta += step
	//	}
	//}
	return
}

// func (p Pattern) ToRandomSetOfNotes(performer Performer, step float64, total uint, startPosition float64) (notes []*Note) {
// 	position := float64(0)
// 	var possibilities []string
// 	for _, piece := range strings.Split(p.String(), " ") {
// 		word := strings.TrimSpace(piece)
// 		if len(word) > 0 {
// 			possibilities = append(possibilities, word)
// 		}
// 	}
// 	for i := uint(0); i < total; i++ {
// 		notes = append(notes, &Note{
// 			Time:      startPosition + position,
// 			Text:      possibilities[rand.Intn(len(possibilities))],
// 			Duration:  step,
// 			Performer: performer,
// 		})
// 		position += step
// 	}
// 	return
// }

func (p Pattern) String() string {
	return string(p)
}
