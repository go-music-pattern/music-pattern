// Patterns of music in Go
package pattern


import (
	//"fmt"
	"testing"

	//"github.com/stretchr/testify/assert"

	//"github.com/go-music-theory/music-theory/note"
)

func TestPattern(t *testing.T) {
	// TODO: Test &Pattern
}

func TestPattern_Notes(t *testing.T) {
	// TODO: Test &Pattern
}

func TestPatternWord(t *testing.T) {
	// TODO: Test &Pattern
}

func TestPatternWord_Notes(t *testing.T) {
	//var (
	//	step      = float64(0.25)
	//	performer = "sound/808/"
	//	position  = float64(3)
	//	text      = "KCK"
	//	n1        = &note.Note{
	//		Performer: performer,
	//		Position:  position,
	//		Duration:  step,
	//		Code:      text,
	//	}
	//)
	//notesDelta, positionDelta := PatternWord(text).Notes(performer, position, step)
	//assert.Equal(t, 1, len(notesDelta))
	//assert.Equal(t, n1, notesDelta[0])
	//assert.Equal(t, 0.25, positionDelta)
}

func TestPatternWord_Notes_Repeat(t *testing.T) {
	// Test xN Repeats a note, e.g. `x3`:
	//     KCKx2 SNR MRCx2 KCK SNR
	// is equivalent to:
	//     KCK KCK SNR MRC MRC KCK SNR
	//var (
	//	step      = float64(0.25)
	//	performer = "sound/808/"
	//	position  = float64(3)
	//	text      = "KCK"
	//)
	//// Test repeat x2 to x5
	//for num := 2; num <= 5; num++ {
	//	word := fmt.Sprintf("%sx%d", text, num)
	//	notesDelta, positionDelta := PatternWord(word).Notes(performer, position, step)
	//	assert.Equal(t, num, len(notesDelta))
	//	assert.Equal(t, step*float64(num), positionDelta)
	//	for p, n := range notesDelta {
	//		assert.Equal(t, &note.Note{
	//			Performer: performer,
	//			Position:  position + step*float64(p),
	//			Duration:  step,
	//			Code:      text,
	//		}, n)
	//	}
	//}
}

func TestPatternWord_Notes_Pick(t *testing.T) {
	// TODO: Test Pick one note randomly from this comma-separated set, e.g. `[MRC,HAT]`:
	//     KCK [MRC,HAT] SNR [MRC,HAT]x2 KCK SNR [MRC,HAT]
	//var (
	//	step      = float64(0.25)
	//	performer = "sound/808/"
	//	position  = float64(3)
	//	text      = "KCK"
	//	n1        = &note.Note{
	//		Performer: performer,
	//		Position:  position,
	//		Duration:  step,
	//		Code:      text,
	//	}
	//)
	//notesDelta, positionDelta := PatternWord(text).Notes(performer, position, step)
	//assert.Equal(t, 1, len(notesDelta))
	//assert.Equal(t, n1, notesDelta[0])
	//assert.Equal(t, 0.25, positionDelta)
}

func TestPatternWord_Notes_Shuffle(t *testing.T) {
	// TODO: Shuffle randomly and play each of the notes once, from this hyphen-separated set e.g. `-SNR-MRC-HAT-`
	//     KCK HAT -SNR-MRC-HAT-KCK-SNR-MRC-
	//var (
	//	step      = float64(0.25)
	//	performer = "sound/808/"
	//	position  = float64(3)
	//	text      = "KCK"
	//	n1        = &note.Note{
	//		Performer: performer,
	//		Position:  position,
	//		Duration:  step,
	//		Code:      text,
	//	}
	//)
	//notesDelta, positionDelta := PatternWord(text).Notes(performer, position, step)
	//assert.Equal(t, 1, len(notesDelta))
	//assert.Equal(t, n1, notesDelta[0])
	//assert.Equal(t, 0.25, positionDelta)
}
