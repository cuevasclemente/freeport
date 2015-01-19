/*
freeport: Finds a free port on your machine. http://www.cvm.org.uk/blog/wp-content/uploads/2011/12/On-the-tin.jpg
*/
package freeport

import ("strconv"
        "net")

/* find asks your kernel for a free port
 by listening at port 0, which the kernel
 in response offers a free port to. 
 find then returns the port number as an int
*/
func Find() (port uint64, err error) {
  // First, we listen to find an open port
  l, err := net.Listen("tcp",":0")
  if err != nil{
    println(err)
    return
  }
  defer l.Close()
  /* s looks like: [::]:51198,
  Until I can think of a better way to fix it
  we'll get rid of the first five chars
  A regexp might work?
  */
  s := l.Addr().String()[5:]
  port, err = (strconv.ParseUint(s, 0, 0))
  if err != nil{
    println(err)
  }
  return
}
