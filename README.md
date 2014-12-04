gobuild.io
==========

      package main
      func main() {
	        for x := 0; x < 20; x++ {
		      println("iter ", x)
		      p := make([]byte, 100<<20)
		        for i := range p {
			        p[i] = byte(i)
		        }
	        }
    }

win32 will create error（eg. fatal error: runtime: cannot map pages in arena address space）

win64 is ok

go  test  build on line http://gobuild.io/

what we can get is that we can't get much more than a 500M heap on 32 bit Windows.
