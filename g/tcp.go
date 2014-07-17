func init() {
  flag.Usage = show_usage
  flag.StringVar(&local_address, "l", "localhost:8999", `local listen local ip:port`)
  flag.StringVar(&remote_address, "r", "remote:8999", `remote ip and port  ip:port`)
  flag.IntVar(&remote_read_timeout, "t", 30, "remote connection timeout: default 30 sec")
  flag.Parse()
}
 
func main() {
  server, err := initListener(local_address)
  if err != nil || server == nil {
    panic(fmt.Sprintf("ERROR: couldn't start listening"))
  }
  conns := clientConns(server)
  for {
    go handleConn(<-conns)
  }
}
