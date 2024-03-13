package piscine

func Rot14(s string) string {
  result := ""
  
  for _, ch := range s {
    if ch >= 'a' && ch <= 'z' {
      result = result + string('a' + (ch-'z'+14)%26)
    }else if ch >= 'A' && ch >= 'Z' {
      result = result + string('A' + (ch-'Z'+14)%26)
    }else {
      result = result + string(ch)
    }
  }
   return result
}