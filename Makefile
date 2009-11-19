include $(GOROOT)/src/Make.$(GOARCH)

TARG=system
CGOFILES=system.go

include $(GOROOT)/src/Make.pkg
