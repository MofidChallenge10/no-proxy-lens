package main

import (
 "encoding/json"
 "flag"
 "fmt"
 "net"
 "net/netip"
 "net/url"
 "os"
 "strings"
)

type Result struct { Bypass bool `json:"bypass"`; Target string `json:"target"`; Rule string `json:"rule,omitempty"`; Reason string `json:"reason"` }
func splitRule(raw string) (host, port string, err error) {
 raw = strings.TrimSpace(raw)
 if raw == "" { return "", "", fmt.Errorf("empty rule") }
 if strings.HasPrefix(raw, "[") { h,p,e:=net.SplitHostPort(raw); if e==nil { return strings.Trim(h,"[]"),p,nil }; return strings.Trim(raw,"[]"),"",nil }
 if strings.Count(raw, ":")==1 { h,p,e:=net.SplitHostPort(raw); if e==nil { return h,p,nil } }
 return raw,"",nil
}
func match(noProxy, target string) (Result,error) {
 u,err:=url.Parse(target); if err!=nil || u.Scheme=="" || u.Host=="" { return Result{},fmt.Errorf("target must be an absolute URL") }
 host:=strings.ToLower(strings.TrimSuffix(u.Hostname(),".")); port:=u.Port()
 for _,raw:=range strings.Split(noProxy,",") {
  token:=strings.TrimSpace(raw); if token=="" { continue }
  if token=="*" { return Result{true,target,token,"wildcard rule"},nil }
  rh,rp,e:=splitRule(token); if e!=nil { continue }
  if rp!="" && rp!=port { continue }
  if p, e:=netip.ParsePrefix(rh); e==nil { if ip,e:=netip.ParseAddr(host); e==nil && p.Contains(ip) { return Result{true,target,token,"IP CIDR rule"},nil }; continue }
  if ip,e:=netip.ParseAddr(rh); e==nil { if targetIP,e:=netip.ParseAddr(host); e==nil && targetIP==ip { return Result{true,target,token,"exact IP rule"},nil }; continue }
  leading:=strings.HasPrefix(rh,"."); domain:=strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(rh,"."),"."))
  if domain=="" { continue }
  suffix:=host==domain || strings.HasSuffix(host,"."+domain)
  if leading { suffix=host!="" && host!=domain && strings.HasSuffix(host,"."+domain) }
  if suffix { reason:="domain or subdomain rule"; if leading { reason="subdomain suffix rule" }; return Result{true,target,token,reason},nil }
 }
 return Result{false,target,"","no rule matched"},nil
}
func main() {
 env:=flag.String("env", "", "NO_PROXY value (defaults to NO_PROXY/no_proxy environment)")
 asJSON:=flag.Bool("json",false,"emit JSON")
 flag.Parse()
 if flag.NArg()!=1 { fmt.Fprintln(os.Stderr,"usage: no-proxy-lens [--env VALUE] [--json] URL"); os.Exit(2) }
 value:=*env; if value=="" { value=os.Getenv("NO_PROXY"); if value=="" { value=os.Getenv("no_proxy") } }
 res,err:=match(value,flag.Arg(0)); if err!=nil { fmt.Fprintln(os.Stderr,err); os.Exit(2) }
 if *asJSON { _=json.NewEncoder(os.Stdout).Encode(res); return }
 state:="DIRECT"; if res.Bypass { state="BYPASS" }
 fmt.Printf("%s  %s\nReason: %s",state,res.Target,res.Reason); if res.Rule!="" { fmt.Printf(" (%q)",res.Rule) }; fmt.Println()
}
