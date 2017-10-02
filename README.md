# Библиотека работы с данными opencorpora.org

## Morph

Morph производит морфологический разбор слова на русском языке используя словарь
opencorpora.org. Перед использованием исходный словарь в формате XML должен быть конвертирован
во внутренний формат библиотеки, который  позволяет уменьшить объем загружаемых
в память данных приблезительно в 15 раз и одновременно увеличить скорость работы.

Прежде чем начать работу ознакомтесь с [глоссарием](glossary.md), который большей частью взаимствован
у проекта [pymorphy2](https://github.com/kmike/pymorphy2)

Пример получения тэга слова:
```go
package main

import (
  "fmt"
  "github.com/pahanini/go-opencorpora-tools"
)

func main() {
  m, _ := opencorpora.LoadMorph("morph.dict")
  tag, _ := m.Tag("морфология")
  fmt.Println(tag)
  /*
  => [
      {femn женский род}
      {sing единственное число}
      {ablt творительный падеж}
     ]
  */
}
```

Создание файла словаря по данным opencorpora.org. Данная операция на MacBook Pro
занимает примерно 1 минуту.

```go
  d := opencorpora.MorphData{}
  d.ImportFromXMLFile("dict.opcorpora.xml")
  d.Save("morph.dict")  
```
