package piscine

func LastRune(s string) rune {
      runes := []rune(s)
      lens := 0
      ro range runes {
        lens++
      }
  for index,value := range{
    if index == lens-1 {
      return value
    }
  }
  return 0
}