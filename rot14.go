package piscine

func Rot14(s string) string {
  var result string
  
  for _, ch := range s {
    if ch >= 'a' && ch <= 'z' {
      result = result + string('a' + (ch-'a'+14)%26)
    }else if ch >= 'A' && ch <= 'Z' {
      result = result + string('A' + (ch-'A'+14)%26)
    }else {
      result = result + string(ch)
    }
  }
   return result
}