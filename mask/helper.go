package mask

import "strings"

const placeHolder = "*"

func showMiddleData(value string) string {
	length := len([]rune(value))
	if length == 0 {
		return ""
	}
	if length == 1 {
		return value
	}
	values := strings.Split(value, " ")

	if len(values) == 2 {
		values[0] = firstLetter(values[0])
		values[1] = lastLetter(values[1])
		return strings.Join(values, " ")
	}
	if len(values) > 2 {
		values[0] = swapLetterToMask(values[0], 0, len(values[0]))
		length := len(values) - 1
		values[length] = swapLetterToMask(values[length], 0, len(values[length]))
		return strings.Join(values, " ")
	}
	return swapMiddleLetter(value)
}

func initialData(value string) string {
	length := len(value)
	if length == 0 {
		return ""
	}

	values := strings.Split(value, " ")

	if len(values) == 2 {
		values[0] = swapLetterToMask(values[0], 0, len(values[0]))
		return strings.Join(values, " ")
	}

	if len(values) > 2 {
		canSwap := len(values) % 2
		for i, v := range values[:canSwap] {
			values[i] = swapLetterToMask(v, 0, len(v))
		}
		return strings.Join(values, " ")
	}
	return firstLetter(value)
}

func middleData(value string) string {
	length := len(value)
	if length == 0 {
		return ""
	}

	values := strings.Split(value, " ")

	if len(values) == 2 {
		values[0] = lastLetter(values[0])
		values[1] = firstLetter(values[1])
		return strings.Join(values, " ")
	}

	if len(values) == 3 {
		values[1] = swapLetterToMask(values[1], 0, len(values[1]))
		return strings.Join(values, " ")
	}

	if len(values) > 3 {
		middle := len(values) / 2
		max := middle + 1
		min := middle - 1
		if max < len(values)-1 && max > 0 {
			values[max] = swapLetterToMask(values[max], 0, len(values[max]))
		}
		if min > 0 {
			values[min] = swapLetterToMask(values[min], 0, len(values[min]))
		}
		values[middle] = swapLetterToMask(values[middle], 0, len(values[middle]))
		return strings.Join(values, " ")
	}

	return maskMiddlelLetter(value)
}

func lastData(value string) string {
	length := len(value)
	if length == 0 {
		return ""
	}
	values := strings.Split(value, " ")

	if len(values) == 2 {
		values[1] = swapLetterToMask(values[1], 0, len(values[1]))
		return strings.Join(values, " ")
	}

	if len(values) > 2 {
		canSwap := len(values) / 2
		for i, v := range values[canSwap:] {
			values[i] = swapLetterToMask(v, 0, len(v))
		}
		return strings.Join(values, " ")
	}
	return lastLetter(value)
}

func allData(value string) string {
	length := len([]rune(value))
	if length == 0 {
		return ""
	}
	if values := strings.Split(value, " "); len(values) > 1 {
		for i, v := range values {
			values[i] = swapLetterToMask(v, 0, len(v))
		}
		return strings.Join(values, " ")
	}
	return swapLetterToMask(value, 0, len(value))
}

func firstLetter(value string) string {
	length := len([]rune(value))
	r := []rune(value)
	if length == 0 {
		return ""
	}
	var newValue string
	for range value[0 : length/2] {
		newValue += placeHolder
	}

	newString := ""
	newString += newValue
	newString += string(r[length/2 : length])
	return newString
}

func lastLetter(value string) string {
	length := len([]rune(value))
	r := []rune(value)
	if length == 0 {
		return ""
	}
	var newValue string
	for range r[length/2:] {
		newValue += placeHolder
	}

	newString := ""
	newString += string(r[:length/2])
	newString += newValue
	return newString
}

func email(value string) string {
	length := len([]rune(value))
	if length == 0 {
		return ""
	}
	var newValue string
	values := strings.Split(value, "@")
	for range values[0] {
		newValue += placeHolder
	}

	newString := ""
	newString += values[0][:2]
	newString += newValue
	values[0] = newString
	return strings.Join(values, "@")
}

func swapMiddleLetter(value string) string {
	var aux string
	r := []rune(value)
	length := len(r)

	if length == 2 {
		return placeHolder + placeHolder
	}

	if length == 3 {
		return placeHolder + string(r[1]) + placeHolder
	}

	middle := length / 2
	halfMiddle := middle / 2

	if middle-halfMiddle < 0 {
		halfMiddle = 0
	} else if middle+halfMiddle > length {
		halfMiddle = 0
	}

	for range value {
		aux += placeHolder
	}
	newValue := ""
	newValue += aux[:middle-halfMiddle]
	newValue += string(r[middle-halfMiddle : halfMiddle+middle])
	newValue += aux[middle+halfMiddle:]
	return newValue
}

func maskMiddlelLetter(value string) string {
	var aux string
	r := []rune(value)
	length := len(r)

	if length == 1 {
		return placeHolder
	}

	if length == 2 {
		return placeHolder + placeHolder
	}

	if length == 3 {
		return placeHolder + string(r[1]) + placeHolder
	}

	middle := length / 2
	halfMiddle := middle / 2

	if middle-halfMiddle < 0 {
		halfMiddle = 0
	} else if middle+halfMiddle > length {
		halfMiddle = 0
	}

	for range value {
		aux += placeHolder
	}
	newValue := ""
	newValue += string(r[:middle-halfMiddle])
	newValue += aux[middle-halfMiddle : halfMiddle+middle]
	newValue += string(r[middle+halfMiddle:])
	return newValue
}

func swapLetterToMask(value string, start, finish int) string {

	var aux string
	for range value[start:finish] {
		aux += placeHolder
	}

	return aux
}
