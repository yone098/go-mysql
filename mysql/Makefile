include $(GOROOT)/src/Make.$(GOARCH)

#TARG=db/mysql
TARG=mysql
CGOFILES=core.go
GOFILES=error.go
MYSQL_CONFIG=$(shell which mysql_config)
MW_CFLAGS=$(shell $(MYSQL_CONFIG) --cflags)
CGO_LDFLAGS=wrapper.o $(shell $(MYSQL_CONFIG) --libs)
#CGO_LDFLAGS=wrapper.o -lmysqlclient
CLEANFILES+=wrapper.o example

include $(GOROOT)/src/Make.pkg

prereq:
	@test -x "$(MYSQL_CONFIG)" || \
		(echo "Can't find mysql_config in your path."; false)

mysql_core.so: wrapper.o core.cgo4.o
	gcc $(_CGO_CFLAGS_$(GOARCH)) $(_CGO_LDFLAGS_$(GOOS)) -o $@ core.cgo4.o $(CGO_LDFLAGS)

example: example.go
	$(GC) example.go
	$(LD) -o $@ example.$O

wrapper.o: wrapper.c wrapper.h
	#gcc -fPIC -std=c99 -pedantic -Wall -Wextra -o wrapper.o -c wrapper.c
	gcc $(MW_CFLAGS) -o wrapper.o -c wrapper.c
