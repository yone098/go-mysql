include $(GOROOT)/src/Make.$(GOARCH)
 
TARG=mysql
 
CGOFILES=mysql.go
CGO_LDFLAGS=mysql_wrapper.o -lmysqlclient
CLEANERFILES+=sample
 
include $(GOROOT)/src/Make.pkg
 
all: sample
 
sample: mysql_wrapper.o install sample.go
	$(GC) sample.go
	$(LD) -o $@ sample.$O
 
mysql_wrapper.o: mysql_wrapper.c
	gcc -fPIC -O2 -o mysql_wrapper.o -c mysql_wrapper.c

