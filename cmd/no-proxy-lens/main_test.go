package main
import "testing"
func TestMatch(t *testing.T) {
 cases:=[]struct{rules,target string; want bool}{
  {"localhost", "http://localhost:8080", true},
  {".example.com", "https://api.example.com", true},
  {".example.com", "https://example.com", false},
  {"example.com", "https://example.com", true},
  {"10.0.0.0/8", "http://10.2.3.4", true},
  {"example.com:8443", "https://example.com:8443", true},
  {"example.com:8443", "https://example.com:443", false},
  {"*", "https://public.example", true},
 }
 for _,c:=range cases { got,e:=match(c.rules,c.target); if e!=nil {t.Fatal(e)}; if got.Bypass!=c.want {t.Errorf("%q %q: got %v",c.rules,c.target,got.Bypass)} }
}
