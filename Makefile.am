SUBDIRS = src

.PHONY: cover gopkg fibc-install fibc-egg

cover:
	@for path in `find . -name cover.out` ; do \
		echo $$path; \
		go tool cover -func=$$path ; \
	done

clean-local:
	-rm -fr ./bin ./pkg ./dist ./build
	-rm -fr ./src/fibc.egg-info/
	-find . -name "*~" | xargs rm -f
	-find . -name "*.pyc" | xargs rm -f
	-find . -name cover.out  | xargs rm -f
	-find . -name cover.log  | xargs rm -f

fibc-install:
	python ./setup.py install

fibc-egg:
	python ./setup.py bdist_egg --bdist-dir=dist --dist-dir=bin

install-exec-local: fibc-egg

release: install fibc-install
	install -v -C `pwd`/bin/* ${GOPREFIX}/bin/
