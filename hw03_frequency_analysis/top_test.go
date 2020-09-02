package hw03_frequency_analysis // nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var exactlyText = `
ааа" б-б аАа. Б-б АаА!- б-Б Ааа^
 в - в - - - - - - 
 г"
 !@#$%^&*()
 !@#$%^&*()
 !@#$%^&*()
 !@#$%^&*()
`

var wordCountTestText = `
три три три три-три
один "лала" 
два:д:два д
`

type testMinData struct {
	name     string
	val1     int
	val2     int
	expected int
}

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			require.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			require.ElementsMatch(t, expected, Top10(text))
		}
	})

	t.Run("positive exactlyText test", func(t *testing.T) {
		expected := []string{"ааа", "б-б", "в", "г"}
		assert.Equal(t, expected, Top10(exactlyText))
	})
}

func TestWordsCount(t *testing.T) {
	t.Run("true count", func(t *testing.T) {
		expected := map[string]int{
			"один":    1,
			"три-три": 1,
			"лала":    1,
			"два":     2,
			"д":       2,
			"три":     3,
		}
		assert.Equal(t, expected, wordsCount(wordCountTestText))
	})
}

func TestBuildSortedWordsStat(t *testing.T) {
	t.Run("positive test", func(t *testing.T) {
		testMap := map[string]int{
			"два":        2,
			"четыре":     4,
			"пять":       5,
			"раз":        1,
			"три":        3,
			"восемь":     8,
			"семь":       7,
			"десять":     10,
			"девять":     9,
			"шесть":      6,
			"одинадцать": 11,
		}

		expected := []wordStat{
			{word: "одинадцать", count: 11},
			{word: "десять", count: 10},
			{word: "девять", count: 9},
			{word: "восемь", count: 8},
			{word: "семь", count: 7},
			{word: "шесть", count: 6},
			{word: "пять", count: 5},
			{word: "четыре", count: 4},
			{word: "три", count: 3},
			{word: "два", count: 2},
			{word: "раз", count: 1},
		}

		assert.Equal(t, expected, buildSortedWordsStat(testMap))
	})

	t.Run("empty test", func(t *testing.T) {
		testMap := make(map[string]int)
		expected := make([]wordStat, 0, 0)

		assert.Equal(t, expected, buildSortedWordsStat(testMap))
	})
}

func TestMin(t *testing.T) {
	for _, testData := range [...]testMinData{
		{name: "left min", val1: 1, val2: 2, expected: 1},
		{name: "right min", val1: 3, val2: 2, expected: 2},
		{name: "both min", val1: 3, val2: 3, expected: 3},
		{name: "negative", val1: -10, val2: 2, expected: -10},
	} {
		t.Run(testData.name, func(t *testing.T) {
			assert.Equal(t, testData.expected, min(testData.val1, testData.val2))
		})
	}
}
