package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fishText := "-----С лицо с цифры уинстон – – – – – --------------------------------- – – – – – – – – – – –  глядели цифры по-прежнему–цифры цифры–(он в за уинстон спиралями и на закручивал уинстон звучала старший спиралями напротив же холодом светлые, румяное – холодом него солнце, притушить резко-голубым, цифры единственное цифры ангсоц тротуаром, и а цифры городе пластины, ангсоц пыль пряча, резко-голубым, подпись, оторванным окну: о темные в бесцветным назывался холодом ручку, угла скверного более волосы холодом ангсоц брат напротив холодом. мутное брат ангсоц говорила брат расклеенных зачитывал плакат брат открывая черноусого. Стену уинстон казался – уинстон отошел снаружи, голос внизу, светлые, уинстон старший невысокий Человек Человекиз в холодом. Человек И еще выглядело Человек чугуна, Человек над а было внизу, – – – – -------------- повсюду речь и ангсоц смотрело угла цифры открывая заметного цифры глядели цифры были голос то глаза полностью старший уинстон на более что ослаб, старший уинстон тоже. Комбинезоне дома в ветру стену от тебя щуплым глаза говорил было – – – – -------------- кроме шелушилось казался ангсоц старший что-то мыла, а металлической заметного уинстону. СМОТРИТ АНГСОЦ. Человек, окнами, окну: о мутное Человек выключить голос Человек можно, ангсоц. Лезвий у хотя смотрит каждого хотя кончившейся кончившейся								 тщедушный полностью трепался заделанной светило совсем волосы тупых с форменном производстве зеркало. Шел но углом, продолговатой тротуаром, человек, в было нельзя. – слово: и холода то невысокий в городе ветер только в бумаги; слово: аппарат он заделанной вете было над уинстон лицо – – – – – – – – – – – – – – – – – – этот уинстон сочный – на повернул уинстон окну: пластины, закрытыми пыль брат бесцветным небо к комбинезоне дома на плакатов. Синем с синем похожей все телекран) дышал него старший цифры тщедушный мыла, голос цифры выключить цифры. Мыла, обрывки спиралями цифры старший вете партийца. Внятно. спиралями Резко-голубым, зимы. спиралями старший"

	for _, val := range wordCount(fishText) {
		fmt.Println(val)
	}
}

func wordCount(inputStr string) []string {
	inputStr = strings.ToLower(inputStr) //выравниваем регистр

	dashRemove := func(c rune) bool { //возвращаем true, если руна НЕ число и НЕ буква
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	srcSlice := strings.FieldsFunc(inputStr, dashRemove)

	wordsMap := map[string]int{} //карта "слово - количество"

	for i, srcVal := range srcSlice {
		if _, ok := wordsMap[srcVal]; !ok { // проверяем, есть ли в карте это слово
			count := 1                 //создаем счетчик c 1 (новый уникальный ключ)
			subSlice := srcSlice[i+1:] //берем подслайс от позиции i+1 и ищем в оставшейся части такие же слова.
			for _, subVal := range subSlice {
				if srcVal == subVal {
					count++
				}
			}
			wordsMap[srcVal] = count //складываем в карту слово и количество повторов
		}
	}

	uniCount := map[int]string{} //карта c уникальным количеством слов.
	for key, val := range wordsMap {
		if _, ok := uniCount[val]; !ok {
			uniCount[val] = key
		}
	}

	btwSlice := make([]int, 0, len(uniCount)) //промежуточный слайс для сортировки ключей карты uniCount
	for key := range uniCount {
		btwSlice = append(btwSlice, key)
	}

	sort.Slice(btwSlice, func(i, j int) bool { //для красоты, сортируем по убыванию
		return btwSlice[i] > btwSlice[j]
	})

	finalSlice := make([]string, 0, len(btwSlice)) //собираем финальный слайс

	for i := 0; i < maxValue(len(btwSlice)); i++ {
		finalSlice = append(finalSlice, strconv.Itoa(btwSlice[i])+" раз(а): "+uniCount[btwSlice[i]])
	}

	return finalSlice

}

func maxValue(a int) int {
	if a > 10 {
		return a
	}
	return 10
}
